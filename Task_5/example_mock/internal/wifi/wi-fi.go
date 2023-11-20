package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type MyWiFiService struct {
	WiFi WiFi
}

func New(wifi WiFi) MyWiFiService {
	return MyWiFiService{WiFi: wifi}
}

func (service MyWiFiService) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	addrs := make([]net.HardwareAddr, 0, len(interfaces))

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service MyWiFiService) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	nameList := make([]string, 0, len(interfaces))

	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}

	return nameList, nil
}
