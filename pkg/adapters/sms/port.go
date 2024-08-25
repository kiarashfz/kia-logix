package sms

type ISMSService interface {
	Send(phone string, template string) error
}
