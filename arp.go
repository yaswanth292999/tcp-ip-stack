package main

import "fmt"

func ARP_RESPONSE(ethernetHdr *EthernetHeader) {

	fmt.Println(ethernetHdr.destinationMacAddress, ethernetHdr.sourceMacAddress, "From ARP RESPONSE PACKET")
	fmt.Println(ethernetHdr.payload)
}
