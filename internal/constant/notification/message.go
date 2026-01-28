package notification

const (
	NewBooking = "🟢 <b>Новая запись!</b>\n\n" +
		"📆 %s\n" +
		"🛞 Радиус: %s\n" +
		"🔧 Тип работ: %s\n" +
		"💰 Ориентировочная цена: %d\n" +
		"📞 Номер телефона: <code>%s</code>"

	CancelBooking = "🔴 <b>Клиент отменил запись!</b>\n\n" +
		"📆 %s\n" +
		"📞 Номер телефона: <code>%s</code>"

	CompleteBooking = "Ваша запись успешно завершена. Спасибо, что выбираете Bandd! 🙌"
)
