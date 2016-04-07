package discovery

import (
	"errors"

	"github.com/hashicorp/mdns"
)

// Discover will register
type Discover interface {
	Register() (*mdns.Server, error)
}

// New will new
func New(discoveryType string, componentName string, serviceName string) (Discover, error) {
	switch discoveryType {
	case "mDNS":
		return NewMdnsServer(componentName, serviceName)
	}
	return nil, errors.New("Unknown discovery service")
}
