package taskrunner

// Runner 任务运行类
type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

// NewRunner 初始化一个Runner
func NewRunner(size int, longLived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(controlChan, 1),
		Error:      make(controlChan, 1),
		Data:       make(dataChan, size),
		longLived:  longLived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			if c == ReadyToDispatch {
				if err := r.Dispatcher(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- ReadyToExecute
				}
			}

			if c == ReadyToExecute {
				if err := r.Executor(r.Data); err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- ReadyToDispatch
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
		default:
		}
	}
}

// StartAll 开始所有任务
func (r *Runner) StartAll() {
	r.Controller <- ReadyToDispatch
	r.startDispatch()
}
