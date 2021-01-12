package handler

import (
	"encoding/json"
	"github.com/takahiko-tanaka-10antz/practice/ddd/usecase"
	"net/http"
	"time"
)

// UserHandler : User における Handler のインターフェース
type UserHandler interface {
	Index(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

// NewUserUseCase : User データに関する Handler を生成
func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

// UserIndex : GET /users -> User データ一覧を返す
func (uu userHandler) Index(w http.ResponseWriter, r *http.Request) {
	// userField : response 内で使用する User を表す構造体
	type userField struct {
		Id       int64     `json:"id"`
		PlayerId string    `json:"player_id"`
		Ap       int       `json:"ap"`
		Lp       int       `json:"lp"`
		LoginAt  time.Time `json:"login_at"`
	}

	// response : API のレスポンス
	type response struct {
		Users []userField `json:"users"`
	}

	ctx := r.Context()

	// ユースケースの呼出
	users, err := uu.userUseCase.GetAll(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(response)
	for _, user := range users {
		var bf userField
		bf = userField(*user)
		res.Users = append(res.Users, bf)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
