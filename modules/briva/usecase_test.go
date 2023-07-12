package briva

import (
	"context"
	"errors"
	"github.com/adjust/rmq/v5"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories/mocks"
	mocksTransactor "nashrul-be/crm/utils/db/mocks"
	"reflect"
	"testing"
)

type useCaseMock struct {
	brivaRepo *mocks.BrivaRepositoryInterface
	auditRepo *mocks.AuditRepositoryInterface
	queue     rmq.Queue
}

func defaultMock(t *testing.T) useCaseMock {
	return useCaseMock{
		brivaRepo: mocks.NewBrivaRepositoryInterface(t),
		auditRepo: mocks.NewAuditRepositoryInterface(t),
		queue:     rmq.NewTestQueue("mock-queue"),
	}
}

func defaultUseCase(mock useCaseMock) useCase {
	return useCase{
		brivaRepo: mock.brivaRepo,
		auditRepo: mock.auditRepo,
		queue:     mock.queue,
	}
}

func Test_useCase_GetByBrivaNo(t *testing.T) {
	type args struct {
		ctx     context.Context
		brivano string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.Briva
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:     context.Background(),
				brivano: "12121",
			},
			want:    entities.Briva{},
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
			ucMock.brivaRepo.EXPECT().GetByBrivaNo(tt.args.ctx, tt.args.brivano).Return(tt.want, err)
			got, err := uc.GetByBrivaNo(tt.args.ctx, tt.args.brivano)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByBrivaNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByBrivaNo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_Update(t *testing.T) {
	type args struct {
		ctx   context.Context
		briva entities.Briva
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:   context.Background(),
				briva: entities.Briva{},
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
			auditEnt := entities.Audit{}
			ucMock.brivaRepo.EXPECT().MakeAuditUpdate(tt.args.ctx, tt.args.briva).Return(auditEnt, err)
			tx := mocksTransactor.NewTransactor(t)
			ucMock.brivaRepo.EXPECT().Begin().Return(tx)
			ucMock.auditRepo.EXPECT().Begin().Return(tx)
			ucMock.brivaRepo.EXPECT().New(tx).Return(uc.brivaRepo)
			ucMock.auditRepo.EXPECT().New(tx).Return(uc.auditRepo)
			ucMock.brivaRepo.EXPECT().Update(tt.args.ctx, tt.args.briva).Return(err)
			ucMock.auditRepo.EXPECT().Create(auditEnt).Return(err)
			tx.EXPECT().Commit().Return(&gorm.DB{}).Times(2)
			if err := uc.Update(tt.args.ctx, tt.args.briva); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
