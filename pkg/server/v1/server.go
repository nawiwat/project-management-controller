package server

import (
	"app-controller/pkg/middlewares"
	"app-controller/pkg/migrations"
	Users "app-controller/pkg/repositories/users"
	Projects "app-controller/pkg/repositories/projects"
	Tasks "app-controller/pkg/repositories/tasks"
	"app-controller/pkg/services/contlr"
	"net"


	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	echo   *echo.Echo
	config *Config
}

func getSQLDB(driver, conn string) (*gorm.DB, error) {
	switch driver {
	case "postgres":
		return gorm.Open(postgres.Open(conn))
	case "sqlite":
		return gorm.Open(sqlite.Open(conn))
	default:
		return gorm.Open(postgres.Open(conn))
	}
}

func (s *Server) startTLS(addr, cert, key string) error {
	if err := s.echo.StartTLS(addr, cert, key); err != nil {
		return errors.Wrapf(err, "fail to start server with")
	}
	return nil
}

func (s *Server) start(addr string) error {
	if err := s.echo.Start(addr); err != nil {
		return errors.Wrapf(err, "fail to start server with")
	}
	return nil
}

// Start to start server
func (s *Server) Start() {
	logger := logrus.StandardLogger()

	// Server
	addr := net.JoinHostPort(s.config.Server.Host, s.config.Server.Port)
	if s.config.Server.CertFile != "" && s.config.Server.KeyFile != "" {
		go func() {
			if err := s.startTLS(addr, s.config.Server.CertFile, s.config.Server.KeyFile); err != nil {
				logger.Fatalf("error on start server: %s", err.Error())
			}
		}()
	} else {
		go func() {
			if err := s.start(addr); err != nil {
				logger.Fatalf("error on start server: %s", err.Error())
			}
		}()
	}

	logger.Infof("start server on %s", addr)
}

func NewServer() (*Server, error) {
	godotenv.Load()

	filePath := ""

	if filePath == "" {
		filePath = "config.toml"
	}

	c := LoadConfig(filePath)

	// set log
	logLevel, err := logrus.ParseLevel(c.Log.Level)
	if err != nil {
		return nil, errors.Wrap(err, "fail to parse log level")
	}
	logrus.SetLevel(logLevel)

	// connect to db
	db, err := getSQLDB(c.SQL.Driver, c.SQL.Connection)
	if err != nil {
		return nil, errors.Wrap(err, "fail to open db")
	}

	if err := migrations.Migrate(db); err != nil {
		return nil, err
	}

	s := &Server{
		echo:   echo.New(),
		config: c,
	}

	e := s.echo
	e.Validator = &customValidator{validator: validator.New()}

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	middlewares.RegisterHTTPErrorHandling(e)

	// repositories
	UsersRepo := Users.New(db)
	ProjectsRepo := Projects.New(db)
	TasksRepo := Tasks.New(db)

	// services
	contlrsvc := contlr.NewControllerService(UsersRepo,ProjectsRepo,TasksRepo)

	registerHealthCheckRouteV1(s.echo)
	registerRouterV1(
		s.echo,
		contlrsvc,
	)

	return s, nil
}