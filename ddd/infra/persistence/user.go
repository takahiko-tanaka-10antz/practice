package persistence

import (
	"context"
	"github.com/takahiko-tanaka-10antz/practice/ddd/domain/model"
	"github.com/takahiko-tanaka-10antz/practice/ddd/domain/repository"
	"time"
)

type userPersistence struct{}

// NewUserPersistence : User データに関する Persistence を生成
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

// GetAll : DB から User データを全件取得（UserRepository インターフェースの GetAll() を実装したもの）
//  -> 本来は DB からデータを取得する（モックデータを返却）
func (up userPersistence) GetAll(context.Context) ([]*model.User, error) {
	user1 := model.User{}
	user1.Id = 1
	user1.PlayerId = "01010101"
	user1.Ap = 10
	user1.Lp = 5
	user1.LoginAt = time.Now().Add(-24 * time.Hour)

	user2 := model.User{}
	user2.Id = 2
	user2.PlayerId = "02020202"
	user2.Ap = 10
	user2.Lp = 4
	user2.LoginAt = time.Now().Add(-24 * 7 * time.Hour)

	return []*model.User{&user1, &user2}, nil
}
