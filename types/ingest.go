package types

type Ingester struct {
	Subject  string
	Message  string
	Metadata map[string]string
	Raw      interface{}
}
