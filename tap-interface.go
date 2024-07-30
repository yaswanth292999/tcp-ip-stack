package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	TUNSETIFF = 0x400454ca // ioctl command to set the interface
	IFF_TAP   = 0x0002     // TAP interface type
	IFF_NO_PI = 0x1000     // No packet information
	IFNAMSIZ  = 16         // Maximum interface name size
)

// ifreq structure used for ioctl calls
type ifreq struct {
	Name  [IFNAMSIZ]byte
	Flags int32
}

func main() {
	// Create the TAP interface
	count := 1
	tapName := "tap0"
	tap, err := createTapInterface(tapName)
	if err != nil {
		fmt.Println("Error creating TAP interface:", err)
		return
	}

	// Read packets from the TAP interface

	fmt.Println("Read packets from TAP interface before for loop")
	packet := make([]byte, 1500)
	for {
		fmt.Println("In for loop for...")
		fmt.Println(count)
		count = count + 1
		sizeOfPacket, err := tap.Read(packet)
		if err != nil {
			fmt.Println("Error reading from TAP interface:", err)
			return
		}

		fmt.Printf("Read %d bytes: %x\n", sizeOfPacket, packet[:sizeOfPacket])
		EthernetFrame(packet, sizeOfPacket)
	}
}

func createTapInterface(name string) (*os.File, error) {
	// Open the TUN/TAP device
	fd, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}

	// Prepare the ifreq structure
	var ifr ifreq
	copy(ifr.Name[:], name)
	ifr.Flags = IFF_TAP | IFF_NO_PI

	// Perform the ioctl request
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
	if errno != 0 {
		fd.Close()
		return nil, fmt.Errorf("ioctl failed: %v", errno)
	}

	return fd, nil
}
