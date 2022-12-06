package test

import (
	"testing"

	"github.com/ShadowFl0w/logger"
)

func TestLogger(t *testing.T) {
	logger.Info("Info here!")
	logger.Println("Println here!")

	logger.SetLevel(logger.LPrintln)

	logger.Info("Info Level here!")
	logger.Println("Println Level here!")
}
