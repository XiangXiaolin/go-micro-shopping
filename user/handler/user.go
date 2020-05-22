package handler

import (
	"context"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"
	"golang.org/x/crypto/bcrypt"
	"shopping/user/model"
	"shopping/user/repository"

	user "shopping/user/proto/user"
)

type User struct {
	Repo *repository.User
}

func (this *User) Register(ctx context.Context, req *user.RegisterRequest, rsp *user.Response) error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &model.User{
		Name:     req.User.Name,
		Phone:    req.User.Phone,
		Password: string(hashedPwd),
	}

	// TODO 是否应该验证用户已经存在与否

	if err := this.Repo.Create(newUser); err != nil {
		log.Log("create error")
		return err
	}

	rsp.Code = "200"
	rsp.Msg = "注册成功"

	return nil
}

func (this *User) Login(ctx context.Context, req *user.LoginRequest, rsp *user.Response) error {
	loginUser, err := this.Repo.FindByField("phone", req.Phone, "id, password")
	if err != nil {
		return err
	}

	if loginUser == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该手机号码不存在！")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", "密码错误！")
	}

	rsp.Code = "200"
	rsp.Msg = "登录成功"

	return nil
}

func (this *User) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, rsp *user.Response) error {
	updateUser, err := this.Repo.Find(req.Uid)
	if updateUser == nil {
		return errors.Unauthorized("go.micro.srv.user.login", "该用户不存在")
	}
	if err != nil {
		return err
	}

	// 验证旧密码是否有效
	if err = bcrypt.CompareHashAndPassword([]byte(req.OldPassword), []byte(updateUser.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.Password", "旧密码不正确")
	}

	// 验证通过后，对新密码hash存下来
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	updateUser.Password = string(hashedPwd)
	_, _ = this.Repo.Update(updateUser)

	rsp.Code = "200"
	rsp.Msg = updateUser.Name + "，您的密码修改成功！"
	return err
}
