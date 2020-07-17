package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
	"time"
)

var (
	validFrom = ""
	validFor  = 2 * 365 * 24 * time.Hour
	isCA      = true
	rsaBits   = 2048
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

func certificatePresenceCheck(certPath string, keyPath string) (present bool) {
	if _, err := os.Stat(certPath); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func generateSelfSignedCertificate(certPath, keyPath, host string) error {
	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		log.Printf("failed to generate private key: %s", err)
		return err
	}

	notBefore, err := parseNotBefore(validFrom)
	if err != nil {
		log.Printf("failed to parse 'Not Before' value of cert using validFrom %q, error was: %s", validFrom, err)
		return err
	}

	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Printf("failed to generate serial number: %s", err)
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Yoti"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Printf("Failed to create certificate: %s", err)
		return err
	}

	err = createPemFile(certPath, derBytes)
	if err != nil {
		log.Printf("failed to create pem file at %q: %s", certPath, err)
		return err
	}
	log.Printf("written %s\n", certPath)

	err = createKeyFile(keyPath, priv)
	if err != nil {
		log.Printf("failed to create key file at %q: %s", keyPath, err)
		return err
	}
	log.Printf("written %s\n", keyPath)

	return nil
}

func createPemFile(certPath string, derBytes []byte) error {
	certOut, err := os.Create(certPath)

	if err != nil {
		log.Printf("failed to open "+certPath+" for writing: %s", err)
		return err
	}

	defer certOut.Close()
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	return err
}

func createKeyFile(keyPath string, privateKey interface{}) error {
	keyOut, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)

	if err != nil {
		log.Print("failed to open "+keyPath+" for writing:", err)
		return err
	}

	defer keyOut.Close()
	err = pem.Encode(keyOut, pemBlockForKey(privateKey))

	return err
}

func parseNotBefore(validFrom string) (notBefore time.Time, err error) {
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", validFrom)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse creation date: %s\n", err)
			return time.Time{}, err
		}
	}

	return notBefore, nil
}
