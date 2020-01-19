package serializer

import b64 "encoding/base64"

// Base64Serializer creates a message by converting byte array to base64-encoded string.
type Base64Serializer struct {
}

// Serialize serializes a message
func (mm *Base64Serializer) Serialize(p []byte) (res []byte) {
	return []byte(b64.StdEncoding.EncodeToString(p))
}
