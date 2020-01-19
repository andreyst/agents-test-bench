package serializer

import b64 "encoding/base64"

import "strings"

// FluentBitMessageSerializer is a message maker that creates a JSON message as a base64-encoded string.
type FluentBitMessageSerializer struct {
}

// Serialize serializes a message.
func (mm *FluentBitMessageSerializer) Serialize(p []byte) (res []byte) {
	var b strings.Builder

	b.WriteString("[\"")
	b.WriteString(b64.StdEncoding.EncodeToString(p))
	b.WriteString("\"]")
	return []byte(b.String())
}
