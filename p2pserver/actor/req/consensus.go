package req

import "github.com/zhaohaijun/go-async-queue/actor"

var ConsensusPid *actor.PID

func SetConsensusPid(conPid *actor.PID) {
	ConsensusPid = conPid
}
