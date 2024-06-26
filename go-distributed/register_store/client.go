package register_store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// 请求注册服务业务的客户端
func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	fmt.Println("post regiser service url:", ServicesURL)
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(" Failed to registerservice.Registry serviceresponded with code %v", res.StatusCode)
	}
	return nil
}
func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL,
		bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/plain")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service. Registry"+
			" service responded with code %v", res.StatusCode)
	}
	return nil
}
