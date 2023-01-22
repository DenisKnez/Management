package http

import (
	"fmt"
	"net/http"

	"github.com/DenisKnez/management/todo/api/service"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService service.Service
	Logger      echo.Logger
}

// CreateTodo godoc
// @Summary      Create todo
// @Description  Create todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo   body      CreateTodoRequest  true  "Text"
// @Success      201 CreateTodoResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /todos [post]
func (t *TodoHandler) CreateTodo(c echo.Context) error {
	var req CreateTodoRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Reponse{
			Error: "invalid json body",
		})
	}

	err = t.TodoService.CreateTodo(c.Request().Context(), req.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Reponse{
			Error: "failed to create todo",
		})
	}

	return c.JSON(http.StatusCreated, Reponse{})
}

// UpdateTodo godoc
// @Summary      Update todo
// @Description  Update todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id"
// @Param        todo   body      UpdateTodoRequest  true  "Text"
// @Success      200 {object} UpdateTodoResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /todos/{id} [post]
func (t *TodoHandler) UpdateTodo(c echo.Context) error {
	// bind todo body
	var req UpdateTodoRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Reponse{
			Error: "invalid request",
		})
	}

	// update todo
	err = t.TodoService.UpdateTodo(c.Request().Context(), service.Todo{
		ID:   req.ID,
		Text: req.Text,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Reponse{
			Error: fmt.Sprintf("failed to update todo %v", req.ID),
		})
	}

	return c.JSON(http.StatusOK, Reponse{})
}

// DeleteTodo godoc
// @Summary      Delete todo
// @Description  Delete todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id"
// @Success      200 {object} UpdateTodoResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /todos/{id} [delete]
func (t *TodoHandler) DeleteTodo(c echo.Context) error {
	var req DeleteTodoRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Reponse{
			Error: "invalid json body",
		})
	}

	err = t.TodoService.DeleteTodo(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Reponse{
			Error: fmt.Sprintf("failed to delete todo %v", req.ID),
		})
	}

	return c.JSON(http.StatusOK, Reponse{})
}

// GetTodo godoc
// @Summary      Get todo
// @Description  Get todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Id"
// @Success      200 {object} GetTodoResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /todos/{id} [get]
func (t *TodoHandler) GetTodo(c echo.Context) error {
	// bind todo id
	var req GetTodoRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Reponse{
			Error: "invalid json body",
		})
	}

	todo, err := t.TodoService.GetTodo(c.Request().Context(), req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Reponse{
			Error: fmt.Sprintf("failed to get todo %v", req.ID),
		})
	}

	return c.JSON(http.StatusOK, Reponse{
		Data: GetTodoResponse{
			ID:   todo.ID,
			Text: todo.Text,
		},
	})
}
