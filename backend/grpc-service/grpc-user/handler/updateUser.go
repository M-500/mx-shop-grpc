package handler

import (
	"context"
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc-user/model"
	"grpc-user/proto"
	"grpc-user/svc"
	"strings"
	"time"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 17:11
//

func (u *UserService) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (resp *emptypb.Empty, err error) {
	//个人中心更新用户
	var user model.UserInfoModel
	db := svc.GetSrvCtx().MysqlConn
	result := db.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = db.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (u *UserService) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (resp *proto.CheckResponse, err error) {
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.CheckResponse{Success: check}, nil
}
