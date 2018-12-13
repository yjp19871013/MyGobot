package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var (
	rgb = [][]byte{
		{255, 0, 0},
		{0, 255, 0},
		{0, 0, 255},
	}

	current = 0
)

func main() {
	adapter := firmata.NewAdaptor("/dev/ttyACM0")
	rgbLed := gpio.NewRgbLedDriver(adapter, "11", "10", "9")
	button := gpio.NewButtonDriver(adapter, "8")

	work := func() {
		timerRGBLed(rgbLed)
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{adapter},
		[]gobot.Device{rgbLed, button},
		work,
	)

	master := gobot.NewMaster()
	master.AddRobot(robot)
	api.NewAPI(master).Start()

	robot.Start()
}

func timerRGBLed(rgbLed *gpio.RgbLedDriver) {
	gobot.Every(500*time.Millisecond, func() {
		err := rgbLed.SetRGB(rgb[current][0], rgb[current][1], rgb[current][2])
		if err != nil {
			fmt.Println(err)
			return
		}

		current = (current + 1) % 3
	})
}

func controlRGBLedWithButton(rgbLed *gpio.RgbLedDriver, button *gpio.ButtonDriver) {
	err := rgbLed.SetRGB(rgb[current][0], rgb[current][1], rgb[current][2])
	if err != nil {
		fmt.Println(err)
		return
	}

	current = current + 1

	state := button.Active
	gobot.Every(50*time.Millisecond, func() {
		if button.Active != state {
			if !button.Active {
				err := rgbLed.SetRGB(rgb[current][0], rgb[current][1], rgb[current][2])
				if err != nil {
					fmt.Println(err)
					return
				}

				current = (current + 1) % 3
			}

			state = button.Active
		}
	})
}
