/*
 domain層はDBアクセスなどの技術的な実装は持たず、それらはinfrastructure層が担当。
*/

package repository

import (
	"context"
	"github.com/takahiko-tanaka-10antz/practice/ddd/domain/model"
)

// UserRepository : User における Repository のインターフェース
// 依存性逆転の法則により infra層は domain層（このインターフェース）に依存
type UserRepository interface {
	GetAll(context.Context) ([]*model.User, error)
}
