package discovery

import (
	"fmt"
	"net"

	"github.com/hashicorp/mdns"
)

// Mdns for data
type Mdns struct {
	server        *mdns.Server
	componentName string
	serviceName   string
}

// Hosts will have information about services available
type Hosts struct {
	IP, Port string
}

// NewMdnsServer creates a new mDNS server
func NewMdnsServer(componentName string, serviceName string) (*Mdns, error) {
	return &Mdns{
		server:        nil,
		componentName: componentName,
		serviceName:   serviceName,
	}, nil

}

// Register will register the service
func (dir *Mdns) Register() error {
	var IP []net.IP
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				IP = append(IP, ipnet.IP)
			}
		}
	}

	fmt.Println("List IPs:", IP)
	service, err := mdns.NewMDNSService(dir.serviceName,
		"_rave._tcp.",
		"",
		"",
		1623,
		IP,
		[]string{dir.componentName},
	)
	fmt.Println(service, err)
	// Create the mDNS server, defer shutdown
	dir.server, err = mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return err
	}
	return nil
}

// Close will shutdown the registered server
func (dir *Mdns) Close() error {
	return dir.server.Shutdown()
}

// GetService will list all available services
func GetService(key string) (h []Hosts) {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry:\n\tName: %v\n\tHost: %v\n\tIp: %v\n\tPort: %v\n\tInfo: %v\n", entry.Name, entry.Host, entry.AddrV4, entry.Port, entry.InfoFields)
			for _, k := range entry.InfoFields {
				fmt.Println(fmt.Sprintf("compare:\"%v %v\"", k, key))
				if k == key {
					h = append(h, Hosts{IP: fmt.Sprintf("%s", entry.AddrV4), Port: fmt.Sprintf("%v", entry.Port)})

					//					ip = fmt.Sprintf("%s", entry.AddrV4)
					//					port = fmt.Sprintf("%v", entry.Port)
				}
			}
		}
	}()

	// Start the lookup
	mdns.Lookup("_rave._tcp", entriesCh)
	close(entriesCh)

	return
}
