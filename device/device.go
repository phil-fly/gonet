package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func main(){

}

//获取linux主机物理网卡
func getNicbridge(){
	netdevicePath := "/sys/class/net/"
	netdevices, _ := ioutil.ReadDir(netdevicePath)
	for _, f := range netdevices {
		filelink, err := os.Readlink(fmt.Sprintf("%s%s",netdevicePath,f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		//过率虚拟网卡
		if strings.Contains(filelink,"virtual") {
			continue
		}
		fmt.Println(f.Name())
	}
}

//获取linux主机物理网卡
func getNiclist(){
	netdevicePath := "/sys/class/net/"
	netdevices, _ := ioutil.ReadDir(netdevicePath)
	for _, f := range netdevices {
		filelink, err := os.Readlink(fmt.Sprintf("%s%s",netdevicePath,f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		//过率虚拟网卡
		if strings.Contains(filelink,"virtual") {
			continue
		}
		fmt.Println(f.Name())
	}
}