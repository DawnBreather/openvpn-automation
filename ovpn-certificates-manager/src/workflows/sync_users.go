package workflows

import (
	"fmt"
	"github.com/thoas/go-funk"
	"log"
	"resources"
)



func SyncUsers() {

	log.Printf("#######################")
	log.Printf("# INFO: Synchronizing #")
	log.Printf("#######################")

	var _forDeletion resources.Users
	var _forBlocking resources.Users

	var _forUnblocking resources.Users
	var _forCreation resources.Users

	// stage 1. Fetch users from Auth0
	newUsers := *resources.FetchUsers()

		//create copy of newUsers array
	_forCreation = newUsers

	// stage 2. Identify changes for known users
	for i, u := range *resources.Session.Data.Users{
		newUsrs := funk.Filter(newUsers, func(usr resources.User) bool{
			return usr.Name == u.Name
		}).([]resources.User)

		if len(newUsrs) == 0 {
			log.Printf("INFO: User marked for deactivation (removal): %s", u.Name)
			_forDeletion = append(_forDeletion, u)
			continue
		}

		newUser := newUsrs[0]

		if u.Blocked != newUser.Blocked{
			if newUser.Blocked{
				_forBlocking = append(_forBlocking, u)
				log.Printf("INFO: User marked for deactivation (sync): %s", u.Name)
			} else {
				_forUnblocking = append(_forUnblocking, u)
				log.Printf("INFO: User marked for activation (sync): %s", u.Name)
			}
			(*resources.Session.Data.Users)[i].Blocked = !(*resources.Session.Data.Users)[i].Blocked
		} else {
			if u.Blocked && len(u.VpnClients) > 0 {
				_forBlocking = append(_forBlocking, u)
				log.Printf("INFO: User marked for deactivation (align): %s", u.Name)
			}

			if ! u.Blocked && len(u.VpnClients) == 0 {
				_forUnblocking = append(_forUnblocking, u)
				log.Printf("INFO: User marked for activation (align): %s", u.Name)
			}
		}

		//index := funk.IndexOf(_forCreation, newUser)
		//_forCreation = removeElementFromUsersSlice(_forCreation, index)
		_forCreation.RemoveUser(newUser)
	}

	for _, u := range _forCreation{
		log.Printf("INFO: User marked for activation (creation): %s\n", u.Name)
	}

	if len(_forCreation) == 0 && len(_forDeletion) == 0 && len(_forUnblocking) == 0 && len (_forBlocking) == 0{
		log.Printf("INFO: No actions required")
	}


	processBlockUsersQueue(_forBlocking)
	processDeleteUsersQueue(_forDeletion)
	processUnblockUsersQueue(_forUnblocking)
	processCreateUsersQueue(_forCreation)

}

func processCreateUsersQueue(_forCreation resources.Users){
	for _, u := range _forCreation{

		vc := resources.CreateVpnClient(fmt.Sprintf("%s@default", u.Name))
		if vc != nil {

			index := funk.IndexOf(*resources.Session.Data.Users, u)
			(*resources.Session.Data.Users)[index].VpnClients = []resources.VpnClient{*vc}

			resources.Session.Data.Users.AddUser(u)

			log.Printf("INFO: User created: %s", u.Name)
			u.GenerateOvpnAndSendOverEmail(*vc)
		} else {
			log.Printf("WARNING: User creation failed %s", u.Name)
		}

	}
}

func processUnblockUsersQueue(_forUnblocking resources.Users){
	for _, u := range _forUnblocking {

		vc := resources.CreateVpnClient(fmt.Sprintf("%s@default", u.Name))
		if vc != nil {

			index := funk.IndexOf(*resources.Session.Data.Users, u)
			(*resources.Session.Data.Users)[index].VpnClients = []resources.VpnClient{*vc}

			log.Printf("INFO: User activated: %s", u.Name)
			u.GenerateOvpnAndSendOverEmail(*vc)
		} else {
			log.Printf("WARNING: User activation failed %s", u.Name)
		}
	}
}

func processBlockUsersQueue(_forBlocking resources.Users){
	for _, u := range _forBlocking{
		for _, vc := range u.VpnClients {
			vc.Delete()
		}

		index := funk.IndexOf(*resources.Session.Data.Users, u)
		(*resources.Session.Data.Users)[index].VpnClients = nil

		log.Printf("INFO: User deactivated: %s", u.Name)
	}
}

func processDeleteUsersQueue(_forDeletion resources.Users){
	for _, u := range _forDeletion{
		for _, vc := range u.VpnClients {
			vc.Delete()
		}

		resources.Session.Data.Users.RemoveUser(u)

		log.Printf("INFO: User deleted: %s", u.Name)
	}
}



func removeElementFromVpnClientsSlice(slice []resources.VpnClient, i int) []resources.VpnClient {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func stage1_fetchUsersFromAuth0Proxy(){

}