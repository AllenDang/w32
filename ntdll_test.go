package w32

import "testing"

var testPortName = "\\TestAlpcPort"

func TestNtAlpcCreatePort(t *testing.T) {

	objAttr := NewObjectAttributes(testPortName, 0, 0, nil)
	portAttr := ALPC_PORT_ATTRIBUTES{MaxMessageLength: 0x148}
	hPort, err := NtAlpcCreatePort(&objAttr, &portAttr)
	if err != nil {
		t.Errorf("failed to create ALPC port %v: %v", testPortName, err)
	}

	t.Logf("[OK] Created ALPC port %v with handle 0x%x", testPortName, hPort)
}
