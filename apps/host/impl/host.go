package impl

import (
	"context"
	"github.com/infraboard/mcube/logger"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Debug("create host")
	i.l.Debugf("create host %s", ins.Name)
	//携带了metedata, 常用语trace
	i.l.With(logger.NewAny("request-id","req01")).Debug("create hsot with meta kv")


	return nil, nil
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
















