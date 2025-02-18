package network_packet_analyzer

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers" // Add this import
	"github.com/google/gopacket/pcap"
)

func RunNetworkPacketAnalyzer() {
	// Get a list of all network devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print available devices
	fmt.Println("Available devices:")
	for _, device := range devices {
		fmt.Printf("\nName: %s\n", device.Name)
		fmt.Printf("Description: %s\n", device.Description)
		for _, address := range device.Addresses {
			fmt.Printf("- IP address: %s\n", address.IP)
			fmt.Printf("- Subnet mask: %s\n", address.Netmask)
		}
	}

	// Open device for capturing (replace "eth0" with your device name)
	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter for specific traffic (optional)
	err = handle.SetBPFFilter("tcp")
	if err != nil {
		log.Fatal(err)
	}

	// Create packet source
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	// Process packets
	for packet := range packetSource.Packets() {
		// Print timestamp
		fmt.Printf("\n%s\n", time.Now().Format(time.RFC3339))

		// Print packet info
		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip, _ := ipLayer.(*layers.IPv4)
			fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		}

		if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp, _ := tcpLayer.(*layers.TCP)
			fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
		}

		// Print application layer data if available
		if appLayer := packet.ApplicationLayer(); appLayer != nil {
			fmt.Printf("Payload: %s\n", string(appLayer.Payload()))
		}
	}
}
