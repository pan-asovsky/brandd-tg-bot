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

	MarkSlotUnavailable = `
		UPDATE available_slots 
		SET is_available = false 
		WHERE date = $1 
		AND start_time = $2 
		AND end_time = $3`

	GetSlotByDateAndTime = `
		SELECT * FROM available_slots 
		WHERE date = $1 
		AND start_time = $2 
		AND end_time = $3`

	GetServiceTypeByCode = `
		SELECT * FROM service_types
		WHERE service_code = $1`

	FindActiveByChatID = `SELECT * FROM bookings WHERE chat_id = $1 AND is_active = true`

	SaveBooking = `INSERT INTO bookings (chat_id, slot_id, service_type_id,
                      					 rim_radius, is_active, 
                      					 status, created_at, updated_at)  
                   VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	SetPhoneByChatID = `UPDATE bookings SET user_phone = $1 WHERE chat_id = $2 and is_active = true`
)
