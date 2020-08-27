package bitlog

// Logger declares the method set for a logging object in bitChat
type Logger interface {
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
}
