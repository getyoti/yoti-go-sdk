package facecapture

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
)

type UploadFaceCaptureImagePayload struct {
	ImageContentType string
	ImageContents    []byte

	body   *bytes.Buffer
	writer *multipart.Writer
}

func NewUploadFaceCaptureImagePayload(contentType string, contents []byte) *UploadFaceCaptureImagePayload {
	return &UploadFaceCaptureImagePayload{
		ImageContentType: contentType,
		ImageContents:    contents,
	}
}

func (p *UploadFaceCaptureImagePayload) Prepare() error {
	if p.ImageContentType == "" {
		return fmt.Errorf("ImageContentType must not be empty")
	}
	if len(p.ImageContents) == 0 {
		return fmt.Errorf("ImageContents must not be empty")
	}

	p.body = &bytes.Buffer{}
	p.writer = multipart.NewWriter(p.body)

	header := textproto.MIMEHeader{}
	header.Set("Content-Disposition", `form-data; name="binary-content"; filename="face-capture-image"`)
	header.Set("Content-Type", p.ImageContentType)

	part, err := p.writer.CreatePart(header)
	if err != nil {
		return fmt.Errorf("failed to create multipart part: %w", err)
	}

	_, err = part.Write(p.ImageContents)
	if err != nil {
		return fmt.Errorf("failed to write image contents: %w", err)
	}

	return p.writer.Close()
}

func (p *UploadFaceCaptureImagePayload) MultipartFormBody() *bytes.Buffer {
	return p.body
}

// ✅ Fixed return type to match SignedRequest expectations
func (p *UploadFaceCaptureImagePayload) Headers() map[string][]string {
	return map[string][]string{
		"Content-Type": {p.writer.FormDataContentType()},
	}
}
