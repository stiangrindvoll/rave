package discovery

import (
	"fmt"

	"github.com/hashicorp/mdns"
)

// Mdns for data
type Mdns struct {
	componentName string
	serviceName   string
}

// NewMdnsServer creates a new mDNS server
func NewMdnsServer(componentName string, serviceName string) (*Mdns, error) {
	return &Mdns{
		componentName: componentName,
		serviceName:   serviceName,
	}, nil

}

// Register will register the service
func (dir *Mdns) Register() (*mdns.Server, error) {
	service, err := mdns.NewMDNSService(dir.serviceName,
		"_rave._tcp",
		"",
		"",
		1623,
		nil,
		[]string{dir.componentName},
	)
	fmt.Println(service, err)
	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return nil, err
	}
	return server, nil
}

// GetService will list all available services
func GetService(key string) (ip, port string) {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry:\n\tName: %v\n\tHost: %v\n\tIp: %v\n\tPort: %v\n\tInfo: %v\n", entry.Name, entry.Host, entry.AddrV4, entry.Port, entry.InfoFields)
			for _, k := range entry.InfoFields {
				fmt.Println(fmt.Sprintf("compare:\"%v %v\"", k, key))
				if k == key {
					ip = fmt.Sprintf("%s", entry.AddrV4)
					port = fmt.Sprintf("%v", entry.Port)
					return
				}
			}

		}
	}()

	// Start the lookup
	mdns.Lookup("_rave._tcp", entriesCh)
	close(entriesCh)

	return
}
