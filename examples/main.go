package main

import (
	"github.com/stancarney/govoipms"
	"github.com/stancarney/govoipms/v1"
	"log"
	"time"
)

func main() {
	v1c := govoipms.NewV1Client("https://voip.ms/api/v1/rest.php", "email", "passwd", true)
	//GeneralFunctions(v1c)
	//AccountFunctions(v1c)
	//CDRFunctions(v1c) 
	//ClientFunctions(v1c)
	DIDFunctions(v1c)
}

func GeneralFunctions(v1c *v1.VOIPClient) {

	general := v1c.NewGeneralAPI()

	log.Println(general.GetBalance(true))
	//log.Println(general.GetTransactionHistory(time.Now().Add(time.Hour * -8760), time.Now()))
}

func AccountFunctions(v1c *v1.VOIPClient) {
	accounts := v1c.NewAccountsAPI()

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

	log.Println(accounts.CreateSubAccount(a))
	log.Println(a)

	a.Description = "New Description."
	log.Println(accounts.SetSubAccount(a))
	log.Println(a)

	log.Println(accounts.DelSubAccount(a.Id))

	log.Println(accounts.GetAllowedCodecs(""))
	log.Println(accounts.GetAuthTypes(0))
	log.Println(accounts.GetDeviceTypes(0))
	log.Println(accounts.GetDTMFModes(""))
	log.Println(accounts.GetDTMFModes(""))
	log.Println(accounts.GetLockInternational(""))
	log.Println(accounts.GetMusicOnHold(""))
	log.Println(accounts.GetMusicOnHold(""))
	log.Println(accounts.GetNAT(""))
	log.Println(accounts.GetProtocols(0))
	log.Println(accounts.GetRegistrationStatus("100000_a"))
	log.Println(accounts.GetReportEstimatedHoldTime(""))
	log.Println(accounts.GetRoutes(0))
	log.Println(accounts.GetSubAccounts(a.Account))

	sub := &v1.Account{
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

	log.Println(accounts.SetSubAccount(sub))
}

func CDRFunctions(v1c *v1.VOIPClient) {
	cdr := v1c.NewCDRAPI()
	mst, _ := time.LoadLocation("America/Edmonton")

	//log.Println(cdr.GetCallAccounts(""))
	//log.Println(cdr.GetCallBilling())

	//log.Println(cdr.GetCallAccounts(""))
	//log.Println(cdr.GetCallBilling())
	//log.Println(cdr.GetCallTypes(""))
	cs := v1.CallStatus{true, false, false, false}
	//log.Println(cdr.GetCDR(time.Now().Add(time.Hour * -300), time.Now(), cs, mst, "all", "", ""))
	//log.Println(cdr.GetRates("1", "Canada"))
	log.Println(cdr.GetResellerCDR(time.Now().Add(time.Hour * -300), time.Now(), "12345", cs, mst, "all", "", ""))
}

func ClientFunctions(v1c *v1.VOIPClient) {
	client := v1c.NewClientsAPI()
	//log.Println(client.GetPackages(""))
	//log.Println(client.AddCharge("589758", "Test_Charge", 0.01, true))
	//log.Println(client.AddPayment("589758", "Test_Payment", 0.01, true))
	//log.Println(client.GetBalanceManagement("0"))
	//log.Println(client.GetCharges("589758"))
	//log.Println(client.GetClientPackages("589758"))
	//log.Println(client.GetClients(""))
	//log.Println(client.GetClientThreshold("589758"))
	//log.Println(client.GetDeposits("589758"))
	//log.Println(client.GetResellerBalance("589758"))


	setClient := &v1.Client{
		Client: "589758",
		Email: "stan.testvoipms@moohoffa.com",
		Password: "P@ssw0rd!",
		Company: "New2Company",
		FirstName: "Stan",
		LastName: "Updated",
		Address: "",
		City: "",
		State: "",
		Country: "",
		Zip: "",
		PhoneNumber: "4035920968",
		BalanceManagement: "",
	}
	log.Println(client.SetClient(setClient))

	//log.Println(client.SetClientThreshold("589758", "5", "stan.voipmstest3threshold@moohoffa.com"))

	/*
	signupClient := &v1.Client{
		Email: "stan.test2voipms@moohoffa.com",
		Password: "P@ssw0rd!",
		Company: "",
		FirstName: "FirstNew",
		LastName: "LastNew",
		Address: "123 Fake St.",
		City: "Calgary",
		State: "AB",
		Country: "CA",
		Zip: "",
		PhoneNumber: "4035920968",
		BalanceManagement: "",
	}
	log.Println(client.SignupClient(signupClient, "stan.test2voipms@moohoffa.com", "P@ssw0rd!", false))
	*/
}

func DIDFunctions(v1c *v1.VOIPClient) {
	did := v1c.NewDIDsAPI()
	/*
	boUSA := &v1.BackOrder{
		Quantity: 1,
		State: "CA", //For US Only
		//Province: "AB", //For Canada Only.
		Ratecenter: "GRENADA",
		Routing: v1.AccountRoute("1234"),
		FailoverBusy: "none:",
		FailoverUnreachable: "none:",
		FailoverNoanswer: "none:",
		Voicemail: "101",
		POP: "1",
		Dialtime: 60,
		CNAM: true,
		CalleridPrefix: "Govoipms",
		Note: "This is a test",
		BillingType: 2,
		Test: true,
	}

	log.Println(did.BackOrderDIDUSA(bo))
	*/

	boCAN := &v1.BackOrder{
		Quantity: 1,
		//State: "CA", //For US Only
		Province: "AB", //For Canada Only.
		Ratecenter: "1",
		Routing: v1.NewAccountRoute("1234"),
		FailoverBusy: v1.NewNoneRoute(),
		FailoverUnreachable: v1.NewNoneRoute(),
		FailoverNoanswer: v1.NewNoneRoute(),
		Voicemail: "101",
		POP: "1",
		Dialtime: 60,
		CNAM: true,
		CalleridPrefix: "Govoipms",
		Note: "This is a test",
		BillingType: 2,
		Test: true,
	}

	log.Println(did.BackOrderDIDCAN(boCAN))

	log.Println(did.GetRateCentersCAN("AB"))
	//log.Println(did.GetRateCentersUSA("CA"))
	//log.Println(did.GetStates())
}
