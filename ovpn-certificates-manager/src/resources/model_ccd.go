package resources

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Ccds []Ccd

type Ccd struct {
	Name    string
	Content string
	fullPath string
}

func (c *Ccd) Delete(){
	ip, _ := c.getIp()
	Session.Data.Ips.Delete(ip)

	err := os.Remove(c.fullPath)
	if err != nil {
		log.Printf("ERROR: Error removing CCD %s:\n%v\n", c.Name, err)
	}
}

func fetchCcds(ccdsDirectoryPath string) Ccds{

	var resCcds = Ccds{}

	files, err := ioutil.ReadDir(ccdsDirectoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		c := parseCcd(filepath.FromSlash(fmt.Sprintf("%s/%s", ccdsDirectoryPath, f.Name())))
		if c != nil{
			resCcds = append(resCcds, *c)
		}
	}

	return resCcds
}

func (cs *Ccds) getIps() Ips{

	ips := Ips{}

	for _, c := range *cs{
		ip, _ := c.getIp()
		ips[Ip(ip)] = struct{}{}
	}

	return ips
}

func (c *Ccd) getKeyValues(key string) ([]string, int){

	keyRgexp := fmt.Sprintf("^%s.*", key)
	r := regexp.MustCompile(keyRgexp)

	keys := r.FindAllString(c.Content, -1)
	var resKeys []string

	for _, k := range keys{
		resKeys = append(resKeys, strings.TrimSpace(strings.ReplaceAll(k, key, "")))
	}

	return resKeys, len(resKeys)
}

func (c *Ccd) getIp() (ip, mask string){

	vals, qty := c.getKeyValues("ifconfig-push")

	if qty > 0 {
		ipAddr := strings.Split(vals[qty - 1], " ")
		ip = ipAddr[0]
		mask = ipAddr[1]
	} else {
		log.Printf("WARNING: No IP address found in CCD %s\n", c.Name)
	}

	// TODO implement proper subnet analyzer based on CCDs or define that in configuration instead of dynamic inistalization
	if ! _ipSubnet.IsDefined {
		_ipSubnet.A = 10
		_ipSubnet.B = 253
		_ipSubnet.C = 0
		_ipSubnet.D = 0

		_ipSubnet.Mask = 16

		_ipSubnet.IsDefined = true
	}

	return
}

func parseCcd(path string) *Ccd {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("ERROR: Error reading CCD file:\n%v\n", err)
		return nil
	}

	c := Ccd{
		Name: filepath.Base(path),
		Content: strings.ReplaceAll(string(content), "\n\n", ""),
		fullPath: path,
	}

	return &c
}