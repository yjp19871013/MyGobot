package orvibo

import (
	"encoding/json"
	"strconv"
	"time"
)

const (
	controlUrl  = "/api/control"
	stateGetUrl = "/api/deviceStatus"
)

type Config struct {
	AppId    string
	AppKey   string
	Host     string
	Username string
	Password string
}

type Adapter struct {
	name   string
	config *Config
}

func NewAdapter(name string, config *Config) *Adapter {
	return &Adapter{
		name:   name,
		config: config,
	}
}

func (adapter *Adapter) Name() string {
	return adapter.name
}

func (adapter *Adapter) SetName(name string) {
	adapter.name = name
}

func (adapter *Adapter) Connect() error {
	return nil
}

func (adapter *Adapter) Finalize() error {
	return nil
}

func (adapter *Adapter) Control(uid string, deviceId string, order string) (*OrviboControlResult, error) {
	params := make(map[string]string)
	params["appId"] = adapter.config.AppId
	params["time"] = strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	params["sn"] = GetUUID()
	params["userName"] = adapter.config.Username
	params["password"] = adapter.config.Password
	params["uid"] = uid
	params["deviceId"] = deviceId
	params["order"] = order
	params["value1"] = "0"
	params["value2"] = "0"
	params["value3"] = "0"
	params["value4"] = "0"
	params["delayTime"] = "0"

	sign := CalSign(controlUrl, "post", params, adapter.config.AppKey)
	params["sig"] = sign
	request := ParamsMapToString(params)
	result, err := PostJson(adapter.config.Host+controlUrl, request)
	if err != nil {
		return nil, err
	}

	orviboResult := &OrviboControlResult{}
	json.Unmarshal([]byte(result), orviboResult)
	return orviboResult, nil
}

func (adapter *Adapter) GetState(uid string, deviceId string) (*OrviboDeviceState, error) {
	params := make(map[string]string)
	params["appId"] = adapter.config.AppId
	params["time"] = strconv.Itoa(int(time.Now().UnixNano() / 1e6))
	params["sn"] = GetUUID()
	params["userName"] = adapter.config.Username
	params["password"] = adapter.config.Password
	params["uid"] = uid
	params["deviceId"] = deviceId

	sign := CalSign(stateGetUrl, "post", params, adapter.config.AppKey)
	params["sig"] = sign
	request := ParamsMapToString(params)
	result, err := PostJson(adapter.config.Host+stateGetUrl, request)
	if err != nil {
		return nil, err
	}

	orviboState := &OrviboDeviceState{}
	json.Unmarshal([]byte(result), orviboState)
	return orviboState, nil
}
