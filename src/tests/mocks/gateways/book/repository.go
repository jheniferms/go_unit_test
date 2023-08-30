package mocks

import (
	"context"
	"go_unit_test/src/entities"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Get(_ context.Context, filter entities.BookFillter) (entities.Book, error) {
	args := m.Called(filter)
	return args.Get(0).(entities.Book), args.Error(1)
}
