package usecase

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/syafdia/go-clean-arch/data/entity"
	repo "github.com/syafdia/go-clean-arch/data/repository"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/model"
)

type MockedUserRepository struct {
	mock.Mock
}

func (m *MockedUserRepository) Create(i repo.InputCreateUser) (entity.User, error) {
	args := m.Called(i)

	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockedUserRepository) Get(userID string) (entity.User, error) {
	args := m.Called(userID)

	return args.Get(0).(entity.User), args.Error(1)
}

func (u *MockedUserRepository) Delete(userID string) error {
	return nil
}

type MockedAuthRepository struct {
	mock.Mock
}

func (a *MockedAuthRepository) RefreshToken(token string) (entity.Token, error) {
	return entity.Token{}, nil
}

func (m *MockedAuthRepository) Authenticate(username string, password string) (entity.User, error) {
	args := m.Called(username, password)

	return args.Get(0).(entity.User), args.Error(1)
}

func (a *MockedAuthRepository) SignOut(token string) error {
	return nil
}

func TestAuthUseCase_SignIn(t *testing.T) {
	userRepo := new(MockedUserRepository)
	authRepo := new(MockedAuthRepository)

	type fields struct {
		userRepo repo.UserRepository
		authRepo repo.AuthRepository
	}
	type args struct {
		i model.InputSignIn
	}
	tests := []struct {
		id      int
		name    string
		fields  fields
		args    args
		want    model.ResponseSignIn
		wantErr bool
	}{
		{
			id:   1,
			name: "test 1 success sign in",
			fields: fields{
				userRepo: userRepo,
				authRepo: authRepo,
			},
			args: args{
				i: model.InputSignIn{
					"foousername",
					"foopassword",
				},
			},
			want: model.ResponseSignIn{
				entity.User{
					ID:       "1",
					Username: "Foo Bar",
					Role:     entity.UserRoleDefault,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.id {
			case 1:
				authRepo.On("Authenticate", "foousername", "foopassword").Return(entity.User{
					ID:       "1",
					Username: "Foo Bar",
					Role:     entity.UserRoleDefault,
				}, nil)
			}

			a := &AuthUseCase{
				userRepo: tt.fields.userRepo,
				authRepo: tt.fields.authRepo,
			}

			got, err := a.SignIn(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUseCase.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthUseCase.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
