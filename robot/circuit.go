package main

import (
        "fmt"

	"gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/i2c"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
	af := i2c.NewAdafruitMotorHatDriver(r)
        work := func() {
		if err := af.Start(); err != nil {
			fmt.Println("start error: %v\n", err)
		}
		if err := af.SetDCMotorSpeed(0, 255); err != nil {
			fmt.Printf("set error: %v\n", err)
		}
		if err := af.RunDCMotor(0, i2c.AdafruitForward); err != nil {
			fmt.Printf("Run error: %v\n", err)
		}
	}

        robot := gobot.NewRobot("motorCircuitRunner",
                []gobot.Connection{r},
                []gobot.Device{af},
                work,
        )

        robot.Start()
}
