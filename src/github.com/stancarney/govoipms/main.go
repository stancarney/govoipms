package main

import (
	"github.com/stancarney/govoipms/vclient"
	"log"
)

func main() {
	client := vclient.NewClient("https://voip.ms/api/v1/rest.php", "stan.voip@moohoffa.com", "PokemonGo1", true)
	general := vclient.NewGeneral(client)
	log.Println(general.GetBalance(true))
	log.Println(general.GetCountries(""))
	log.Println(general.GetIP())
	log.Println(general.GetLanguages(""))
	log.Println(general.GetServerInfo(""))
/*
	sub := &vclient.SubAccount{
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