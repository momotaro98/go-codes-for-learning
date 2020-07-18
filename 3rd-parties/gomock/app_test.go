package app

import (
	"testing"

	"github.com/golang/mock/gomock"
)

//go:generate mockgen -source=app.go -destination=mock.go -package=app

func TestRegisterUser(t *testing.T) {
	// Actual mapped UserModel
	var actUserModel *UserModel

	// Mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userRepoMock := NewMockIUserRepository(mockCtrl)

	userRepoMock.EXPECT().
		InsertAUser(gomock.Any()).
		Do(func(um *UserModel) { // Do を使ってモック関数への引数を得ることができる
			actUserModel = um
		}).
		Return(nil)

	us := &userService{
		userRepo: userRepoMock,
	}

	// Input
	input := User{
		UserID: 123,
		Name:   "John",
		Email:  "john@example.com",
	}

	// Act
	_ = us.RegisterUser(&input)

	// Assert to test mapping
	if uint32(actUserModel.ID) != input.UserID {
		t.Error()
	}
	if name, _ := actUserModel.Name.Value(); name != input.Name {
		t.Error()
	}
	if actUserModel.Email != input.Email {
		t.Error()
	}
}
