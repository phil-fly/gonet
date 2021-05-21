package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"github.com/vishvananda/netlink"
)

type Opt struct {
	MTU    int
	IPAddr string
	Name   string
}

const (
	// DefaultMTU is the default MTU for new bridge interfaces.
	DefaultMTU = 1500
)

var (
	// ErrNameEmpty holds the error for when the name is empty.
	ErrNameEmpty = errors.New("name cannot be empty")
)

// Init creates a bridge with the name specified if it does not exist.
func addbr(opt Opt) (*net.Interface, error) {

	if len(opt.Name) < 1 {
		return nil, ErrNameEmpty
	}

	// Set the defaults.
	if opt.MTU < 1 {
		opt.MTU = DefaultMTU
	}

	bridge, err := net.InterfaceByName(opt.Name)
	if err == nil {
		// Bridge already exists, return early.
		return bridge, nil
	}

	if !strings.Contains(err.Error(), "no such network interface") {
		return nil, fmt.Errorf("getting interface %s failed: %v", opt.Name, err)
	}

	// Create *netlink.Bridge object.
	la := netlink.NewLinkAttrs()
	la.Name = opt.Name
	la.MTU = opt.MTU
	br := &netlink.Bridge{LinkAttrs: la}
	if err := netlink.LinkAdd(br); err != nil {
		return nil, fmt.Errorf("bridge creation for %s failed: %v", opt.Name, err)
	}


	// Bring the bridge up.
	if err := netlink.LinkSetUp(br); err != nil {
		return nil, fmt.Errorf("bringing bridge %s up failed: %v", opt.Name, err)
	}

	return net.InterfaceByName(opt.Name)
}

// Delete removes the bridge by the specified name.
func delbr(opt Opt) error {
	// Get the link.
	l, err := netlink.LinkByName(opt.Name)
	if err != nil {
		return fmt.Errorf("getting bridge %s failed: %v", opt.Name, err)
	}

	// Delete the link.
	if err := netlink.LinkDel(l); err != nil {
		return fmt.Errorf("deleting bridge %s failed: %v", opt.Name, err)
	}

	return nil
}