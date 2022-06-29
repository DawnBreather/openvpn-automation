package resources

import (
	"bytes"
	"fmt"
	"github.com/domodwyer/mailyak"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"
	"time"
)

var Session = session{
	Config: config{

		Application: applicationConfig{
			SyncUsersIntervalSeconds: 60,
		},

		CcdDir:   "vpn-resources/ccd",
		CertsDir: "vpn-resources/keys",

		VpnServer: serverConfig{
			ClientTemplatePath: "templates/ovpn-client-config.tmpl",
			Server:        vpnServer{
				Name:         "default",
				Dn:           "us-nodes.vpn.scopicsoftware.com",
				Port:         "1194",
			},
			SrvCertsPaths: srvCertsPaths{
				CaCrt: "vpn-resources/srv-keys/rootCA.crt",
				CaKey: "vpn-resources/srv-keys/rootCA.key",
				TaKey: "vpn-resources/srv-keys/ta.key",
			},
		},

		Pki:pkiConfig{
			Organization: "Scopic Software LLC",
		},

		Net: netConfig{
			mask: "255.255.0.0",
		},
		Mailer: mailerConfig{
			Smtp: smtpConfig{
				host:     "smtp.zoho.com",
				port:     "587",
				Username: "noreply@scopicsoftware.com",
				password: "qNg@UX83$",
			},
			EmailMessage: emailMessage{
				templateFilePath: filepath.Join("templates", "email-template.tmpl"),
				Subject:          "[VPN] Welcome to Scopic!",
			},
		},
	},
	Switcher:switcher{
		sendEmail: true,
	},
}

type session struct{
	Config config
	Data data
	Switcher switcher
}

type switcher struct {
	sendEmail bool
}

type config struct {
	Application applicationConfig
	CcdDir    string
	CertsDir  string
	Net       netConfig
	Mailer    mailerConfig
	VpnServer serverConfig
	Pki pkiConfig
}

	type applicationConfig struct{
			SyncUsersIntervalSeconds time.Duration
	}

	type pkiConfig struct {
		Organization string
	}

	type serverConfig struct {
		Server vpnServer
		ClientTemplatePath string
		SrvCertsPaths srvCertsPaths
	}
		type srvCertsPaths struct {
			CaCrt string
			CaKey string
			TaKey string
		}

		func (sc *serverConfig) Initialize() {

			sc.Server.Name = ""

			caCrt, err := parseCert(sc.SrvCertsPaths.CaCrt)
			if err != nil {
				log.Fatalf("ERROR: Error reading CA public certificate:\n%v", err)
			}

			caKey , err := parseCert(sc.SrvCertsPaths.CaKey)
			if err != nil {
				log.Fatalf("ERROR: Error reading CA private certificate:\n%v", err)
			}

			taKey, err := parseCert(sc.SrvCertsPaths.TaKey)
			if err != nil {
				log.Fatalf("ERROR: Error reading TA private certificate:\n%v", err)
			}

			sc.Server.Certificates.CaCrt = *caCrt
			sc.Server.Certificates.CaKey = *caKey
			sc.Server.Certificates.TaKey = *taKey
		}

	type mailerConfig struct {
		Smtp smtpConfig
		EmailMessage emailMessage
	}

		type smtpConfig struct {
			host     string
			port     string
			Username string
			password string
		}
			func (s smtpConfig) GetMailAgent() *mailyak.MailYak{
				return mailyak.New(
					fmt.Sprintf("%s:%s", s.host, s.port),
					smtp.PlainAuth("", s.Username, s.password, s.host))
			}

		type emailMessage struct {
			templateFilePath string
			Subject          string
		}

			func (em emailMessage) ParseTemplate(data interface{}) (error, string) {
				t, err := template.ParseFiles(em.templateFilePath)
				if err != nil {
					return err, ""
				}
				buf := new(bytes.Buffer)
				if err = t.Execute(buf, data); err != nil {
					return err, ""
				}
				return nil, buf.String()
			}

	type netConfig struct {
		mask string
	}

	type data struct{
		Ips *Ips
		Users *Users
		//Ccds *Ccds
		VpnServerCert *vpnServerCert
	}



func (s *session) Initalize(){

	log.Printf("INFO: Initializing session")

	s.Config.VpnServer.Initialize()

	ccds := fetchCcds(s.Config.CcdDir)

	ips := ccds.getIps()
	s.Data.Ips = &ips

	users := FetchUsers()
	s.Data.Users = users

	s.Data.Users.FetchVpnClients(s.Config.CertsDir, ccds)
}