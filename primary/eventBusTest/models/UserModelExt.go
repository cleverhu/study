package models

import "github.com/golang/protobuf/proto"

type UserOpt func(*UserModel)

type UserOpts []UserOpt

func (this *UserModel) Encode() ([]byte, error) {
	return proto.Marshal(this)
}

func (this *UserModel) Decode(b []byte) error {
	return proto.Unmarshal(b, this)
}

func WithUserID(id int32) UserOpt {
	return func(this *UserModel) {
		this.UserId = id
	}
}

func WithUserName(name string) UserOpt {
	return func(this *UserModel) {
		this.UserName = name
	}
}

func (this UserOpts) apply(u *UserModel) {
	for _, f := range this {
		f(u)
	}
}

func NewUserModel(opts ...UserOpt) *UserModel {
	u := &UserModel{}
	UserOpts(opts).apply(u)
	return u
}
