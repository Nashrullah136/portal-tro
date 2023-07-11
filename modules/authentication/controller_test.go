package authentication

import (
	"context"
	"errors"
	"nashrul-be/crm/entities"
	mocksAudit "nashrul-be/crm/modules/audit/mocks"
	"nashrul-be/crm/modules/user"
	"nashrul-be/crm/utils"
	mocksCrypt "nashrul-be/crm/utils/crypto/mocks"
	"reflect"
	"testing"
)

type controllerMock struct {
	actorUseCase *user.MockUseCaseInterface
	auditUseCase *mocksAudit.UseCaseInterface
	hash         *mocksCrypt.Hash
}

func defaultControllerMock(t *testing.T) controllerMock {
	return controllerMock{
		actorUseCase: user.NewMockUseCaseInterface(t),
		auditUseCase: mocksAudit.NewUseCaseInterface(t),
		hash:         mocksCrypt.NewHash(t),
	}
}

func defaultController(mock controllerMock) ControllerInterface {
	return controller{
		actorUseCase: mock.actorUseCase,
		auditUseCase: mock.auditUseCase,
		hash:         mock.hash,
	}
}

func Test_controller_Login(t *testing.T) {
	actor := entities.User{}
	type args struct {
		request LoginRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.User
		wantErr bool
	}{
		{
			name:    "Normal case",
			args:    args{request: LoginRequest{}},
			want:    &actor,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cMock := defaultControllerMock(t)
			c := defaultController(cMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			cMock.actorUseCase.EXPECT().GetByUsername(nil, tt.args.request.Username).Return(actor, err)
			cMock.hash.EXPECT().Compare(tt.args.request.Password, actor.Password).Return(err)
			ctx := utils.SetUserToContext(context.Background(), actor)
			cMock.auditUseCase.EXPECT().CreateAudit(ctx, "Login").Return(err)
			got, err := c.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_Logout(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Normal case",
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cMock := defaultControllerMock(t)
			c := defaultController(cMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			cMock.auditUseCase.EXPECT().CreateAudit(tt.args.ctx, "Logout").Return(err)
			if err := c.Logout(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
