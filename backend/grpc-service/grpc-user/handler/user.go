package handler

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 12:58
//

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-user/model"
	"grpc-user/proto"
	"grpc-user/svc"
)

type UserService struct {
	proto.UnimplementedUserServer
}

/* 获取用户列表*/
func (u *UserService) GetUserList(ctx context.Context, req *proto.PageInfo) (resp *proto.UserListResponse, err error) {
	// 获取用户列表
	var users []model.UserInfoModel
	srvCtx := svc.GetSrvCtx()
	srvCtx.MysqlConn.Find(&users)
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (u *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (resp *proto.UserInfoResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByMobile not implemented")
}
func (u *UserService) GetUserById(ctx context.Context, req *proto.IdRequest) (resp *proto.UserInfoResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (u *UserService) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (resp *proto.UserInfoResponse, err error) {
	// 新增用户
	var user model.UserInfoModel
	db := svc.GetSrvCtx().MysqlConn
	result := db.Where(&model.UserInfoModel{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected >= 1 {
		return nil, status.Errorf(codes.AlreadyExists, "手机号已经存在，无法重复创建")
	}
	user.Mobile = req.Mobile
	user.NickName = req.NickName
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = db.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	resp = &proto.UserInfoResponse{}
	resp.Id = user.ID
	resp.NickName = user.NickName
	resp.Gender = user.Gender
	return resp, nil
}
func (u *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (resp *emptypb.Empty, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (u *UserService) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (resp *proto.CheckResponse, err error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPassWord not implemented")
}
