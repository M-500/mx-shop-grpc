package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"grpc-user/model"
	"grpc-user/proto"
	"grpc-user/svc"
)

// @Description
// @Author 代码小学生王木木
// @Date 2023/11/17 17:12
func ModelToResponse(user model.UserInfoModel) *proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return &userInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

/* 获取用户列表*/
func (u *UserService) GetUserList(ctx context.Context, req *proto.PageInfo) (resp *proto.UserListResponse, err error) {
	// 获取用户列表
	var users []model.UserInfoModel
	db := svc.GetSrvCtx().MysqlConn
	result := db.Find(&users)
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	db.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, userInfoRsp)
	}
	return rsp, nil
}
func (u *UserService) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (resp *proto.UserInfoResponse, err error) {
	var user model.UserInfoModel
	db := svc.GetSrvCtx()
	result := db.MysqlConn.Where(&model.UserInfoModel{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected != 1 {
		return nil, status.Errorf(codes.NotFound, "手机号不存在")
	}
	return ModelToResponse(user), nil
}
func (u *UserService) GetUserById(ctx context.Context, req *proto.IdRequest) (resp *proto.UserInfoResponse, err error) {
	var user model.UserInfoModel
	db := svc.GetSrvCtx()
	result := db.MysqlConn.Find(req.Id).First(&user)
	if result.RowsAffected != 1 {
		return nil, status.Errorf(codes.NotFound, "用户ID不存在")
	}
	return ModelToResponse(user), nil
}
