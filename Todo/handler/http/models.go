package http

import "github.com/gofrs/uuid"

// GENERAL
type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}

// CREATE TODO MODELS
type CreateTodoRequest struct {
	Text string `json:"text" example:"some text"`
}

type CreateTodoResponse struct{}

// UPDATE TODO MODELS
type UpdateTodoRequest struct {
	ID   uuid.UUID `json:"-" param:"id" example:"4b6158b9-bc03-4058-93b8-a0a1fcab7371"`
	Text string    `json:"text" example:"some text"`
}

type UpdateTodoResponse struct{}

// DELETE TODO MODELS
type DeleteTodoRequest struct {
	ID uuid.UUID `param:"id" example:"4b6158b9-bc03-4058-93b8-a0a1fcab7371"`
}

type DeleteTodoResponse struct{}

// GET TODO MODELS
type GetTodoRequest struct {
	ID uuid.UUID `param:"id" example:"4b6158b9-bc03-4058-93b8-a0a1fcab7371"`
}

type GetTodoResponse struct {
	ID   uuid.UUID `json:"id" example:"4b6158b9-bc03-4058-93b8-a0a1fcab7371"`
	Text string    `json:"text" example:"some text"`
}
