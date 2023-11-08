package digitalidentity

type ReceiptItemKeyResponse struct {
	ID    string `json:"id"`
	Iv    []byte `json:"iv"`
	Value []byte `json:"value"`
}
