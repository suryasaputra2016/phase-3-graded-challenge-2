package repos

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
)

// define mock user repository
type UserRepoMock struct {
	Mock mock.Mock
}

func (m *UserRepoMock) GetUserByUsername(id int) (*entities.User, error) {
	res := m.Mock.Called(id)

	if res.Get(0) == nil {
		return nil, errors.New("user id not found")
	}

	user := res.Get(0).(entities.User)
	return &user, nil
}
