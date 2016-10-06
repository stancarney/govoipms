package main

import (
	"github.com/stancarney/govoipms"
	"log"
)

func main() {
	v1 := govoipms.NewV1Client("https://voip.ms/api/v1/rest.php", "stan.voip@moohoffa.com", "PokemonGo1", true)
	/*
	general := govoipms.NewGeneralAPI(client)
	log.Println(general.GetBalance(true))
	log.Println(general.GetCountries(""))
	log.Println(general.GetIP())
	log.Println(general.GetLanguages(""))
	log.Println(general.GetServerInfo(""))
	*/

	account := v1.NewAccountAPI()
	//log.Println(account.GetAllowedCodecs(""))
	//log.Println(account.GetAuthTypes(0))
	//log.Println(account.GetDeviceTypes(0))
	//log.Println(account.GetDTMFModes(""))
	//log.Println(account.GetDTMFModes(""))
	//log.Println(account.GetLockInternational(""))
	//log.Println(account.GetMusicOnHold(""))
	//log.Println(account.GetMusicOnHold(""))
	//log.Println(account.GetNAT(""))
	//log.Println(account.GetProtocols(0))
	//log.Println(account.GetRegistrationStatus("100000_a"))
	//log.Println(account.GetReportEstimatedHoldTime(""))
	//log.Println(account.GetRoutes(0))
	log.Println(account.GetSubAccounts(""))

/*
	sub := &govoipms.SubAccount{
		Username: "test",
		Protocol: "3",
		Description: "Desc Test",
		AuthType: "2",
		Password: "password",
		IP: "127.0.0.1",
		DeviceType: "1",
		CalleridNumber: "14035551234",
		CanadaRouting: "1",
		LockInternational: "1",
		InternationalRoute: "1",
		MusicOnHold: "jazz",
		AllowedCodecs: "ulaw",
		DTMFMode: "inband",
		NAT: "route",
		InternalExtension: "999",
		InternalVoicemail: "999",
		InternalDialtime: "60",
		//ResellerClient
		//ResellerPackage
		//ResellerNextBilling
		//ResellerChargesetup
	}
	log.Println(general.CreateSubAccount(sub))
	*/
}
