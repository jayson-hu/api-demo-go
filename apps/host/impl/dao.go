package impl

import (
	"context"
	"fmt"
	"github.com/jayson-hu/api-demo-go/apps/host"
)

//完成对象和SQL数据库之间的转换的

// 把Host对象保存到数据库
func (i HostServiceImpl) save(ctx context.Context, ins *host.Host) error {


	var err error
	//初始化事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("start tx error, %s", err)
	}
	//处理事务的提交方式
	//1. 无错误则提交commit
	//2.有报错rollback回滚
	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				i.l.Error("rollback error, %s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error, %s", err)

			}
		}
	}()
	//插入resour语句
	rstmt, err := tx.Prepare(InsertResourceSQL)
	if err != nil {
		return err
	}
	defer rstmt.Close()
	_, err = rstmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}




	dstmt, err := tx.Prepare(InsertDescribeSQL)
	if err != nil {
		return err
	}
	defer dstmt.Close()
	_, err = dstmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}




	return nil

}
