package engine

import "fmt"

type Event interface {
}

type Subscriber interface {
	HandleEvent(event Event)
}

type EventBus struct {
	subscribers map[Event][]Subscriber

	queue []Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: map[Event][]Subscriber{},
	}
}

// Flush should be called between game ticks to ensure systems don't conflict with each other.
func (e *EventBus) Flush() {
	for _, event := range e.queue {
		for _, s := range e.subscribers[eventName(event)] {
			s.HandleEvent(event)
		}
	}
	e.queue = nil
}

func (e *EventBus) Publish(event Event) {
	e.queue = append(e.queue, event)
}

func (e *EventBus) Subscribe(event Event, subscriber Subscriber) {
	name := eventName(event)
	e.subscribers[name] = append(e.subscribers[name], subscriber)
}

func eventName(event Event) string {
	return fmt.Sprintf("%T", event)
}
