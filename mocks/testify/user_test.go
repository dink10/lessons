package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/dink10/lessons/mocks/testify/mocks"
	"github.com/dink10/lessons/mocks/testify/user"
)

func TestUserTestify(t *testing.T) {
	mockDoer := &mocks.Doer{}

	testUser := user.User{Doer: mockDoer}

	mockDoer.On("Do", 1, "work1").Return("done").Once()

	res := testUser.Use()

	assert.Equal(t, "done", res)
}
