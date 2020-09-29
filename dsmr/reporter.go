package dsmr

type Reporter interface {
	Update(telegram Telegram)
	Log(msg string)
	Error(err error)
}
