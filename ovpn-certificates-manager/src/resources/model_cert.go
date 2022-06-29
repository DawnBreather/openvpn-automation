package resources

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//type Certs []Cert

type Cert struct {
	Content string
	Path string
	Name string
	Type string
}

func (c *Cert) Delete(){
	//if Toolbox.File.Exists(c.Path) {
	err := os.Remove(c.Path)
	if err != nil {
		log.Printf("ERROR: Error removing certificate %s:\n%v\n", c.Name, err)
	}
	//}
}

func parseCert(path string) (*Cert, error) {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("ERROR: Error reading certificate file:\n%v\n", err)
		return nil, err
	}

	return &Cert{
		Content: string(b),
		Path: path,
		Name: strings.ReplaceAll(filepath.Base(path), filepath.Ext(path), ""),
		Type: filepath.Ext(path),
	}, nil
}

func fetchCertificates(keysDir string) []Cert {

	var resCerts = []Cert{}

	files, err := ioutil.ReadDir(keysDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		c, err := parseCert(filepath.FromSlash(fmt.Sprintf("%s/%s", keysDir, f.Name())))
		if err == nil {
			resCerts = append(resCerts, *c)
		}
	}

	return resCerts
}