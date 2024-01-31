package delivery

import (
	"fmt"
	"simple-bank/config"
	"simple-bank/delivery/controller"
	"simple-bank/delivery/middleware"
	"simple-bank/manager"
	"simple-bank/utils/common"

	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager  manager.UsecaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
	session    common.Session
}

func (s *Server) setupController() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.session)
	rg := s.engine.Group("/api")
	// register all controller here
	controller.NewUserController(s.ucManager.UserUseCase(), rg).Route()
	controller.NewTransactionController(s.ucManager.TransactionUseCase(), s.ucManager.UserUseCase(), rg, authMiddleware, s.session).Route()
	controller.NewAuthController(s.ucManager.AuthUseCase(), rg, s.session).Route()
}

func (s *Server) Run() {
	s.setupController()
	if err := s.engine.Run(s.host); err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUsecaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.FileConfig)
	session := common.NewSession(cfg.SessionConfig)

	return &Server{
		ucManager:  useCaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
		session:    session,
	}
}
