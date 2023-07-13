package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/localtime"
	"reflect"
	"testing"
)

type controllerMock struct {
	useCase *MockUseCaseInterface
}

func defaultActor() entities.User {
	return entities.User{
		Role:      entities.Role{},
		UpdatedAt: localtime.Now(),
		CreatedAt: localtime.Now(),
	}
}

func defaultControllerMock(t *testing.T) controllerMock {
	return controllerMock{useCase: NewMockUseCaseInterface(t)}
}

func defaultController(mock controllerMock) ControllerInterface {
	return controller{actorUseCase: mock.useCase}
}

func Test_controller_ChangePassword(t *testing.T) {
	type args struct {
		ctx context.Context
		req ChangePasswordRequest
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx: context.Background(),
				req: ChangePasswordRequest{},
			},
			want:    dto.Success("Success update password", nil),
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
			user := mapChangePasswordToUser(tt.args.req)
			cMock.useCase.EXPECT().validateActor(user, mock.Anything).Return(err, err)
			cMock.useCase.EXPECT().ChangePassword(tt.args.ctx, tt.args.req.OldPassword, user).
				Return(err, err)
			got, err := c.ChangePassword(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangePassword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_CreateActor(t *testing.T) {
	req := CreateRequest{}
	actor := mapCreateRequestToActor(req)
	actorResult := defaultActor()
	respActor := mapActorToResponse(actorResult)
	type args struct {
		ctx context.Context
		req CreateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx: context.Background(),
				req: req,
			},
			want:    dto.Created("Success create user", respActor),
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
			cMock.useCase.EXPECT().validateActor(actor, mock.Anything).Return(err, err)
			cMock.useCase.EXPECT().CreateUser(tt.args.ctx, actor).Return(actorResult, err)
			got, err := c.CreateActor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateActor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_DeleteActor(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
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
			fmt.Println(err)
			cMock.useCase.EXPECT().DeleteUser(tt.args.ctx, tt.args.username).Return(err)
			if err := c.DeleteActor(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteActor() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_controller_GetAll(t *testing.T) {
	totalRow := 10
	actorResp := make([]Representation, 0)
	type args struct {
		ctx context.Context
		req PaginationRequest
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx: context.Background(),
				req: PaginationRequest{},
			},
			want:    dto.SuccessPagination("Success retrieve user", 0, 0, totalRow, actorResp),
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
			offset := (tt.args.req.Page - 1) * tt.args.req.PerPage
			actors := make([]entities.User, 0)
			cMock.useCase.EXPECT().GetAll(tt.args.ctx, tt.args.req.Username+"%", tt.args.req.Role,
				uint(tt.args.req.PerPage), uint(offset)).Return(actors, err)
			cMock.useCase.EXPECT().CountAll(tt.args.ctx, tt.args.req.Username+"%", tt.args.req.Role).
				Return(totalRow, err)
			got, err := c.GetAll(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_GetByUsername(t *testing.T) {
	actor := defaultActor()
	actorResponse := mapActorToResponse(actor)
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			want:    dto.Success("Success retrieve user", actorResponse),
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
			cMock.useCase.EXPECT().GetByUsername(tt.args.ctx, tt.args.username).Return(actor, err)
			got, err := c.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_UpdateActor(t *testing.T) {
	actorRes := defaultActor()
	type args struct {
		ctx context.Context
		req UpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx: context.Background(),
				req: UpdateRequest{},
			},
			want:    dto.Success("Success update user", mapActorToResponse(actorRes)),
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
			actor := mapUpdateRequestToActor(tt.args.req)
			cMock.useCase.EXPECT().validateActor(actor, mock.Anything).Return(err, err)
			cMock.useCase.EXPECT().UpdateUser(tt.args.ctx, actor).Return(actorRes, err)
			got, err := c.UpdateActor(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateActor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_UpdateProfile(t *testing.T) {
	actorResp := defaultActor()
	type args struct {
		ctx context.Context
		req UpdateProfile
	}
	tests := []struct {
		name    string
		args    args
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx: context.Background(),
				req: UpdateProfile{},
			},
			want:    dto.Success("Success update user", mapActorToResponse(actorResp)),
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
			user := mapUpdateProfileToUser(tt.args.req)
			cMock.useCase.EXPECT().UpdateUser(tt.args.ctx, user).Return(actorResp, nil)
			got, err := c.UpdateProfile(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateProfile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
