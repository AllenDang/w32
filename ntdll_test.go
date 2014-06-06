package w32

import (
	"testing"
)

var testPortName = "\\TestAlpcPort"

func TestNtAlpcCreatePort(t *testing.T) {

	hPort, err := UnsecuredAlpcPort(testPortName)

	if err != nil {
		t.Errorf("failed to create ALPC port %v: %v", testPortName, err)
	} else {
		t.Logf("[OK] Created ALPC port %v with handle 0x%x", testPortName, hPort)
	}
}
