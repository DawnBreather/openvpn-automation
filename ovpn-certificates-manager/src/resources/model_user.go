package resources

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"github.com/thoas/go-funk"
	"log"
	"regexp"
	"strings"
)

type Users []User

type User struct {
	Name string            `json:"Name,omitempty"`
	Email string           `json:"Email,omitempty"`
	VpnClients []VpnClient `json:"VpnClients,omitempty"`
	Blocked bool           `json:"Blocked,omitempty"`
}

func (u User) GenerateOvpnAndSendOverEmail(vc VpnClient){
	ovpn, err:= vc.GetOvpn()
	if err == nil {
		ok := Toolbox.Mailer.SendEmail(
			u.Email,
			fmt.Sprintf("%s.ovpn", vc.Name),
			ovpn)

		if ok {
			log.Printf("INFO: OVPN file generated and sent to: %s", u.Email)
		} else {
			log.Printf("ERROR: Error sending OVPN file to: %s", u.Email)
		}
	} else {
		log.Printf("ERROR: Error generating .ovpn file %s: \n%v", vc.Name, err)
	}
}

// Fetching Users from Auth0 proxy
func FetchUsers() *Users {
	client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").Get("http://localhost:3000/users")
	if err != nil {
		log.Printf("ERROR: Error fetching Users from Auth0 proxy:\n%v\n", err)
		return nil
	}

	var usrs Users

	//fmt.Println(string(resp.Body()))

	err = json.Unmarshal(resp.Body(), &usrs)
	if err != nil {
		log.Printf("ERROR: Error unmarshalling Users from JSON:\n%v\n", err)
		return nil
	}

	for i := range usrs{
		usrs[i].Name = strings.Split(usrs[i].Email, "@")[0]
	}

	resUsrs := funk.Filter(usrs, func(u User) bool {
		var r = regexp.MustCompile("^[a-zA-Z0-9]{1,20}[.][a-zA-Z0-9]{1,20}$")
		return r.MatchString(u.Name)
	}).([]User)

	res := Users(resUsrs)

	//fmt.Printf("%v", res)

	return & res
}

func (users *Users) RemoveUser(u User){
	index := funk.IndexOf(*users, u)
	copy((*users)[index:], (*users)[index + 1:])
	*users = (*users)[:len(*users) - 1]
}

func (users *Users) AddUser(u User){
	*users = append(*users, u)
}


func (usrs *Users) FetchVpnClients(_certsDir string, _ccds Ccds){

	certs := fetchCertificates(_certsDir)

	for i := range *usrs{

		name := (*usrs)[i].Name

		var vpnClients []VpnClient

		userCerts := funk.Filter(certs, func(c Cert) bool{
			return strings.HasPrefix(c.Name, name + "@") && strings.HasSuffix(c.Path, ".crt")
		})


		for _, uc := range userCerts.([]Cert){

			var resCcd *Ccd
			var resVcc *vpnClientCert

			cs := funk.Filter(_ccds, func(x Ccd) bool{
				return x.Name == uc.Name || x.Name == uc.Name + ".disabled"
			}).([]Ccd)

			if len(cs) == 0{
				log.Printf("WARNING: Unable to find CCDs for %s\n", uc.Name)
			} else {
				resCcd = &cs[0]
			}

			cts := funk.Filter(certs, func(c Cert) bool{
				return c.Name == uc.Name
			}).([]Cert)

			if len(cts) == 0{
				log.Printf("WARNING: Unable to find certificates for %s\n", uc.Name)
			} else {
				resVcc = &vpnClientCert{}
				for _, c := range cts {
					tmpC := c
					switch c.Type {
					case ".key":
						resVcc.Key = &tmpC
					case ".csr":
						resVcc.Csr = &tmpC
					case ".crt":
						resVcc.Crt = &tmpC
					}
				}
			}

			vc := VpnClient{
				Name:          uc.Name,
				Ccd:           resCcd,
				VpnClientCert: resVcc,
			}

			vpnClients = append(vpnClients, vc)
		}

		(*usrs)[i].VpnClients = vpnClients
	}

	//fmt.Printf("%v", *usrs)
}
