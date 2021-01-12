package usecase

import (
	"context"
	"github.com/takahiko-tanaka-10antz/practice/ddd/domain/model"
	"github.com/takahiko-tanaka-10antz/practice/ddd/domain/repository"
)

// UserUseCase : User における UseCase のインターフェース
type UserUseCase interface {
	GetAll(context.Context) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase : User データに関する UseCase を生成
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

// GetAll : User データを全件取得するためのユースケース
func (uu userUseCase) GetAll(ctx context.Context) (users []*model.User, err error) {
	// Persistence（Repository）を呼出
	users, err = uu.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
