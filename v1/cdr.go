package v1

type CDRAPI struct {
	client *VOIPClient
}

func (c *CDRAPI) GetCallAccount(clientId string) error {
	panic("Not implemented!")
}