package user

import (
	"context"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/repositories/mocks"
	"nashrul-be/crm/utils/crypto"
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
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx      context.Context
		username string
		role     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
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
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx   context.Context
		actor entities.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult entities.User
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
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
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
			if err := uc.DeleteUser(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx      context.Context
		username string
		role     string
		limit    uint
		offset   uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
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
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantActor entities.User
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
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
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		ctx   context.Context
		actor entities.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult entities.User
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
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

func Test_useCase_validateActor(t *testing.T) {
	type fields struct {
		actorRepository repositories.ActorRepositoryInterface
		roleRepository  repositories.RoleRepositoryInterface
		hash            crypto.Hash
	}
	type args struct {
		actor       entities.User
		validations []validateFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCase{
				actorRepository: tt.fields.actorRepository,
				roleRepository:  tt.fields.roleRepository,
				hash:            tt.fields.hash,
			}
			got, got1 := uc.validateActor(tt.args.actor, tt.args.validations...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validateActor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateActor() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
