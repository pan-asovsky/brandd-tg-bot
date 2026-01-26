package notification

type Recipient struct {
	ChatID int64
}

type Formatter func(data any) (string, error)
