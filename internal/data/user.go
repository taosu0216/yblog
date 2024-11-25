package data

import (
	v1 "blug/api/blug/v1"
	"blug/internal/data/ent"
	"blug/internal/data/ent/user"
	"blug/internal/pkg"
	"context"
)

func (f *blugRepo) RegisterUserInData(ctx context.Context, req *v1.RegisterUserReq) error {
	err := f.data.DB.User.Create().
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetIsRoot(false).Exec(ctx)

	if err != nil {
		f.log.Error(err)
		return pkg.InternalErr
	}

	f.data.UserCache.RPush(ctx, pkg.UserListKey, req.Username)

	return nil
}

func (f *blugRepo) CheckUserInDB(ctx context.Context, req *v1.UserLoginReq) (bool, error) {
	info, err := f.data.DB.User.Query().
		Where(user.UsernameEQ(req.Username)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, pkg.UserNotFoundErr
		}
		f.log.Error(err)
		return false, pkg.InternalErr
	}
	if info.Password == req.Password {
		return true, nil
	} else {
		return false, nil
	}
}

func (f *blugRepo) GetUserListInData(ctx context.Context) ([]string, error) {
	result, err := f.data.UserCache.LRange(ctx, pkg.UserListKey, 0, -1).Result()
	if err != nil {
		f.log.Error(err)
	} else {
		f.log.Info("get user list shoot cache succeed!")
		return result, nil
	}

	list, err := f.data.DB.User.Query().All(ctx)
	if err != nil {
		f.log.Error(err)
		return nil, pkg.InternalErr
	}
	users := make([]string, len(list)-1)
	for _, v := range list {
		users = append(users, v.Username)
	}
	return users, nil
}
