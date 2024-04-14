package register_store

// 需要被注册的服务信息
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}
type ServiceName string

const (
	LogService = ServiceName(" LogService")
)
