package main

import (
	"github.com/nadavbm/controller/controller"
	"github.com/nadavbm/controller/pkg/logger"
	"k8s.io/apimachinery/pkg/util/runtime"
)

func main() {
	logger := logger.New()

	controller, err := controller.New(logger)
	if err != nil {
		panic(err)
	}
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()

	controller.Run(stopper)
	<-stopper
}
