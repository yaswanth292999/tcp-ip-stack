package main

import (
	"encoding/hex"
	"fmt"
)

//

/*
1. Make a class of type Ethernet Layer 2

2. Extract destination Mac

3. Extract Source Mac

4. Protocol Type

5. Payload

6. Checksum

TO DO:
CONVERT THIS INTO SPACE EFFICIENT MAPPING no need for strings you can store in may be some other like uint16 etc

*/

type EthernetHeader struct {
	destinationMacAddress string
	sourceMacAddress      string
	protocolType          string
	payload               []byte
}

// TODO :
// Investigate the extra bytes that are present in the packet array randomly even after packet size
// Construct arp payload and process the payload and send the response by writing it to the tap interface

func EthernetFrame(packet []byte, sizeOfPacket int) {

	fmt.Println(packet[:sizeOfPacket], "where")

	sizeOfDmac := 6
	sizeofSmac := 6
	sizeofProtocol := 2
	sizeofPayload := sizeOfPacket - (sizeOfDmac + sizeOfDmac + sizeofProtocol)
	fmt.Println(sizeOfPacket, sizeOfDmac, sizeofSmac, sizeofProtocol, sizeofPayload)
	ethernetHdr := EthernetHeader{
		destinationMacAddress: hex.EncodeToString(packet[0:6]),
		sourceMacAddress:      hex.EncodeToString(packet[6:12]),
		protocolType:          GetProtocolName(hex.EncodeToString(packet[12:14])),
		payload:               packet[14 : 14+(sizeOfPacket-14)],
	}

	if ethernetHdr.protocolType == "ARP" {
		ARP_RESPONSE(&ethernetHdr)
	}

}
