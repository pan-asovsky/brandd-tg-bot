package model

type ServiceType struct {
	ID          int64  `db:"id"`
	ServiceCode string `db:"service_code"`
	ServiceName string `db:"service_name"`
	Description string `db:"description"`
}

const (
	ServiceCodeTireChange = "TIRE_SERVICE"
	ServiceCodeBalancing  = "BALANCING"
	ServiceCodeComplex    = "COMPLEX"
	ServiceCodeService    = "SERVICE"
)
