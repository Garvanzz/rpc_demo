package task

import "rpc_demo/common/config"

type PushMessage struct {
}

var pushChan []chan *PushMessage

type Task struct{}

func New() *Task {
	return new(Task)
}

func (t *Task) Run() {
	t.goPush()
}

func (t *Task) goPush() {
	pushChan = make([]chan *PushMessage, config.Conf.Task.PushChan)

	// for range  如果用range 应该也可以遍历
	for i := 0; i < len(pushChan); i++ {
		pushChan[i] = make(chan *PushMessage, config.Conf.Task.PushChanSize)

		go goPush(pushChan[i])
	}
}

func goPush(channel chan *PushMessage) {
	for {
		msg := <-channel

		// TODO:
	}
}
