package bus

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Event struct {
	Data  interface{}
	Topic string
}

type Channel chan Event

type ChannelSlice []Channel

type EventBus struct {
	topicChannels map[string]ChannelSlice
	mutex         sync.RWMutex
}

var eventBusSingleton *EventBus

func GetInstance() *EventBus {
	if eventBusSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if eventBusSingleton == nil {
			fmt.Println("Creating EventBus singleton.")
			eventBusSingleton = &EventBus{
				topicChannels: map[string]ChannelSlice{},
			}
		}
	}
	return eventBusSingleton
}

func (bus *EventBus) Publish(topic string, data interface{}) {
	bus.mutex.RLock()
	defer bus.mutex.RUnlock()
	if slices, found := bus.topicChannels[topic]; found {
		channels := append(ChannelSlice{}, slices...)
		go func(event Event, dataChannelSlices ChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- event
			}
		}(Event{Data: data, Topic: topic}, channels)
	}

}

func (bus *EventBus) Subscribe(topic string, chanel Channel) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()
	if residual, found := bus.topicChannels[topic]; found {
		bus.topicChannels[topic] = append(residual, chanel)
	} else {
		bus.topicChannels[topic] = append([]Channel{}, chanel)
	}
}

func PublishTo(topics []string, data interface{}) {
	busInstance := GetInstance()
	for _, topic := range topics {
		busInstance.Publish(topic, data)
	}
}
