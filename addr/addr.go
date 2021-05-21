package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

type Opt struct {
	IPAddr string
	Name   string
}

func addAddr(opt Opt) (error) {
	iface, err := netlink.LinkByName(opt.Name)

	// Setup ip address for device.
	addr, err := netlink.ParseAddr(opt.IPAddr)
	if err != nil {
		return fmt.Errorf("parsing address %s failed: %v", opt.IPAddr, err)
	}
	if err := netlink.AddrAdd(iface, addr); err != nil {
		return fmt.Errorf("adding address %s to Device %s failed: %v", addr.String(), opt.Name, err)
	}
	return nil
}

// Delete removes the Device by the specified name.
func delAddr(opt Opt) (error) {
	// Get the link.
	l, err := netlink.LinkByName(opt.Name)
	if err != nil {
		return fmt.Errorf("getting Device %s failed: %v", opt.Name, err)
	}

	addr, err := netlink.ParseAddr(opt.IPAddr)
	if err != nil {
		return fmt.Errorf("parsing address %s failed: %v", opt.IPAddr, err)
	}
	// Delete the link.
	if err := netlink.AddrDel(l,addr); err != nil {
		return fmt.Errorf("deleting Device %s failed: %v", opt.Name, err)
	}

	return nil
}