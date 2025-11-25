package config

func (config *Config) IsAdmin(chatID int64) bool {
	for _, adminId := range config.AdminChatID {
		if adminId == chatID {
			return true
		}
	}
	return false
}

func (config *Config) GetAdminChatID() []int64 {
	return config.AdminChatID
}

func (config *Config) AddAdmin(chatID int64) {
	for _, existingId := range config.AdminChatID {
		if existingId == chatID {
			return // Уже существует
		}
	}
	config.AdminChatID = append(config.AdminChatID, chatID)
}
