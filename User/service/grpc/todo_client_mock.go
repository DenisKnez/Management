package grpc

import (
	context "context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockTodoServiceClient struct {
	mock.Mock
}

func (t *MockTodoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	args := t.Called(in)
	return args.Get(0).(*CreateTodoResponse), args.Error(1)
}

func (t *MockTodoServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (Todo_UploadFileClient, error) {
	return nil, nil
}
