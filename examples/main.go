package main

import (
	"github.com/stancarney/govoipms"
	"github.com/stancarney/govoipms/v1"
	"log"
)

func main() {
	v1c := govoipms.NewV1Client("https://voip.ms/api/v1/rest.php", "username", "password", true)
	
	a := &v1.Account{
		Username: "Test1",
		Protocol: "1",
		Description: "Description",
		AuthType: "1",
		Password: "Password1",
		IP: "",
		DeviceType: "2",
		CalleridNumber: "5555551234",
		CanadaRouting: "1",
		LockInternational: "1",
		InternationalRoute: "1",
		MusicOnHold: "default",
		AllowedCodecs: "ulaw;g729",
		DTMFMode: "auto",
		NAT: "yes",
		InternalExtension: "",
		InternalVoicemail: "",
		InternalDialtime: "20",
		ResellerClient: "0",
		ResellerPackage: "0",
		ResellerNextbilling: "0000-00-00",
	}
	
	account := v1c.NewAccountAPI()
	log.Println(account.CreateSubAccount(a))
	log.Println(a)

	a.Description = "New Description."
	log.Println(account.SetSubAccount(a))
	log.Println(a)

	log.Println(account.DelSubAccount(a.Id))
	
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
	//log.Println(account.GetSubAccounts(a.Account))
	//log.Println(account.SetSubAccount(""))

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
