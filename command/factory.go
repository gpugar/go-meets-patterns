package command

import (
	"encoding/json"
	"fmt"
	"sync"
)

const (
	Temperature = "Temperature"
	Humidity    = "Humidity"
	Wind        = "Wind"
	Lightning   = "Lightning"
)

var lock = &sync.Mutex{}

type WeatherEventFactory struct {
	NotificationTopics []string
	AlertTopics        []string
}

var weatherEventFactory *WeatherEventFactory

func GetFactoryInstance() *WeatherEventFactory {
	if weatherEventFactory == nil {
		lock.Lock()
		defer lock.Unlock()
		if weatherEventFactory == nil {
			fmt.Println("Creating single Factory instance.")
			weatherEventFactory = &WeatherEventFactory{
				NotificationTopics: []string{"notify"},
				AlertTopics:        []string{"alert"},
			}
		}
	}
	return weatherEventFactory
}

func deser(rawEvent []byte, eventWrapper *Wrapper) error {
	err := json.Unmarshal(rawEvent, eventWrapper)
	if err != nil {
		return err
	}
	return nil
}

func populate(rawEvent []byte, event Executor) error {
	err := json.Unmarshal(rawEvent, event)
	if err != nil {
		return err
	}
	return nil
}

func (factory WeatherEventFactory) CreateReading(rawEvent []byte) (Executor, error) {
	var eventWrapper Wrapper
	if err := deser(rawEvent, &eventWrapper); err != nil {
		return nil, err
	}
	if spawnEvent, present := eventHatchery[eventWrapper.EventType]; present {
		fmt.Printf("Building event [%s]\n", eventWrapper.EventType)
		e := spawnEvent(factory.NotificationTopics, factory.AlertTopics)
		if err := populate(rawEvent, e); err != nil {
			return nil, err
		}
		return e, nil
	} else {
		return nil, fmt.Errorf("unknown event type [%s]", eventWrapper.EventType)
	}
}

var eventHatchery = map[Type]func(notify []string, alert []string) Executor{
	Temperature: func(notificationTopics []string, alertTopics []string) Executor {
		return &TemperatureReading{TopicHolder: TopicHolder{NotificationTopics: notificationTopics, AlertTopics: alertTopics}}
	},
	Humidity: func(topics []string, _ []string) Executor {
		return &HumidityReading{TopicHolder: TopicHolder{NotificationTopics: topics}}
	},
	Wind: func(notificationTopics []string, alertTopics []string) Executor {
		return &WindReading{TopicHolder: TopicHolder{NotificationTopics: notificationTopics, AlertTopics: alertTopics}}
	},
	Lightning: func(notificationTopics []string, alertTopics []string) Executor {
		return &LightningReading{TopicHolder: TopicHolder{NotificationTopics: notificationTopics, AlertTopics: alertTopics}}
	},
}
