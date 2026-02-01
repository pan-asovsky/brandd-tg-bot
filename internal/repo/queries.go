package repo

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

	GetCompositeServiceTypes = `SELECT * FROM service_types WHERE is_composite = true`
	GetAllRimSizes           = `SELECT DISTINCT rim_size FROM prices`
	IsAutoConfirm            = `SELECT auto_confirm from config`

	MarkSlotUnavailable = `
		UPDATE available_slots 
		SET is_available = false 
		WHERE date = $1 
		AND start_time = $2
		`

	FindActiveNotPending = `SELECT * FROM bookings WHERE chat_id = $1 AND is_active = true AND status != 'PENDING'`
	FindActivePending    = `SELECT * FROM bookings WHERE chat_id = $1 AND is_active = true AND status = 'PENDING'`
	FindAnyActive        = `SELECT * FROM bookings WHERE chat_id = $1 and is_active = true`
	FindAllActive        = `SELECT * FROM bookings WHERE is_active = true`

	StatusesByPeriod = `
		SELECT
			COUNT(*) FILTER (WHERE status = 'CONFIRMED') AS active_count,
			COUNT(*) FILTER (WHERE status = 'CANCELLED') AS cancelled_count,
			COUNT(*) FILTER (WHERE status = 'COMPLETED') AS completed_count,
			COUNT(*) FILTER (WHERE status = 'NO_SHOW') AS no_show_count,
			COUNT(*) FILTER (WHERE status = 'PENDING') AS pending_count
		FROM bookings
		WHERE date >= $1 AND date <= $2;
`

	FindByID = `SELECT * FROM bookings WHERE id = $1`

	BookingExists = `SELECT EXISTS(SELECT 1 FROM bookings WHERE chat_id = $1 AND is_active = true AND status NOT IN ('CANCELLED', 'NO_SHOW'))`

	SaveBooking = `INSERT INTO bookings (chat_id, date, time,
                      					 service, rim_radius, total_price,
                      					 is_active, status, created_at, updated_at)  
                   VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	SetPhone = `UPDATE bookings SET user_phone = $1 WHERE chat_id = $2 and is_active = true`

	UpdateRimRadius = `UPDATE bookings SET rim_radius = $1 WHERE chat_id = $2 and is_active = true`

	UpdateStatus = `UPDATE bookings SET status = $1 WHERE chat_id = $2 AND is_active = true`

	UpdatePrice = `UPDATE bookings SET total_price = $1 WHERE chat_id = $2 AND is_active = true`

	UpdateService = `UPDATE bookings SET service = $1 WHERE chat_id = $2 AND is_active = true`

	ConfirmBooking = `UPDATE bookings SET status = $1, confirmed_by = $2 WHERE chat_id = $3 and is_active = true`

	CancelBooking = `UPDATE bookings SET status = $1, is_active = false WHERE chat_id = $2`

	Close = `UPDATE bookings SET status = $1, is_active = false 
                WHERE chat_id = $2 AND id = $3
                RETURNING id, chat_id,user_phone, date, time, service, rim_radius, total_price, 
                    status, is_active, created_at, updated_at, confirmed_by, cancelled_by, notes;
	`

	GetPricePerSet = `SELECT price_per_set FROM prices WHERE service_type_code = $1 AND rim_size = $2 AND is_active = true`

	FreeUpSlot = `UPDATE available_slots SET is_available = true WHERE date = $1 AND start_time = $2`

	GetActiveAdmins = "SELECT * FROM users WHERE role = 'admin' AND is_active = true"

	GetUserRole = "SELECT role FROM users WHERE chat_id = $1"
)
