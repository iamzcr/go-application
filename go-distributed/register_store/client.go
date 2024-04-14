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
	fmt.Println(ServicesURL)
	res, err := http.Post(ServicesURL, "application/json", buf)
	fmt.Println(err)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(" Failed to registerservice.Registry serviceresponded with code %v", res.StatusCode)
	}
	return nil
}
