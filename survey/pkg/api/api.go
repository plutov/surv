package api

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

// Server type
type Server struct {
	echo    *echo.Echo
	logger  *logrus.Logger
	storage *Storage
}

// New creates server instance
func New(logger *logrus.Logger) *Server {
	return &Server{
		echo:    echo.New(),
		logger:  logger,
		storage: NewStorage(),
	}
}

// Run entrypoint
func (s *Server) Run() {
	name := os.Getenv("NAME")

	s.logger.Infof("Starting %s", name)

	s.echo.Use(middleware.CORS())
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Gzip())

	s.echo.File("/swagger", "swagger.json")

	s.echo.POST("/answers", s.postAnswers)
	s.echo.GET("/answers", s.getAnswers)

	s.logger.WithError(s.echo.Start(":8080")).Error("unable to start server")
}

func (s *Server) postAnswers(c echo.Context) error {
	var answer *Answer
	err := c.Bind(&answer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	s.storage.Save(*answer)

	return c.JSON(http.StatusCreated, nil)
}

func (s *Server) getAnswers(c echo.Context) error {
	return c.JSON(http.StatusOK, s.storage.Get())
}
