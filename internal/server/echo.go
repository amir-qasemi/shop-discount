package server

import (
	"net/http"

	"github.com/amir-qasemi/shop-discount/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server responsible for serving discount requests
type Server struct {
	Echo *echo.Echo
}

// Controller each package wanting to define request handlers should implement this interface
type Controller interface {
	// Setup controllers should set up their routes, middlewares and ... in this method
	Setup(*Server)
}

// New build a new server
func New(serverConfig config.ServerConfig) *Server {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	server := Server{Echo: e}
	return &server
}

// HttpMethod methods which echo server accepts
type HttpMethod string

// Methods copied from echo itself
const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

// Wrapper a utility function to be able to bind sent parameters to the handler
// as golang struct type and validate them.
func Wrapper[A any](s *Server, path string, httpMethod HttpMethod, f func(a A, ctx echo.Context) error) {
	s.Echo.Add(string(httpMethod), path, func(ctx echo.Context) error {
		// Bind
		a := new(A)
		if err := ctx.Bind(a); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Validate
		if err := ctx.Validate(a); err != nil {
			return err
		}

		return f(*a, ctx)
	})
}

// GroupWrapper a utility function to be able to bind sent parameters to the handler
// in a group as golang struct type and validate them.
func GroupWrapper[A any](g *echo.Group, path string, httpMethod HttpMethod, f func(a A, ctx echo.Context) error) {
	g.Add(string(httpMethod), path, func(ctx echo.Context) error {
		// Bind
		a := new(A)
		if err := ctx.Bind(a); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Validate
		if err := ctx.Validate(a); err != nil {
			return err
		}

		return f(*a, ctx)
	})
}
