package host

import (
	"context"
	"github.com/go-playground/validator/v10"
	"time"
)

var (
	validate = validator.New()
)

type Service interface {
	// CreateHost 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// DescribeHost 查询主机详情
	DescribeHost(context.Context, *QueryHostRequest) (*Host, error)
	// UpdateHost 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// DeleteHost 主机删除, 比如前端需要打印当前删除主机的ip或者其他的信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}
type HostSet struct {
	Items []*Host
	Total int
}
type UpdateHostRequest struct {
	*Describe
}
type DeleteHostRequest struct {
	Id string
}
type QueryHostRequest struct {
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

//Host 模型的定义
type Host struct {
	//资源共有属性
	*Resource
	//私有属性
	*Describe
}

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (h *Host) InjectDefault()  {
	if h.CreateAt ==0{
		h.CreateAt = time.Now().UnixNano() / 1e6 //毫秒
	}

}


const (
	Private_IDC Vendor = iota
	ALIYUN
	TXYUN
)

type Vendor int

type Resource struct {
	Id       string `json:"id" validate:"required"` //全局唯一ID
	Vendor   Vendor `json:"vendor"`
	Region   string `json:"region" validate:"required"`
	Zone     string `json:"zone"`
	CreateAt int64  `json:"create_at"`
	ExpireAt int64  `json:"expire_at"`
	Category string `json:"category"`
	Type     string `json:"type" validate:"required"`
	//InstanceID  string `json:"instance_id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status"` //服务商中的状态
	Tags        string `json:"tags"`
	UpdateAt    int64  `json:"update_at"`
	SyncAt      int64  `json:"sync_at"`
	Account     string `json:"account"` //资源所属账号
	PublicIP    string `json:"public_ip"`
	PrivateIP   string `json:"private_ip"`
	PayType     string `json:"pay_type"`
}

type Describe struct {
	CPU          int    `json:"cpu" validate:"required"`
	Memory       int    `json:"memory" validate:"required"`
	GPUAmount    int    `json:"gpu_amount"`
	GPUSpec      string `json:"gpu_spec"`
	OSType       string `json:"os_type"`
	OSName       string `json:"os_name"`
	SerialNumber string `json:"serial_number"`
}
