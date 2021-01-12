package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/takahiko-tanaka-10antz/practice/ddd/handler/rest"
	"github.com/takahiko-tanaka-10antz/practice/ddd/infra/persistence"
	"github.com/takahiko-tanaka-10antz/practice/ddd/usecase"
	"net/http"
)

func main() {
	// DBなどのアクセスの実体を持っているuserPersistence（repository.UserRepositoryを満たす）を生成する。
	userPersistence := persistence.NewUserPersistence()
	// そのuserPersistenceをusecase層のuserUsecase（repository.UserRepositoryをフィールドに持つ）に注入する。
	userUseCase := usecase.NewUserUseCase(userPersistence)
	// 生成したuserUsecaseをuserHandler（userUsecaseをフィールドに持つ）に注入する
	userHandler := handler.NewUserHandler(userUseCase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", userHandler.Index)
	_ = http.ListenAndServe(":5050", r)
}
