package impl

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

// 业务处理层
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("create host")
	i.l.Debugf("create host %s", ins.Name)
	//携带了metedata, 常用语trace
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create hsot with meta kv")

	//检验数据合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	// 把数据入库到 resource表和host表
	// 一次需要往2个表录入数据, 我们需要2个操作 要么都成功，要么都失败, 事务的逻辑
	ins.InjectDefault()
	//由Dao模块负责把对象入库
	err := i.save(ctx, ins)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.QueryHostRequest) (*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}
func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}
func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
