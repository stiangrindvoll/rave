package discovery

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	host, _ := os.Hostname()
	service, err := mdns.NewMDNSService(host,
		"_rave._tcp",
		"",
		"",
		1623,
		nil,
		[]string{dir.serviceName, dir.componentName},
	)
	fmt.Println(service, err)
	// Create the mDNS server, defer shutdown
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return nil, err
	}
	return server, nil
}

// List will list all available services
func List() {
	entriesCh := make(chan *mdns.ServiceEntry, 4)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry:\n\tName: %v\n\tHost: %v\n\tIp: %v\n\tPort: %v\n\tInfo: %v\n", entry.Name, entry.Host, entry.AddrV4, entry.Port, entry.InfoFields)

			res, err := http.Get(fmt.Sprintf("http://%v:%v/index.html", entry.AddrV4, entry.Port))
			if err != nil {
				log.Fatal(err)
			}
			robots, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", robots)

		}
	}()

	// Start the lookup
	mdns.Lookup("_rave._tcp", entriesCh)
	close(entriesCh)
}
