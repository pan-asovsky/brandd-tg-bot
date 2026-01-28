package interfaces

type PriceRepo interface {
	GetAllRimSizes() ([]string, error)
	GetSetPrice(svc string, radius string) (int64, error)
}
