package handler

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:58
//

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-user/proto"
)

type UserService struct {
	proto.UnimplementedUserServer
}

func (u *UserService) GetUserList(context.Context, *proto.PageInfo) (resp *proto.UserListResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (u *UserService) GetUserByMobile(context.Context, *proto.MobileRequest) (resp *proto.UserInfoResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (u *UserService) GetUserById(context.Context, *proto.IdRequest) (resp *proto.UserInfoResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (u *UserService) CreateUser(context.Context, *proto.CreateUserInfo) (resp *proto.UserInfoResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (u *UserService) UpdateUser(context.Context, *proto.UpdateUserInfo) (resp *emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (u *UserService) CheckPassWord(context.Context, *proto.PasswordCheckInfo) (resp *proto.CheckResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}
