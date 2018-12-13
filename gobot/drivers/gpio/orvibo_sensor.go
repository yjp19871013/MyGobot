package gpio

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/orvibo"
)

const (
	valueEvent   = "value"
	statusFailed = 1
	interval     = 5 * time.Second
)

type OrviboSensorDriver struct {
	name       string
	uid        string
	deviceId   string
	deviceType string

	adapter *orvibo.Adapter
	unit    string
	halt    chan bool
	gobot.Commander
	gobot.Eventer
}

func NewOrviboSensorDriver(name string, uid string, deviceId string, deviceType string, unit string, adapter *orvibo.Adapter) *OrviboSensorDriver {
	s := &OrviboSensorDriver{
		name:       name,
		uid:        uid,
		deviceId:   deviceId,
		deviceType: deviceType,
		adapter:    adapter,
		unit:       unit,
		halt:       make(chan bool),
		Commander:  gobot.NewCommander(),
		Eventer:    gobot.NewEventer(),
	}

	s.AddCommand("state", func(params map[string]interface{}) interface{} {
		return s.state()
	})

	s.AddEvent(valueEvent)

	return s
}

func (s *OrviboSensorDriver) Start() (err error) {
	go func() {
		for {
			newValue, err := s.getValue()
			if err != nil {
				s.Publish(valueEvent, err)
			}

			s.Publish(valueEvent, fmt.Sprintf("%.2f", newValue)+s.unit)

			select {
			case <-s.halt:
				return
			case <-time.After(interval):

			}
		}
	}()

	return nil
}

func (s *OrviboSensorDriver) Halt() (err error) {
	s.halt <- true
	return nil
}

func (s *OrviboSensorDriver) Name() string { return s.name }

func (s *OrviboSensorDriver) SetName(n string) { s.name = n }

func (s *OrviboSensorDriver) Connection() gobot.Connection {
	return s.adapter
}

func (s *OrviboSensorDriver) state() string {
	value, err := s.getValue()
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%.2f", value) + s.unit
}

func (s *OrviboSensorDriver) getValue() (float32, error) {
	state, err := s.adapter.GetState(s.uid, s.deviceId)
	if err != nil {
		return 0.0, err
	}

	if state.Status == statusFailed {
		return 0.0, fmt.Errorf(state.Msg)
	}

	switch s.deviceType {
	case "temperature":
		return float32(state.Value1) / 100, nil
	case "humidity":
		return float32(state.Value2) / 100, nil
	default:
		panic("No this device type")
	}
}
