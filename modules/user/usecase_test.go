package user

import (
	"context"
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories/mocks"
	mocksCrypt "nashrul-be/crm/utils/crypto/mocks"
	"reflect"
	"testing"
)

type useCaseMock struct {
	actorRepository *mocks.ActorRepositoryInterface
	roleRepository  *mocks.RoleRepositoryInterface
	hash            *mocksCrypt.Hash
}

func defaultUser() *entities.User {
	return &entities.User{
		Username: "admin",
	}
}

func defaultMock(t *testing.T) useCaseMock {
	return useCaseMock{
		actorRepository: mocks.NewActorRepositoryInterface(t),
		roleRepository:  mocks.NewRoleRepositoryInterface(t),
		hash:            mocksCrypt.NewHash(t),
	}
}

func defaultUseCase(mock useCaseMock) useCase {
	return useCase{
		actorRepository: mock.actorRepository,
		roleRepository:  mock.roleRepository,
		hash:            mock.hash,
	}
}

func Test_useCase_ChangePassword(t *testing.T) {
	type args struct {
		ctx         context.Context
		oldPassword string
		user        entities.User
	}
	tests := []struct {
		name       string
		args       args
		wantResult error
		wantErr    error
	}{
		{
			name: "Normal case",
			args: args{
				ctx:         context.Background(),
				oldPassword: "12",
				user:        entities.User{Password: "12"},
			},
			wantResult: nil,
			wantErr:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			ucMock.actorRepository.EXPECT().GetByUsername(tt.args.ctx, tt.args.user.Username).Return(tt.args.user, tt.wantErr)
			ucMock.hash.EXPECT().Compare(tt.args.oldPassword, tt.args.user.Password).Return(tt.wantResult)
			ucMock.hash.EXPECT().Hash(tt.args.user.Password).Return(tt.args.user.Password, tt.wantErr)
			ucMock.actorRepository.EXPECT().Update(tt.args.ctx, tt.args.user).Return(tt.wantErr)
			ucMock.actorRepository.EXPECT().GetByUsername(tt.args.ctx, tt.args.user.Username).Return(tt.args.user, tt.wantErr)
			gotResult, gotErr := uc.ChangePassword(tt.args.ctx, tt.args.oldPassword, tt.args.user)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ChangePassword() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("ChangePassword() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_useCase_CountAll(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		role     string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:      context.Background(),
				username: "user",
				role:     "user",
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.actorRepository.EXPECT().CountAll(tt.args.ctx, tt.args.username, tt.args.role).Return(tt.want, err)
			got, err := uc.CountAll(tt.args.ctx, tt.args.username, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_CreateUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		actor entities.User
	}
	tests := []struct {
		name       string
		args       args
		wantResult entities.User
		wantErr    bool
	}{
		{
			name: "Normal Case",
			args: args{
				ctx:   context.Background(),
				actor: entities.User{},
			},
			wantResult: entities.User{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.hash.EXPECT().Hash(tt.args.actor.Password).Return(tt.args.actor.Password, err)
			ucMock.actorRepository.EXPECT().Create(tt.args.ctx, tt.args.actor).Return(tt.wantResult, err)
			ucMock.roleRepository.EXPECT().GetByID(tt.args.actor.RoleID).Return(tt.wantResult.Role, err)
			gotResult, err := uc.CreateUser(tt.args.ctx, tt.args.actor)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CreateUser() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_useCase_DeleteUser(t *testing.T) {
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
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.actorRepository.EXPECT().Delete(tt.args.ctx, tt.args.username).Return(err)
			if err := uc.DeleteUser(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		role     string
		limit    uint
		offset   uint
	}
	tests := []struct {
		name    string
		args    args
		want    []entities.User
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:      context.Background(),
				username: "user",
				role:     "user",
				limit:    10,
				offset:   0,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.actorRepository.EXPECT().GetAll(tt.args.ctx, tt.args.username,
				tt.args.role, tt.args.limit, tt.args.offset).Return(tt.want, err)
			got, err := uc.GetAll(tt.args.ctx, tt.args.username, tt.args.role, tt.args.limit, tt.args.offset)
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

func Test_useCase_GetByUsername(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name      string
		args      args
		wantActor entities.User
		wantErr   bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:      context.Background(),
				username: "user",
			},
			wantActor: entities.User{},
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.actorRepository.EXPECT().GetByUsername(tt.args.ctx, tt.args.username).Return(tt.wantActor, err)
			gotActor, err := uc.GetByUsername(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotActor, tt.wantActor) {
				t.Errorf("GetByUsername() gotActor = %v, want %v", gotActor, tt.wantActor)
			}
		})
	}
}

func Test_useCase_UpdateUser(t *testing.T) {
	type args struct {
		ctx   context.Context
		actor entities.User
	}
	tests := []struct {
		name       string
		args       args
		wantResult entities.User
		wantErr    bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:   context.Background(),
				actor: entities.User{},
			},
			wantResult: entities.User{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			if tt.args.actor.Password != "" {
				ucMock.hash.EXPECT().Hash(tt.args.actor.Password).Return(tt.args.actor.Password, err)
			}
			ucMock.actorRepository.EXPECT().Update(tt.args.ctx, tt.args.actor).Return(err)
			ucMock.actorRepository.EXPECT().GetByUsername(tt.args.ctx, tt.args.actor.Username).Return(tt.wantResult, err)
			gotResult, err := uc.UpdateUser(tt.args.ctx, tt.args.actor)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("UpdateUser() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
