package event

import (
	"github.com/zhaohaijun/matrixchain/common"
	"github.com/zhaohaijun/matrixchain/core/types"
	"github.com/zhaohaijun/matrixchain/events"
	"github.com/zhaohaijun/matrixchain/events/message"
)

const (
	EVENT_LOG    = "Log"
	EVENT_NOTIFY = "Notify"
)

// PushSmartCodeEvent push event content to socket.io
func PushSmartCodeEvent(txHash common.Uint256, errcode int64, action string, result interface{}) {
	if events.DefActorPublisher == nil {
		return
	}
	smartCodeEvt := &types.SmartCodeEvent{
		TxHash: txHash,
		Action: action,
		Result: result,
		Error:  errcode,
	}
	events.DefActorPublisher.Publish(message.TOPIC_SMART_CODE_EVENT, &message.SmartCodeEventMsg{smartCodeEvt})
}
