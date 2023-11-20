package wifi

import (
	"errors"
	"example_mock/internal/wifi/mocks"
	"fmt"
	"net"
	"strings"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type testCaseAdresses struct {
	addrs       []string // List of adresses
	errExpected error    // Expected error
}

// mockIfaces - creates Wi-Fi interfaces mocks based on given adresses
func mockIfacesAdresses(addrs []string) []*wifi.Interface {
	mockInterfaces := make([]*wifi.Interface, len(addrs))
	for i, addr := range addrs {
		hardwareAddr, err := net.ParseMAC(addr)
		if err != nil {
			panic(err)
		}

		mockInterfaces[i] = &wifi.Interface{
			HardwareAddr: hardwareAddr,
			Name:         fmt.Sprintf("MockInterface%d", i+1),
		}
	}
	return mockInterfaces
}

// mockIfaces - creates Wi-Fi interfaces mocks based on given names
func mockIfacesNames(names []string) []*wifi.Interface {
	mockInterfaces := make([]*wifi.Interface, len(names))
	for i, name := range names {
		mockInterfaces[i] = &wifi.Interface{
			HardwareAddr: nil,
			Name:         name,
		}
	}
	return mockInterfaces
}

// parseMACs processes adresses list and exports MAC-adresses
func parseMACs(addrs []string) []string {
	macAddresses := make([]string, len(addrs))
	for i, addr := range addrs {
		macAddresses[i] = addr[:17]
	}
	return macAddresses
}

func TestGetAdresses(t *testing.T) {
	mockWifi := mocks.NewWiFi(t)
	wifiService := MyWiFiService{WiFi: mockWifi}

	var testTable = []testCaseAdresses{
		{
			addrs:       []string{"00:11:22:33:44:55", "AA:BB:CC:DD:EE:FF"},
			errExpected: nil,
		},
		{
			addrs:       []string{},
			errExpected: errors.New("some error"),
		},
		{
			addrs:       []string{"11:22:33:44:55:66", "FF:EE:DD:CC:BB:AA", "01:23:45:67:89:AB"},
			errExpected: nil,
		},
	}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfacesAdresses(row.addrs), row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i,
				row.errExpected, err)
			continue
		}

		actualAddrsStr := make([]string, len(actualAddrs))
		for i, addr := range actualAddrs {
			actualAddrsStr[i] = addr.String()
		}
		expectedAddrsUpper := make([]string, len(row.addrs))
		for j, addr := range row.addrs {
			expectedAddrsUpper[j] = strings.ToUpper(addr)
		}

		require.Equal(t, parseMACs(row.addrs), expectedAddrsUpper,
			"row: %d, expected addrs: %s, actual addrs: %s", i,
			parseMACs(row.addrs), expectedAddrsUpper)
		require.NoError(t, err, "row: %d, error must be nil", i)
	}
}

type testCaseNames struct {
	names       []string // name list
	errExpected error    // expected error
}

func TestGetNames(t *testing.T) {
	mockWifi := mocks.NewWiFi(t)
	wifiService := MyWiFiService{WiFi: mockWifi}

	var testTable = []testCaseNames{
		{
			names:       []string{"wifi0", "wifi1"},
			errExpected: nil,
		},
		{
			names:       []string{},
			errExpected: errors.New("some error"),
		},
		{
			names:       []string{"eth0", "eth1", "eth2"},
			errExpected: nil,
		},
	}

	for i, row := range testTable {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfacesNames(row.names), row.errExpected)
		actualNames, err := wifiService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i,
				row.errExpected, err)
			continue
		}

		// Проверка результата
		require.Equal(t, row.names, actualNames,
			"row: %d, expected names: %v, actual names: %v", i,
			row.names, actualNames)

		require.NoError(t, err, "row: %d, error must be nil", i)
	}
}
