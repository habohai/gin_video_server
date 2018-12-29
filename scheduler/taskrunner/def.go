package taskrunner

const (
	// ReadyToDispatch controlChan中的消息 下发任务
	ReadyToDispatch = "d"
	// ReadyToExecute controlChan中的消息 执行任务
	ReadyToExecute = "e"
	// CLOSE controlChan中的消息 关闭消息
	CLOSE = "c"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error
