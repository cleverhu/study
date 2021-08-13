package services

import "jtthinkStudy/primary/eventBusTest/eventbus"

var Bus *eventbus.EventBus

func init() {
	Bus = eventbus.NewEventBus()
}
