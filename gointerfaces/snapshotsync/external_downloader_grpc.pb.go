// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package snapshotsync

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DownloaderClient is the client API for Downloader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloaderClient interface {
	Download(ctx context.Context, in *DownloadSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsInfoReply, error)
}

type downloaderClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloaderClient(cc grpc.ClientConnInterface) DownloaderClient {
	return &downloaderClient{cc}
}

func (c *downloaderClient) Download(ctx context.Context, in *DownloadSnapshotRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snapshotsync.Downloader/Download", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downloaderClient) Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsInfoReply, error) {
	out := new(SnapshotsInfoReply)
	err := c.cc.Invoke(ctx, "/snapshotsync.Downloader/Snapshots", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloaderServer is the server API for Downloader service.
// All implementations must embed UnimplementedDownloaderServer
// for forward compatibility
type DownloaderServer interface {
	Download(context.Context, *DownloadSnapshotRequest) (*emptypb.Empty, error)
	Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsInfoReply, error)
	mustEmbedUnimplementedDownloaderServer()
}

// UnimplementedDownloaderServer must be embedded to have forward compatible implementations.
type UnimplementedDownloaderServer struct {
}

func (UnimplementedDownloaderServer) Download(context.Context, *DownloadSnapshotRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedDownloaderServer) Snapshots(context.Context, *SnapshotsRequest) (*SnapshotsInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshots not implemented")
}
func (UnimplementedDownloaderServer) mustEmbedUnimplementedDownloaderServer() {}

// UnsafeDownloaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloaderServer will
// result in compilation errors.
type UnsafeDownloaderServer interface {
	mustEmbedUnimplementedDownloaderServer()
}

func RegisterDownloaderServer(s grpc.ServiceRegistrar, srv DownloaderServer) {
	s.RegisterService(&_Downloader_serviceDesc, srv)
}

func _Downloader_Download_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloaderServer).Download(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshotsync.Downloader/Download",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloaderServer).Download(ctx, req.(*DownloadSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Downloader_Snapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloaderServer).Snapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snapshotsync.Downloader/Snapshots",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloaderServer).Snapshots(ctx, req.(*SnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Downloader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "snapshotsync.Downloader",
	HandlerType: (*DownloaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Download",
			Handler:    _Downloader_Download_Handler,
		},
		{
			MethodName: "Snapshots",
			Handler:    _Downloader_Snapshots_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snapshot_downloader/external_downloader.proto",
}
