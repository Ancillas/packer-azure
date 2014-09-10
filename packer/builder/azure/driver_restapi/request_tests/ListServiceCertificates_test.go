package request_tests

import (
	"testing"
	"github.com/MSOpenTech/packer-azure/packer/builder/azure/driver_restapi/response"
	"github.com/MSOpenTech/packer-azure/packer/builder/azure/driver_restapi/cert"
)

func _TestListServiceCertificates(t *testing.T) {

	errMassage := "TestListServiceCertificates: %s\n"

	reqManager, err := getRequestManager()
	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	serviceName := "PkrSrvmosl521tfw"
	requestData := reqManager.ListServiceCertificates(serviceName)
	resp, err := reqManager.Execute(requestData)

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	list, err := response.ParseServiceCertificateList(resp.Body)

	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	t.Logf("ServiceCertificateList:\n\n")

	for _, val := range(list.Certificates){
		t.Logf("%v\n\n", val)
	}

	base64X509 := list.Certificates[0].Data
	certPath := "d:\\temp\\shch.cer"
	err = cert.ToX509File(base64X509, certPath)
	if err != nil {
		t.Errorf(errMassage, err.Error())
	}

	t.Error("eom")
}