package resources

type vpnServer struct {
	Name string
	Dn string
	Port string
	Certificates vpnServerCert
}

	type vpnServerCert struct {
		CaKey Cert
		CaCrt Cert
		TaKey Cert
	}

