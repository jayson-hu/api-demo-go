package impl

import (
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/sqlbuilder"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

// 业务处理层
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	i.l.Info("create host")
	i.l.Debugf("create host %s", ins.Name)
	//携带了metedata, 常用语trace
	i.l.With(logger.NewAny("request-id", "")).Debug("create hsot with meta kv")

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

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescirbeHostRequest) (*host.Host, error) {
	b := sqlbuilder.NewBuilder(QueryHostSQL)
	b.Where("r.id = ?", req.Id)
	querySQL, args := b.Build()
	i.l.Debugf("describe sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	ins := host.NewHost()
	err = stmt.QueryRowContext(ctx, args...).Scan(&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP,
		&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber)
	if err != nil {
		return nil, err
	}
	return ins, nil

}
func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	b := sqlbuilder.NewBuilder(QueryHostSQL)

	if req.Keywords != "" {
		// (r.`name`='%' OR r.description='%' OR r.private_ip='%' OR r.public_ip='%')
		//  10.10.1, 接口测试
		b.Where("r.`name`LIKE ? OR r.description LIKE ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	b.Limit(req.OffSet(), req.GetPageSize())
	querySQL, args := b.Build()
	i.l.Debugf("query sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next() {
		// 没扫描一行,就需要读取出来
		// h.cpu, h.memory, h.gpu_spec, h.gpu_amount, h.os_type, h.os_name, h.serial_number
		ins := host.NewHost()
		if err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
			&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
			&ins.Account, &ins.PublicIP, &ins.PrivateIP,
			&ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)

	}

	// total统计
	countSQL, args := b.BuildCount()
	i.l.Debugf("count sql: %s, args: %v", countSQL, args)
	countStmt, err := i.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
}
func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	// 获取已有的对象
	ins, err := i.DescribeHost(ctx, host.NewDescribeHostWithId(req.Id))
	if err != nil {
		return nil, err
	}
	//更新的更新模式
	switch req.UpdateMode {
	case host.UPDATE_MODE_PUT:
		err := ins.Put(req.Host)
		if err != nil {
			return nil, err
		}
	case host.UPDATE_MODE_PATCH:
		err := ins.Patch(req.Host)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("update mode only put or patch")

	}
	//检查更新的数据是否合法
	err = ins.Validate()
	if err != nil{
		return nil, err
	}
	//更新数据库里面的数据
	if err := i.update(ctx,ins); err != nil{
		return nil, err
	}
	//返回更新的后的对象

	return ins, nil
}
func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
