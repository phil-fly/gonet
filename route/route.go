package main

import (
	"github.com/vishvananda/netlink"
	"log"
	"net"
)

func addRoute(){
	link, err := netlink.LinkByName("br0")
	if err != nil {
		log.Fatal(err)
	}
	//目的子网
	dst := &net.IPNet{
		IP:   net.IPv4(5,5,5,0),
		Mask: net.CIDRMask(24, 32),
	}
	//网关
	gw := net.ParseIP("5.5.5.1").To4()

	//指定源
	src := net.ParseIP("5.5.5.122").To4()

	defaultRoute := netlink.Route{
		LinkIndex:link.Attrs().Index,
		Dst: dst,
		Gw:  gw,
		Src: src,
	}
	if err := netlink.RouteAdd(&defaultRoute); err != nil {
		log.Println(err.Error())
	}
}

func delRoute(){

}


func addDefaultRoute() {
	// Add default route
	gw := net.ParseIP("10.10.27.1").To4()

	defaultRoute := netlink.Route{
		Dst: nil,
		Gw:  gw,
	}

	if err := netlink.RouteAdd(&defaultRoute); err != nil {
		log.Fatal(err)
	}
}