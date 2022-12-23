package command

import (
	"fmt"
	"go-command-pattern/bus"
	"go-command-pattern/service/alert"
)

type Executor interface {
	Execute() error
}

func (reading TemperatureReading) Execute() error {
	message := fmt.Sprintf("Temperature is %f", reading.Reading)
	go bus.PublishTo(reading.NotificationTopics, message)
	switch temperature := reading.Reading; {
	case temperature >= 40.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: message, Priority: alert.Critical})
		return nil
	case temperature >= 37.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: message, Priority: alert.High})
		return nil
	case temperature <= 4.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: message, Priority: alert.High})
		return nil
	case temperature <= -10.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: message, Priority: alert.Critical})
		return nil
	}
	return nil
}

func (reading HumidityReading) Execute() error {
	bus.PublishTo(reading.NotificationTopics, reading)
	return nil
}

func (reading WindReading) Execute() error {
	go bus.PublishTo(reading.NotificationTopics, reading)
	switch gust := reading.Gust; {
	case gust >= 120.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: reading, Priority: alert.Critical})
		return nil
	case gust >= 60.0:
		go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: reading, Priority: alert.High})
		return nil
	}
	return nil
}

func (reading LightningReading) Execute() error {
	go bus.PublishTo(reading.AlertTopics, alert.Alert{Reading: reading, Priority: alert.Critical})
	return nil
}
