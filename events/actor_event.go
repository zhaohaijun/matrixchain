package events

import (
	"fmt"

	"github.com/zhaohaijun/go-async-queue/actor"
	"github.com/zhaohaijun/go-async-queue/eventhub"
)

var DefEvtHub *eventhub.EventHub
var DefPublisherPID *actor.PID
var DefActorPublisher *ActorPublisher
var defPublisherProps *actor.Props

func Init() {
	DefEvtHub = eventhub.GlobalEventHub
	defPublisherProps = actor.FromFunc(func(context actor.Context) {})
	var err error
	DefPublisherPID, err = actor.SpawnNamed(defPublisherProps, "DefPublisherActor")
	if err != nil {
		panic(fmt.Errorf("DefPublisherPID SpawnNamed error:%s", err))
	}
	DefActorPublisher = NewActorPublisher(DefPublisherPID)
}

func NewActorPublisher(publisher *actor.PID, evtHub ...*eventhub.EventHub) *ActorPublisher {
	var hub *eventhub.EventHub
	if len(evtHub) == 0 {
		hub = DefEvtHub
	} else {
		hub = evtHub[0]
	}
	if publisher == nil {
		publisher = DefPublisherPID
	}
	return &ActorPublisher{
		EvtHub:    hub,
		Publisher: publisher,
	}
}

type ActorPublisher struct {
	EvtHub    *eventhub.EventHub
	Publisher *actor.PID
}

func (this *ActorPublisher) Publish(topic string, msg interface{}) {
	event := &eventhub.Event{
		Publisher: this.Publisher,
		Message:   msg,
		Topic:     topic,
		Policy:    eventhub.PublishPolicyAll,
	}
	this.EvtHub.Publish(event)
}

func (this *ActorPublisher) PublishEvent(evt *eventhub.Event) {
	this.EvtHub.Publish(evt)
}

type ActorSubscriber struct {
	EvtHub     *eventhub.EventHub
	Subscriber *actor.PID
}

func NewActorSubscriber(subscriber *actor.PID, evtHub ...*eventhub.EventHub) *ActorSubscriber {
	var hub *eventhub.EventHub
	if len(evtHub) == 0 {
		hub = DefEvtHub
	} else {
		hub = evtHub[0]
	}
	return &ActorSubscriber{
		EvtHub:     hub,
		Subscriber: subscriber,
	}
}

func (this *ActorSubscriber) Subscribe(topic string) {
	this.EvtHub.Subscribe(topic, this.Subscriber)
}

func (this *ActorSubscriber) Unsubscribe(topic string) {
	this.EvtHub.Unsubscribe(topic, this.Subscriber)
}
