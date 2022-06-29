package main

import (
	"log"
	"resources"
	"time"
	"workflows"
)


func main(){

	resources.Session.Initalize()

	for {
		workflows.SyncUsers()
		log.Printf("INFO: %d seconds interval", resources.Session.Config.Application.SyncUsersIntervalSeconds)
		time.Sleep(resources.Session.Config.Application.SyncUsersIntervalSeconds * time.Second)
	}
}