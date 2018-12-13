package orvibo

const (
	StatusSuccess = iota
	StatusFailed
)

type OrviboControlResult struct {

	//状态码
	Status int `json:"status"`

	//说明信息
	Msg string `json:"msg"`
}

type OrviboDeviceState struct {
	Status   int    `json:"status"`
	Msg      string `json:"msg"`
	Uid      string `json:"uid"`
	DeviceId string `json:"deviceId"`

	Value1 int `json:"value1"`
	Value2 int `json:"value2"`
	Value3 int `json:"value3"`
	Value4 int `json:"value4"`

	// 在线状态 离线:0 在线:1
	Online int `json:"online"`
}

type OrviboGetEnergyDayResult struct {
	Status int                  `json:"status"`
	Msg    string               `json:"msg"`
	Data   []OrviboEnergyDetail `json:"data"`
}

type OrviboEnergyDetail struct {
	DeviceId    string  `json:"deviceId"`
	Energy      float64 `json:"energy"`
	WorkingTime int     `json:"workingTime"`
	Day         string  `json:"day"`
}

type OrviboGetDistributionBoxLastValueResult struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`

	// 电流,单位毫安
	Current int `json:"CURRENT"`

	// 电压,单位100毫伏
	Voltage int `json:"VOLTAGE"`

	// 功率, 单位0.1瓦
	Power int `json:"POWER"`

	// 功率因素PowerFactor = 功率因素真实值*100
	PowerFactor int `json:"POWER_FACTOR"`

	// 电量,单位 千瓦/时
	Energy float64 `json:"ENERGY"`
}
