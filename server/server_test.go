package server

import (
	"github.com/boltdb/bolt"

	"github.com/lachlan351/alexa/events"
	"github.com/lachlan351/alexa/parser"
	"github.com/lachlan351/alexa/response"
	"github.com/lachlan351/alexa/validations"
)

func ExampleServer() {
	d, err := bolt.Open("info.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	validations.DB = d

	ev := events.New().
		Add("HelloWorld",
			func(ev *parser.Event) (*response.Response, error) {
				return response.New().
					AddSpeech("Hello, world!"), nil
			}).
		Add("HelloName",
			func(ev *parser.Event) (*response.Response, error) {
				name := ev.Request.Intent.Slots["Name"].Value

				return response.New().
					AddSpeech("Hello, " + name), nil
			})

	Run(ev)
}
