package postgres

const (
	IsTodayAvailable = `
		SELECT COUNT(*) > 0 FROM available_slots
		WHERE date = CURRENT_DATE
		AND start_time > NOW()::time
		AND is_available = true
	`
	FindByDate = `
		SELECT DISTINCT zone
		FROM (
			SELECT CASE
				WHEN start_time >= '09:00' AND start_time < '12:00' THEN '09:00-12:00'
				WHEN start_time >= '12:00' AND start_time < '15:00' THEN '12:00-15:00'
				WHEN start_time >= '15:00' AND start_time < '18:00' THEN '15:00-18:00'
				WHEN start_time >= '18:00' AND start_time < '21:00' THEN '18:00-21:00'
			END AS zone
			FROM available_slots
			WHERE date = $1
			AND (date > CURRENT_DATE OR (date = CURRENT_DATE AND start_time > NOW()::time))
			AND is_available = true
		) t
		WHERE zone IS NOT NULL
		ORDER BY zone
	`
	GetZonesByDate = `
		SELECT * FROM available_slots
		WHERE date = $1
  		AND (date > CURRENT_DATE OR (date = CURRENT_DATE AND start_time > NOW()::time))
  		AND is_available = true
		ORDER BY start_time
	`
)
