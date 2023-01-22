package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DenisKnez/management/todo/api/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	e := echo.New()

	body := `{"text":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()

	c := e.NewContext(req, w)

	todoService := new(service.MockTodoService)

	todoService.On("CreateTodo", service.Todo{
		Text: "test",
	}).Return(nil)

	handler := TodoHandler{
		TodoService: todoService,
		Logger:      e.Logger,
	}

	err := handler.CreateTodo(c)

	todoService.AssertExpectations(t)

	require.Nil(t, err)
	require.Equal(t, http.StatusCreated, w.Code)
}
