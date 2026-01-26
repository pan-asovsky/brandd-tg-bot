package interfaces

type ConfigRepo interface {
	IsAutoConfirm() (bool, error)
}
