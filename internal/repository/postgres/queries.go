package postgres

const (
	IsTodayAvailable = `
		SELECT COUNT(*) > 0 FROM available_slots
		WHERE date = CURRENT_DATE
		AND start_time > NOW()::time
		AND is_available = true
	`
	GetZonesByDate = `
		SELECT * FROM available_slots
		WHERE date = $1
  		AND (date > CURRENT_DATE OR (date = CURRENT_DATE AND start_time > NOW()::time))
  		AND is_available = true
		ORDER BY start_time
	`
	GetAllServiceTypes = `SELECT * FROM service_types`
	GetAllRimSizes     = `SELECT DISTINCT rim_size FROM prices`
	IsAutoConfirm      = `SELECT auto_confirm from config`
)
