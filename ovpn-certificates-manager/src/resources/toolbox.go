package resources

var Toolbox toolbox = toolbox{
	Mailer: mailer{},
	Pki:    pki{},
	File: file{},
}

type toolbox struct {
	Mailer mailer
	Pki pki
	File file
}