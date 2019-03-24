package parser

// PortParser is used to parse incoming data
type PortParser interface {
	Parse(string)
}
