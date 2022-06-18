package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Light struct {
	Entity  Entity
	Fetched bool
}

type Entity struct {
	State string `json:"state"`
	Attr  struct {
		Brightness float64 `json:"brightness"`
	} `json:"attributes"`
}

type Output struct {
	Text       string `json:"text,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Tooltip    string `json:"tooltip,omitempty"`
	Class      string `json:"class,omitempty"`
	Percentage int8   `json:"percentage"`
}

//
// Function Definitions
//

func (light *Light) fetchState() Entity {
	if light.Fetched {
		return light.Entity
	}

	body, statusCode := request.Get(config.API.Resources.State)

	var entity Entity

	if statusCode < 200 || statusCode > 299 {
		log.Fatal(fmt.Sprint(statusCode, " ", http.StatusText(statusCode)))
	}

	json.Unmarshal(body, &entity)

	light.Entity = entity
	light.Fetched = true

	return entity
}

/*
 * Outputs the current state of the light in json formatted for waybar
 */
func (light *Light) State() {
	if !light.Fetched {
		light.fetchState()
	}

	output := Output{
		Alt:        light.Entity.State,
		Class:      "light-" + light.Entity.State,
		Percentage: light.Brightness(),
	}

	marshalledOutput, err := json.Marshal(&output)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(marshalledOutput))
}

/*
 * Toggles the light on or off
 */
func (light *Light) Toggle() {
	request.Post(
		config.API.Resources.Service+"/toggle",
		[]byte(fmt.Sprintf(`{"entity_id":"%s"}`, config.API.EntityId)),
	)
}

/*
 * Get the current brightness as a percentage value
 */
func (light *Light) Brightness() int8 {
	if !light.Fetched {
		light.fetchState()
	}

	// The brightness is represented by a value from 0 to 255, dimmest to brightest,
	// so let's convert it to a percentage from 0 to 100 instead
	return int8((light.Entity.Attr.Brightness / 255.0) * 100.0)
}

func (light *Light) AlterBrightness(amount int64, absolute bool) {
	attr := "brightness_step_pct"

	if absolute {
		attr = "brightness_pct"
	}

	request.Post(
		config.API.Resources.Service+"/turn_on",
		[]byte(fmt.Sprintf(`{"entity_id":"%s","%s":%d}`, config.API.EntityId, attr, amount)),
	)
}
