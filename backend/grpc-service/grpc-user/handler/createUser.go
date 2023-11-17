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
	"grpc-user/model"
	"grpc-user/proto"
	"grpc-user/svc"
)

type UserService struct {
	proto.UnimplementedUserServer
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
