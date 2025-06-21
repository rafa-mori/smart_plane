package main

import (
	l "github.com/rafa-mori/logz"
	gl "github.com/rafa-mori/smart_plane/logger"
)

var logger l.Logger

// main initializes the logger and creates a new SmartPlane instance.
func main() {
	if err := RegX().Command().Execute(); err != nil {
		gl.Log("fatal", err.Error())
	}
}
