package logic

import (
	"context"
	"gf-cx/internal/dao"
	"gf-cx/internal/model"
	"gf-cx/internal/model/do"
	"gf-cx/internal/model/entity"
	"gf-cx/internal/service"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(NewUser())
}

func NewUser() service.IUser {
	return &sUser{}
}

func (s *sUser) SignUp(ctx context.Context, in model.UserSignUpInput) (userId uint, err error) {
	if in.RePassword != in.Password {
		return 0, gerror.Newf("重复密码错误", in.RePassword)
	}
	aUserName, err := s.IsUserNameAvailable(ctx, in.UserName)
	if err != nil {
		return 0, err
	}
	if !aUserName {
		return 0, gerror.Newf("用户名已被注册", in.UserName)
	}

	aMobile, err := s.IsMobileAvailable(ctx, in.Mobile)
	if err != nil {
		return 0, err
	}
	if !aMobile {
		return 0, gerror.Newf("手机号已被注册", in.Mobile)
	}

	in.Password = s.EncryptPassword(in.UserName, in.Password)
	uId, err := dao.Users.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return 0, err
	}
	return uint(uId), err
}

func (s *sUser) SignIn(ctx context.Context, in model.UserSignInInput) (err error) {
	var user *entity.Users
	err = dao.Users.Ctx(ctx).
		Where(do.Users{
			UserName: in.UserName,
			Password: s.EncryptPassword(in.UserName, in.Password),
		}).
		Scan(&user)
	if err != nil {
		return err
	}
	if user == nil {
		return gerror.New("用户名或密码错误！")
	}
	err = service.Session().SetUser(ctx, user)
	if err != nil {
		return err
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		UserName: user.UserName,
		Mobile:   user.Mobile,
		Age:      user.Age,
	})
	return nil
}

func (s *sUser) UserInfo(ctx context.Context) (out model.UserInfoOut, err error) {
	var outUser *model.UserOutItem
	user := service.Session().GetUser(ctx)
	err = gconv.Struct(user, &outUser)
	if err != nil {
		return out, err
	}
	return model.UserInfoOut{User: outUser}, err
}

func (s *sUser) SignOut(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}

func (s *sUser) IsUserNameAvailable(ctx context.Context, userName string) (b bool, err error) {
	count, err := dao.Users.Ctx(ctx).
		Where(do.Users{UserName: userName}).
		Count()
	if err != nil {
		return false, err
	}
	return count == 0, err
}

func (s *sUser) IsMobileAvailable(ctx context.Context, mobile string) (b bool, err error) {
	count, err := dao.Users.Ctx(ctx).
		Where(do.Users{Mobile: mobile}).
		Count()
	if err != nil {
		return false, err
	}
	return count == 0, err
}

func (s *sUser) EncryptPassword(userName string, password string) string {
	return gmd5.MustEncrypt(userName + password)
}
