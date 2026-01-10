package entity

type ServiceType struct {
	ID          int64  `db:"id"`
	ServiceCode string `db:"service_code"`
	ServiceName string `db:"service_name"`
	IsComposite bool   `db:"is_composite"`
}
