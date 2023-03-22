package main

import "fmt"

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify(message string)
}

type Observer interface {
	Update(message string)
}

type MessagePublisher struct {
	observers []Observer
}

func (mp *MessagePublisher) Attach(observer Observer) {
	mp.observers = append(mp.observers, observer)
}

func (mp *MessagePublisher) Detach(observer Observer) {
	for i, obs := range mp.observers {
		if obs == observer {
			mp.observers = append(mp.observers[:i], mp.observers[i+1:]...)
			break
		}
	}
}

func (mp *MessagePublisher) Notify(message string) {
	for _, observer := range mp.observers {
		observer.Update(message)
	}
}

type MessageSubscriber struct {
	name string
}

func (ms *MessageSubscriber) Update(message string) {
	fmt.Printf("%s received message: %s\n", ms.name, message)
}

func main() {
	publisher := &MessagePublisher{}

	subscriber1 := &MessageSubscriber{name: "Subscriber 1"}
	subscriber2 := &MessageSubscriber{name: "Subscriber 2"}

	publisher.Attach(subscriber1)
	publisher.Attach(subscriber2)

	publisher.Notify("Hello, world!")

	publisher.Detach(subscriber2)

	publisher.Notify("How are you doing?")

}
