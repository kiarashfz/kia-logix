package sms

import "fmt"

type KaveNegarService struct {
	APIKey string
}

func NewKaveNegarService(APIKey string) ISMSService {
	return &KaveNegarService{APIKey: APIKey}
}

func (k *KaveNegarService) Send(phone string, template string) error {
	fmt.Printf("kavenegar sent an sms to: %s, with template %s\n", phone, template)
	return nil
}
