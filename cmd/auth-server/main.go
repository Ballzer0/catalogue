package main

import "github.com/Ballzer0/catalogue/pkg/authorizationserver"

func main() {
	// ### oauth2 server ###
	authorizationserver.RegisterHandlers() // the authorization server (fosite)
}
