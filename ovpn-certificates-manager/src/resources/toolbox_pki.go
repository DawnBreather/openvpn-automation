package resources

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type pki struct {

}

func (p pki) GenerateKeys(_name string) (crtFullPath, keyFullPath string, e error) {

	var _caCrtPath = Session.Config.VpnServer.SrvCertsPaths.CaCrt
	var _caKeyPath = Session.Config.VpnServer.SrvCertsPaths.CaKey
	var _certsDir = Session.Config.CertsDir
	var _pkiOrg = Session.Config.Pki.Organization

	// Load CA
	catls, err := tls.LoadX509KeyPair(
		_caCrtPath,
		_caKeyPath)

	if err != nil {
		panic(err)
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		panic(err)
	}

	// Prepare certificate
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{_pkiOrg},
			CommonName:   _name,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 4096)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	crtFullPath = filepath.Join(_certsDir, fmt.Sprintf("%s.crt", _name))
	certOut, err := os.Create(crtFullPath)
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	if err != nil{
		log.Printf("ERROR: Error writing generated .crt certificate %s:\n%v\n", _name, err)
		return "", "", err
	}
	certOut.Close()
	//log.Print("written cert.pem\n")

	// Private key
	keyFullPath = filepath.Join(_certsDir, fmt.Sprintf("%s.key", _name))
	keyOut, err := os.OpenFile(keyFullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	err = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	if err != nil {
		log.Printf("ERROR: Error writing generated .key certificate %s:\n%v\n", _name, err)
		return "", "", err
	}
	keyOut.Close()
	//log.Print("written key.pem\n")


	/*crtBt, err := ioutil.ReadFile(crtFullPath)
	log.Printf("ERROR: Error reading generated .crt certificate %s:\n%v\n", err)
	keyBt, err := ioutil.ReadFile(keyFullPath)
	log.Printf("ERROR: Error reading generated .key certificate %s:\n%v\n", err)

	crtContent = string(crtBt)
	keyContent = string(keyBt)*/

	return crtFullPath, keyFullPath, nil


}
