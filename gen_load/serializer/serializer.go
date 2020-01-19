package serializer

// Serializer is an interface to create a string message from byte array.
type Serializer interface {
	Serialize(p []byte) (res []byte)
}
