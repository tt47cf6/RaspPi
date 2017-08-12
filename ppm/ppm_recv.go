package main

import (
	"log"
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/raspi"
)

func main() {
        r := raspi.NewAdaptor()
        pin := gpio.NewDirectPinDriver(r, "18")

        work := func() {
		vals := []int64{}
		valsP := &vals

		go func() {
			for {
				v := *valsP
				if len(v) > 0 {
					sum := 0.0
					for _, e := range v {
						sum += float64(e)
					}
					log.Printf("Len: %d  Avg: %f\n", len(v), sum/float64(len(v)))

					v = []int64{}
					valsP = &v
				}
				time.Sleep(250*time.Millisecond)
			}
		}()

		updateVals := func(d time.Duration) {
			v := *valsP
			i := int64(d / time.Microsecond)
			v = append(v, i)
			valsP = &v	
		}

		wasUp := false
		lstCh := time.Now()
		for {
			val, err := pin.DigitalRead()
			if err != nil {
				log.Printf("DigitalRead() error: %v\n", err)
			}
			if val > 0 && !wasUp {
				timeDn := time.Now().Sub(lstCh)
				//log.Printf("Was down for %v\n", timeDn)
				updateVals(timeDn)
				wasUp = true
				lstCh = time.Now()
			}
			if val == 0 && wasUp {
				//timeUp := time.Now().Sub(lstCh)
				//updateVals(timeUp)
				//log.Printf("Was up for %v\n", timeUp)
				wasUp = false
				lstCh = time.Now()
			}
		}
        }

        robot := gobot.NewRobot("ppmBot",
                []gobot.Connection{r},
                []gobot.Device{pin},
                work,
        )

        robot.Start()
}
