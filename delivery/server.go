package delivery

import (
	"bank/delivery/controller"
	"bank/manager"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	s.srv.Use(controller.LoggerMiddleware())
	s.srv.Run()

}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()
	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}
}
