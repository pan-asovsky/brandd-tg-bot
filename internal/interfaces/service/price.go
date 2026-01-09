package service

type PriceService interface {
	Calculate(service, radius string) (int64, error)
}
