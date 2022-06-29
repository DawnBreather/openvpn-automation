package resources

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"text/template"
)

// User - person from organization
// server - server from organization
// VpnClient - User-specific or server-specific vpn client implementation
// vpnServer - OpenVPN server implementation


type server struct {
	Name string
	VpnClients []VpnClient
}

type VpnClient struct{
	Name          string
	Ccd           *Ccd
	VpnClientCert *vpnClientCert
}

func (vc VpnClient) GetOvpn() (string, error){
	t, err := template.ParseFiles(Session.Config.VpnServer.ClientTemplatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, struct{
		ServerHost string
		ServerPort string
		Ca string
		Crt string
		Key string
		Ta string
	}{
		ServerHost: Session.Config.VpnServer.Server.Dn,
		ServerPort: Session.Config.VpnServer.Server.Port,
		Ca:         Session.Config.VpnServer.Server.Certificates.CaCrt.Content,
		Crt:        vc.VpnClientCert.Crt.Content,
		Key:        vc.VpnClientCert.Key.Content,
		Ta:         Session.Config.VpnServer.Server.Certificates.TaKey.Content,
	}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (vc VpnClient) Delete() {
	if vc.VpnClientCert.Crt != nil {
		vc.VpnClientCert.Crt.Delete()
	}
	if vc.VpnClientCert.Key != nil {
		vc.VpnClientCert.Key.Delete()
	}
	if vc.VpnClientCert.Csr != nil {
		vc.VpnClientCert.Csr.Delete()
	}

	if vc.Ccd != nil {
		vc.Ccd.Delete()
	}
}

	type vpnClientCert struct {
		Csr *Cert
		Key *Cert
		Crt *Cert
	}

func CreateVpnClient(name string) *VpnClient{

	var vcc *vpnClientCert
	var ccd *Ccd

	vcc = createVpnClientCert(name)
	if vcc != nil {
		ccd = createCcd(name)
		if ccd == nil {
			return nil
		}
	} else {
		return nil
	}

	return &VpnClient{
		Name: name,
		Ccd: createCcd(name),
		VpnClientCert: createVpnClientCert(name),
	}
}

	func createVpnClientCert(name string) *vpnClientCert{
		crtFullPath, keyFullPath, err := Toolbox.Pki.GenerateKeys(name)

		if err != nil {
			return nil
		}

		key, _ := parseCert(keyFullPath)
		crt, _ := parseCert(crtFullPath)

		return &vpnClientCert{
			Key: key,
			Crt: crt,
		}
	}

	func createCcd(name string) *Ccd{

		fullPath := filepath.FromSlash(fmt.Sprintf("%s/%s", Session.Config.CcdDir, name))

		if !Toolbox.File.Exists(fullPath) {

			ip := Session.Data.Ips.GetAvailableIp()
			content := fmt.Sprintf("ifconfig-push %s %s\n", ip, Session.Config.Net.mask)

			err := ioutil.WriteFile(fullPath, []byte(content), 0x644)
			if err != nil {
				log.Printf("ERROR: Error saving generated CCD %n:\n%v\n", name, err)
				return nil
			}

			Session.Data.Ips.Append(ip)

			return &Ccd{
				Name:     name,
				Content:  content,
				fullPath: fullPath,
			}
		}

		return parseCcd(fullPath)
	}


//type vpnClients map[VpnClient]struct{}