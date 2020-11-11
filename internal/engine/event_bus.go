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

	debug bool
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: map[Event][]Subscriber{},
		debug:       true,
	}
}

// Flush should be called between game ticks to ensure systems don't conflict with each other.
func (e *EventBus) Flush() {
	// The outer loop is needed, because events can trigger more events.
	for len(e.queue) > 0 {
		queue := e.queue
		e.queue = nil
		for _, event := range queue {
			for _, s := range e.subscribers[eventName(event)] {
				if e.debug {
					fmt.Printf("%T -> %T\n", event, s)
				}

				s.HandleEvent(event)
			}
		}
	}
}

func (e *EventBus) Publish(event Event) {
	if e.debug {
		fmt.Printf("Publishing %T\n", event)
	}
	e.queue = append(e.queue, event)
}

func (e *EventBus) Subscribe(event Event, subscriber Subscriber) {
	if e.debug {
		fmt.Printf("Subscribing %T -> %T\n", event, subscriber)
	}
	name := eventName(event)
	e.subscribers[name] = append(e.subscribers[name], subscriber)
}

func eventName(event Event) string {
	return fmt.Sprintf("%T", event)
}
