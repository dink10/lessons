package gomock

import (
	"testing"

	mocks "github.com/dink10/lessons/mocks/mock"
	"github.com/dink10/lessons/mocks/user"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserWork(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testuser := &user.User{Doer: mockDoer}

	mockDoer.EXPECT().Do(gomock.Any(), gomock.Any()).Return("done").Times(1)

	res := testuser.Use()

	assert.Equal(t, "done", res)
}
