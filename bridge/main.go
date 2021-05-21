package main

import "log"

func main() {
	Addbr()
	//Delbr()
}

func Addbr() {
	br ,err := addbr(Opt{
		Name: "br0",
	})

	log.Println(br ,err)
}

func Delbr() {
	err := delbr(Opt{
		Name: "br0",
	})

	log.Println(err)
}