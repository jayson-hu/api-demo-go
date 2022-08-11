package host

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type Vendor int
type UPDATE_MODE string

const (
	UPDATE_MODE_PUT UPDATE_MODE = "put"

	UPDATE_MODE_PATCH UPDATE_MODE = "patch"
)
const (
	Private_IDC Vendor = iota
	ALIYUN
	TXYUN
)

type Service interface {
	// CreateHost 录入主机
	CreateHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// DescribeHost 查询主机详情
	DescribeHost(context.Context, *DescirbeHostRequest) (*Host, error)
	// UpdateHost 主机更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// DeleteHost 主机删除, 比如前端需要打印当前删除主机的ip或者其他的信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}
type HostSet struct {
	Total int
	Items []*Host
}

func NewPutUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		UpdateMode: UPDATE_MODE_PUT,
		Host: h,
		//Describe:&Describe{},
		//Resource: &Resource{},

	}
}
func NewPatchUpdateHostRequest(id string) *UpdateHostRequest {
	h := NewHost()
	h.Id = id
	return &UpdateHostRequest{
		//Id: id,
		UpdateMode: UPDATE_MODE_PATCH,
		//Describe: &Describe{},
		//Resource: &Resource{},
		Host: h,
	}
}

type UpdateHostRequest struct {
	//Id         string      `json:"id"`
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Host
	//*Describe
	//*Resource
}
type DeleteHostRequest struct {
	Id string
}
type QueryHostRequest struct {
	PageSize   int    `json:"page_size"`
	PageNumber int    `json:"page_number"`
	Keywords   string `json:"kws"`
}

func NewHost() *Host {
	return &Host{
		Resource: &Resource{},
		Describe: &Describe{},
	}
}

func NewHostSet() *HostSet {
	return &HostSet{
		Items: []*Host{},
	}
}

type DescirbeHostRequest struct {
	Id string
}

func (s *HostSet) Add(item *Host) {
	s.Items = append(s.Items, item)
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

func (h *Host) InjectDefault() {
	if h.CreateAt == 0 {
		h.CreateAt = time.Now().UnixNano() / 1e6 //毫秒
	}

}

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

func NewQueryHostFromHTTP(r *http.Request) *QueryHostRequest {
	req := NewQueryHostRequest()
	// query string
	qs := r.URL.Query()
	pss := qs.Get("page_size")
	if pss != "" {
		req.PageSize, _ = strconv.Atoi(pss)
	}

	pns := qs.Get("page_number")
	if pns != "" {
		req.PageNumber, _ = strconv.Atoi(pns)
	}

	req.Keywords = qs.Get("kws")
	return req
}

func NewQueryHostRequest() *QueryHostRequest {
	return &QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

func NewDescribeHostWithId(id string) *DescirbeHostRequest {
	return &DescirbeHostRequest{
		Id: id,
	}
}

func (req *QueryHostRequest) GetPageSize() uint {
	return uint(req.PageSize)
}

func (req *QueryHostRequest) OffSet() int64 {
	return int64((req.PageNumber - 1) * req.PageSize)
}

// 对象的全量更新
func (h *Host) Put(obj *Host) error {
	//h.Description = obj // 这种是直接对象更换，把对象直接更换
	//*h.Description = *obj // 不是指针的copy， 不是值的更换
	if obj.Id != h.Id {
		return fmt.Errorf("id not equal")
	}
	*h.Resource = *obj.Resource
	*h.Describe = *obj.Describe
	return nil

}

func (h *Host) Patch(obj *Host) error {
	if obj.Name != "" {
		h.Name = obj.Name
	}
	if obj.CPU != 0 {
		h.CPU = obj.CPU
	}

	return nil

}
