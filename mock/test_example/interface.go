package chat

import (
	"github.com/stretchr/testify/mock"
)

type MockHandler struct {
	mock.Mock
}

type MockUsecase struct {
	mock.Mock
}

type MockRepository struct {
	mock.Mock
}
