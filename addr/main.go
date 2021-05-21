package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	AddAddr("br0","5.5.5.122","255.255.255.0")
	//DelAddr("br0","5.5.5.122","255.255.255.0")
}

func AddAddr(nicName,ipaddr,maskAddr string) {
	addr := net.ParseIP(maskAddr).To4()

	sz, _ := net.IPv4Mask(addr[0], addr[1], addr[2], addr[3]).Size()

	err := addAddr(Opt{
		Name: nicName,
		IPAddr: fmt.Sprintf("%s/%d",ipaddr,sz),
	})

	log.Println(err)
}

func DelAddr(nicName,ipaddr,maskAddr string) {
	addr := net.ParseIP(maskAddr).To4()

	sz, _ := net.IPv4Mask(addr[0], addr[1], addr[2], addr[3]).Size()

	err := delAddr(Opt{
		Name: nicName,
		IPAddr: fmt.Sprintf("%s/%d",ipaddr,sz),
	})

	log.Println(err)
}