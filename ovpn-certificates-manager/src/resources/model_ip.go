package resources

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

var _ipSubnet = ipSubnet{
	IsDefined: false,
}

type ipSubnet struct{
	A    int
	B    int
	C    int
	D    int

	Mask int

	IsDefined bool
}

type Ips map[Ip]struct{}
type Ip string

func (ips Ips) Append(ip string){
	ips[Ip(ip)] = struct{}{}
	log.Printf("INFO: added IP %s to the list", ip)
}

func (ips Ips) Delete(ip string){
	delete(ips, Ip(ip))
	log.Printf("INFO: removed IP %s from the list", ip)
}

func (ips Ips) Contains(ip string) bool{

	if net.ParseIP(ip) == nil{
		return true
	}

	if _, ok := ips[Ip(ip)]; ok {
		return true
	} else {
		return false
	}
}

// considering ipv4 address is a.b.c.d
func (ips Ips) GetAvailableIp() (ip string){

	switch {

	case _ipSubnet.Mask == 16:
		for ips.Contains(ip) {
			ip = generateIp()
		}
		return ip

	}

	return
}

func generateIp() (ip string){
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255

	var a = _ipSubnet.A
	var b = _ipSubnet.B
	var c = _ipSubnet.C
	if _ipSubnet.Mask <= 8 {
		b = rand.Intn(max - min + 1) + min
	}
	if _ipSubnet.Mask <= 16 {
		c = rand.Intn(max - min + 1) + min
	}

	var d = 0

	for d == 0 || d == 255 {
		d = rand.Intn(max - min + 1) + min
	}

	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}