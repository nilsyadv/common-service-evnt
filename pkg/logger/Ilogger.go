package logger

type ILogger interface {
	Error(error)
	Panic(interface{})
	Println(interface{})
	Warning(interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Printf(string, ...interface{})
	Warningf(string, ...interface{})
}
