package logic

import (
	"bookstore/rpc/model"
	"context"

	"bookstore/rpc/add/internal/svc"
	"bookstore/rpc/add/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.AddReq) (*pb.AddResp, error) {
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Book{
		Book:  in.Book,
		Price: in.Price,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddResp{
		Ok: true,
	}, nil

}
