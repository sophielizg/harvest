package helloworld

import (
	"github.com/sophielizg/harvest/api/pkg/app"
)

type Success struct {
	Hello string `json:"name"`
}

type Failure struct {
	Good string `json:"good"`
}

func SayHello(currentApp app.App, name string) interface{} {
	if name != "" {
		return Success{Hello: name}
	}

	return Failure{Good: "bye"}
}
