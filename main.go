package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

var (
	config  Config
	request Request
)

func main() {
	// We don't need timestamps for logs, this isn't a long running program.
	log.SetFlags(0)

	var (
		configFilePath string
		entityId       string
		token          string
		host           string
	)

	flag.StringVar(&configFilePath, "c", "", "path to an alternate configuration file")
	flag.StringVar(&configFilePath, "config", "", "path to an alternate configuration file")
	flag.StringVar(&entityId, "entity", "", "specify an light entity to control")
	flag.StringVar(&token, "token", "", "specify the auth token to use for api requests")
	flag.StringVar(&host, "host", "", "specify the host to use for api requests")

	flag.Parse()

	config = Config{
		API: APIConfig{
			AuthToken: token,
			Host:      host,
			EntityId:  entityId,
		},
	}

	config.Load(configFilePath)

	config.API.Resources.State = config.API.Host + "/api/states/" + config.API.EntityId
	config.API.Resources.Service = config.API.Host + "/api/services/light"

	request = Request{}
	light := Light{}

	switch flag.Arg(0) {
	case "toggle":
		light.Toggle()
	case "brightness":
		if flag.NArg() > 1 {
			absolute := false
			brightnessValue := flag.Arg(1)

			amount, err := strconv.ParseInt(brightnessValue, 10, 0)

			if err != nil {
				log.Fatal("Invalid brightness value, must be an integer of the format +0, 0, or -0")
			}

			// If the first character is not a + or -, set an absolute value
			if brightnessValue[:1] != "+" && brightnessValue[:1] != "-" {
				absolute = true
			}

			if amount > 100 || amount < -100 {
				lowerBound := -100

				if absolute {
					lowerBound = 0
				}

				log.Fatalf("Brightness value out-of-bounds, must be an integer between %d to 100", lowerBound)
			}

			light.AlterBrightness(amount, absolute)
		} else {
			fmt.Println(light.Brightness())
		}
	default:
		light.State()
	}
}
