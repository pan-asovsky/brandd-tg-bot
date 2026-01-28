package service

type ConfigService interface {
	IsAutoConfirm() (bool, error)
}
