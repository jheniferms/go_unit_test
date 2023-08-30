package book

import (
	"context"
	"errors"
	"fmt"
	"go_unit_test/src/entities"
	mocks "go_unit_test/src/tests/mocks/gateways/book"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type getUsecaseSuite struct {
	suite.Suite
	repository *mocks.RepositoryMock
	ctx        context.Context
	usecase    GetUsecase
}

func (s *getUsecaseSuite) SetupTest() {
	s.repository = new(mocks.RepositoryMock)
	s.ctx = context.Background()
	s.usecase = New(s.repository)
}

func TestGetUsecaseSuite(t *testing.T) {
	suite.Run(t, &getUsecaseSuite{})
}

func (s *getUsecaseSuite) TestExecute_Success() {

	filter := entities.BookFillter{
		Title: "the colors",
	}

	responseExpected := entities.Book{
		Title:  "the colors",
		Author: "J. April",
	}

	s.repository.On("Get", filter).Return(responseExpected, nil)

	response, err := s.usecase.Execute(s.ctx, filter)
	s.NoError(err)
	s.Equal(responseExpected, response)

	s.repository.AssertNumberOfCalls(s.T(), "Get", 1)
}

func (s *getUsecaseSuite) TestExecute_Success_Multiple() {

	testsCase := []struct {
		title string
	}{
		{title: "the colors"},
		{title: "the animals"},
	}

	for _, testCase := range testsCase {
		name := fmt.Sprintf("TesteExecute_Success_%s", testCase.title)

		s.T().Run(name, func(t *testing.T) {
			filter := entities.BookFillter{
				Title: testCase.title,
			}

			responseExpected := entities.Book{
				Title:  testCase.title,
				Author: "J. April",
			}

			s.repository.On("Get", filter).Return(responseExpected, nil)

			response, err := s.usecase.Execute(s.ctx, filter)
			s.NoError(err)
			s.Equal(responseExpected, response)
		})
	}
}

func (s *getUsecaseSuite) TestExecute_Error() {

	filter := entities.BookFillter{
		Title: "the colors",
	}

	errorFoo := errors.New("error")
	s.repository.On("Get", mock.Anything).Return(entities.Book{}, errorFoo)

	response, err := s.usecase.Execute(s.ctx, filter)
	s.Error(err)
	s.ErrorIs(err, errorFoo)
	s.Empty(response)
	s.repository.AssertNumberOfCalls(s.T(), "Get", 1)
}
