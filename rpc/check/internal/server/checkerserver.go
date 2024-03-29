// Code generated by goctl. DO NOT EDIT!
// Source: check.proto

package server

import (
	"context"

	"bookstore/rpc/check/internal/logic"
	"bookstore/rpc/check/internal/svc"
	"bookstore/rpc/check/pb/pb"
)

type CheckerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCheckerServer
}

func NewCheckerServer(svcCtx *svc.ServiceContext) *CheckerServer {
	return &CheckerServer{
		svcCtx: svcCtx,
	}
}

func (s *CheckerServer) Check(ctx context.Context, in *pb.CheckReq) (*pb.CheckResp, error) {
	l := logic.NewCheckLogic(ctx, s.svcCtx)
	return l.Check(in)
}
