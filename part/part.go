package part

type Part interface {
	Magic() uint32
	Init() error
	Startup() error
	Shutdown() error
}
