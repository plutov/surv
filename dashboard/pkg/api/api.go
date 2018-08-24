package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

// Server type
type Server struct {
	echo    *echo.Echo
	logger  *logrus.Logger
	config  *Config
	storage *Storage
}

// New creates server instance
func New(logger *logrus.Logger) *Server {
	cfg, err := GetConfig()
	if err != nil {
		logger.WithError(err).Error("unable to read config")
		return nil
	}

	logger.WithField("config", cfg).Info("config")

	return &Server{
		echo:    echo.New(),
		logger:  logger,
		storage: NewStorage(),
	}
}

// Run entrypoint
func (s *Server) Run() {
	s.echo.Use(middleware.CORS())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Gzip())

	s.echo.File("/swagger", "swagger.json")

	s.echo.POST("/request", s.requestDataFetch)
	s.echo.GET("/dashboard", s.getDashboard)

	s.logger.WithError(s.echo.Start(":8080")).Error("unable to start server")
}

func (s *Server) requestDataFetch(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}

func (s *Server) getDashboard(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	if limit <= 0 {
		limit = 10
	}

	rows := s.storage.Get(limit, offset)

	return c.JSON(http.StatusOK, rows)
}
