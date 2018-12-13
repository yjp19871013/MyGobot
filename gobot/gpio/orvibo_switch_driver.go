package gpio

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/orvibo"
)

const (
	stateEvent = "state"
)

type OrviboSwitchDriver struct {
	name     string
	uid      string
	deviceId string

	adapter *orvibo.Adapter
	state   bool
	gobot.Commander
	gobot.Eventer
}

func NewOrviboSwitchDriver(name string, uid string, deviceId string, adapter *orvibo.Adapter) *OrviboSwitchDriver {
	s := &OrviboSwitchDriver{
		name:      name,
		uid:       uid,
		deviceId:  deviceId,
		adapter:   adapter,
		state:     false,
		Commander: gobot.NewCommander(),
		Eventer:   gobot.NewEventer(),
	}

	s.AddCommand("Toggle", func(params map[string]interface{}) interface{} {
		return s.SwitchToggle()
	})

	s.AddCommand("On", func(params map[string]interface{}) interface{} {
		return s.SwitchOn()
	})

	s.AddCommand("Off", func(params map[string]interface{}) interface{} {
		return s.SwitchOff()
	})

	s.AddEvent(stateEvent)

	return s
}

func (s *OrviboSwitchDriver) Start() (err error) { return nil }

func (s *OrviboSwitchDriver) Halt() (err error) { return nil }

func (s *OrviboSwitchDriver) Name() string { return s.name }

func (s *OrviboSwitchDriver) SetName(n string) { s.name = n }

func (s *OrviboSwitchDriver) Connection() gobot.Connection {
	return s.adapter
}

// State return true if the led is On and false if the led is Off
func (s *OrviboSwitchDriver) State() bool {
	return s.state
}

// On sets the led to a high state.
func (s *OrviboSwitchDriver) SwitchOn() (err error) {
	result, err := s.adapter.Control(s.uid, s.deviceId, "on")
	if err != nil {
		s.Publish(stateEvent, err)
		return err
	}

	if result.Status == orvibo.StatusFailed {
		err := fmt.Errorf("switch control status error:", s.deviceId, "on")
		s.Publish(stateEvent, err)
		return err
	}

	s.state = true
	s.Publish(stateEvent, "on")

	return nil
}

// Off sets the led to a low state.
func (s *OrviboSwitchDriver) SwitchOff() (err error) {
	result, err := s.adapter.Control(s.uid, s.deviceId, "off")
	if err != nil {
		s.Publish(stateEvent, err)
		return err
	}

	if result.Status == orvibo.StatusFailed {
		err := fmt.Errorf("switch control status error:", s.deviceId, "off")
		s.Publish(stateEvent, err)
		return err
	}

	s.state = false
	s.Publish(stateEvent, "off")

	return nil
}

// Toggle sets the led to the opposite of it's current state
func (s *OrviboSwitchDriver) SwitchToggle() (err error) {
	if s.state {
		return s.SwitchOff()
	} else {
		return s.SwitchOn()
	}
}
