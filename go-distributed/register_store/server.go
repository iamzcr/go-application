package register_store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// 注册服务的服务端
const ServerPort = ":3000"
const ServicesURL = "http://127.0.0.1" + ServerPort + "/services"

// 存放所有被注册的服务
type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

// 注册一个服务
func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

// 声明包级别的全局变量
var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

// RegistryService 接受注册服务的服务请求的服务端逻辑
type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	//通过发送post请求过来注册一个服务，例如日志服务
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL:%s\\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
	case http.MethodDelete:
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("Removing service at URL: %s", url)
		err = reg.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// 取消服务
func (r *registry) remove(url string) error {
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			reg.registrations = append(reg.registrations[:i], reg.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("Service at URL %s not found", url)
}
