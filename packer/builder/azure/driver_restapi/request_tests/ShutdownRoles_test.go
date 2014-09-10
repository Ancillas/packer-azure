package request_tests

import (
	"fmt"
	"testing"
)

func _TestShutdownRoles(t *testing.T) {

	errMassage := "TestCaptureVmImage: %s\n"

	reqManager, err := getRequestManager()
	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	requestData := 	reqManager.ShutdownRoles(g_tmpServiceName, g_tmpVmName)
	fmt.Println(fmt.Sprintf("requestData:\n %v", requestData))
	_, err = reqManager.Execute(requestData)

	if err != nil {
		t.Errorf("reqManager error: %s\n", err.Error())
	}
}
