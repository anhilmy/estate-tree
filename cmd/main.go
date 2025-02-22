package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Use(logger)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}

func logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log request body
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			log.Printf("Failed to read request body: %v", err)
		}
		// Re-create the body to be able to read it again
		c.Request().Body = io.NopCloser(bytes.NewReader(reqBody))
		log.Printf("Endpoint: %s:%s ;Request Body: %s", c.Request().Method, c.Request().URL, reqBody)

		// Capture the original response writer
		rec := &responseRecorder{ResponseWriter: c.Response().Writer, body: new(bytes.Buffer)}
		c.Response().Writer = rec

		// Call the next handler
		err = next(c)

		// Log the response body
		if err != nil {
			log.Printf("Error processing request: %v", err)
		} else {
			log.Printf("Status: %d, Response Body: %s", c.Response().Status, rec.body.String())
		}

		return err
	}
}

type responseRecorder struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r *responseRecorder) Write(p []byte) (n int, err error) {
	// Capture the body
	r.body.Write(p)
	return r.ResponseWriter.Write(p)
}
