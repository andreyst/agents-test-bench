package transport

import (
	"fmt"
	"net"
)

// TCPTransport is a TCP transport.
type TCPTransport struct {
	endpoint string
	conn     net.Conn
}

// NewTCPTransport creates new TCP transport with specified endpoint.
func NewTCPTransport(endpoint string) (transport *TCPTransport, err error) {
	transport = new(TCPTransport)
	transport.endpoint = endpoint
	conn, err := net.Dial("tcp", transport.endpoint)
	if err != nil {
		return nil, err
	}
	transport.conn = conn

	return transport, nil
}

func (transport *TCPTransport) Write(p []byte) (n int, err error) {
	n, err = fmt.Fprintf(transport.conn, string(p))
	return n, err
}
