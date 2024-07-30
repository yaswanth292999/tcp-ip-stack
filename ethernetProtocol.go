package main

import "strings"

func GetProtocolName(ethType string) string {

	ethType = strings.ToUpper(ethType)

	switch ethType {
	case "0800":
		return "IPv4"
	case "0806":
		return "ARP"
	case "0842":
		return "Wake-on-LAN"
	case "22F3":
		return "IETF TRILL Protocol"
	case "6003":
		return "DECnet Phase IV"
	case "8035":
		return "Reverse ARP"
	case "809B":
		return "AppleTalk"
	case "80F3":
		return "AppleTalk ARP"
	case "8100":
		return "VLAN-tagged frame (IEEE 802.1Q)"
	case "8137":
		return "IPX"
	case "8204":
		return "QNX Qnet"
	case "86DD":
		return "IPv6"
	case "8808":
		return "Ethernet flow control"
	case "8809":
		return "Ethernet Slow Protocols"
	case "8819":
		return "CobraNet"
	case "8847":
		return "MPLS unicast"
	case "8848":
		return "MPLS multicast"
	case "8863":
		return "PPPoE Discovery Stage"
	case "8864":
		return "PPPoE Session Stage"
	case "8870":
		return "Jumbo frames (Obsoleted draft-ietf-isis-ext-eth-01)"
	case "887B":
		return "HomePlug 1.0 MME"
	case "888E":
		return "EAP over LAN (IEEE 802.1X)"
	case "8892":
		return "PROFINET Protocol"
	case "889A":
		return "HyperSCSI (SCSI over Ethernet)"
	case "88A2":
		return "ATA over Ethernet"
	case "88A4":
		return "EtherCAT Protocol"
	case "88A8":
		return "Provider Bridging (IEEE 802.1ad)"
	case "88AB":
		return "Ethernet Powerlink"
	case "88CC":
		return "LLDP"
	case "88CD":
		return "SERCOS III"
	case "88E1":
		return "HomePlug AV MME"
	case "88E3":
		return "Media Redundancy Protocol (IEC62439-2)"
	case "88E5":
		return "MAC security (IEEE 802.1AE)"
	case "88E7":
		return "Provider Backbone Bridges (PBB) (IEEE 802.1ah)"
	case "88F7":
		return "Precision Time Protocol (IEEE 1588)"
	case "8902":
		return "IEEE 802.1ag Connectivity Fault Management (CFM) Protocol"
	case "8906":
		return "Fibre Channel over Ethernet (FCoE)"
	case "8914":
		return "FCoE Initialization Protocol"
	case "8915":
		return "RDMA over Converged Ethernet (RoCE)"
	case "891D":
		return "TTEthernet Protocol Control Frame (TTE)"
	case "892F":
		return "High-availability Seamless Redundancy (HSR)"
	case "9000":
		return "Ethernet Configuration Testing Protocol"
	default:
		return "UNKNOWN ETHER TYPE IN ETHERNET FRAME"
	}

}
