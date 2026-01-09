package interfaces

type ConfigService interface {
	IsAutoConfirm() (bool, error)
}
