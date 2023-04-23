// Code generated by goctl. DO NOT EDIT.
// Source: topicsrv.proto

package topicsrv

import (
	"context"

	"github.com/lixvyang/betxin-micro/service/topic/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddTopicReq      = pb.AddTopicReq
	AddTopicResp     = pb.AddTopicResp
	DelTopicReq      = pb.DelTopicReq
	DelTopicResp     = pb.DelTopicResp
	GetTopicByIdReq  = pb.GetTopicByIdReq
	GetTopicByIdResp = pb.GetTopicByIdResp
	SearchTopicReq   = pb.SearchTopicReq
	SearchTopicResp  = pb.SearchTopicResp
	Topic            = pb.Topic
	UpdateTopicReq   = pb.UpdateTopicReq
	UpdateTopicResp  = pb.UpdateTopicResp

	Topicsrv interface {
		// -----------------------topic-----------------------
		AddTopic(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*AddTopicResp, error)
		UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error)
		DelTopic(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*DelTopicResp, error)
		GetTopicById(ctx context.Context, in *GetTopicByIdReq, opts ...grpc.CallOption) (*GetTopicByIdResp, error)
		SearchTopic(ctx context.Context, in *SearchTopicReq, opts ...grpc.CallOption) (*SearchTopicResp, error)
	}

	defaultTopicsrv struct {
		cli zrpc.Client
	}
)

func NewTopicsrv(cli zrpc.Client) Topicsrv {
	return &defaultTopicsrv{
		cli: cli,
	}
}

// -----------------------topic-----------------------
func (m *defaultTopicsrv) AddTopic(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*AddTopicResp, error) {
	client := pb.NewTopicsrvClient(m.cli.Conn())
	return client.AddTopic(ctx, in, opts...)
}

func (m *defaultTopicsrv) UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error) {
	client := pb.NewTopicsrvClient(m.cli.Conn())
	return client.UpdateTopic(ctx, in, opts...)
}

func (m *defaultTopicsrv) DelTopic(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*DelTopicResp, error) {
	client := pb.NewTopicsrvClient(m.cli.Conn())
	return client.DelTopic(ctx, in, opts...)
}

func (m *defaultTopicsrv) GetTopicById(ctx context.Context, in *GetTopicByIdReq, opts ...grpc.CallOption) (*GetTopicByIdResp, error) {
	client := pb.NewTopicsrvClient(m.cli.Conn())
	return client.GetTopicById(ctx, in, opts...)
}

func (m *defaultTopicsrv) SearchTopic(ctx context.Context, in *SearchTopicReq, opts ...grpc.CallOption) (*SearchTopicResp, error) {
	client := pb.NewTopicsrvClient(m.cli.Conn())
	return client.SearchTopic(ctx, in, opts...)
}
