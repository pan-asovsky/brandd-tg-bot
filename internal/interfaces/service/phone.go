package service

type PhoneService interface {
	Normalize(phone string) (string, error)
	Detect(text string) (string, bool)
}
