package retrieve

import "fmt"

type RequiredResourceResponse interface {
	GetType() string
	String() string
}

type BaseRequiredResource struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	State string `json:"state"`
}

func (b *BaseRequiredResource) GetType() string {
	return b.Type
}

func (b *BaseRequiredResource) String() string {
	return fmt.Sprintf("Type: %s, ID: %s, State: %s", b.Type, b.ID, b.State)
}
