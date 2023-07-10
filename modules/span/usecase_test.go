package span

import (
	"context"
	"errors"
	"github.com/adjust/rmq/v5"
	"gorm.io/gorm"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories/mocks"
	"nashrul-be/crm/utils"
	mocksTransactor "nashrul-be/crm/utils/db/mocks"
	"reflect"
	"testing"
)

type useCaseMock struct {
	spanRepo  *mocks.SpanRepositoryInterface
	auditRepo *mocks.AuditRepositoryInterface
	queue     rmq.Queue
}

func defaultUser() *entities.User {
	return &entities.User{
		Username: "admin",
	}
}

func defaultMock(t *testing.T) useCaseMock {
	return useCaseMock{
		spanRepo:  mocks.NewSpanRepositoryInterface(t),
		auditRepo: mocks.NewAuditRepositoryInterface(t),
		queue:     rmq.NewTestQueue("mock-queue"),
	}
}

func defaultUseCase(mock useCaseMock) useCase {
	return useCase{
		spanRepo:  mock.spanRepo,
		auditRepo: mock.auditRepo,
		queue:     mock.queue,
	}
}

func Test_useCase_GetByDocumentNumber(t *testing.T) {
	type args struct {
		ctx            context.Context
		documentNumber string
	}
	tests := []struct {
		name    string
		args    args
		want    entities.SPAN
		wantErr bool
	}{
		{
			name: "Normal test case",
			args: args{
				ctx:            context.Background(),
				documentNumber: "12123232",
			},
			want:    entities.SPAN{},
			wantErr: false,
		},
		{
			name: "Fail test case",
			args: args{
				ctx:            context.Background(),
				documentNumber: "12123232",
			},
			want:    entities.SPAN{},
			wantErr: true,
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
			ucMock.spanRepo.EXPECT().GetBySpanDocumentNumber(tt.args.ctx, tt.args.documentNumber).Return(tt.want, err)
			got, err := uc.GetByDocumentNumber(tt.args.ctx, tt.args.documentNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByDocumentNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByDocumentNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_UpdatePatchBankRiau(t *testing.T) {
	type args struct {
		ctx  context.Context
		span entities.SPAN
		user entities.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:  utils.SetUserToContext(context.Background(), *defaultUser()),
				span: entities.SPAN{DocumentNumber: "123123"},
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
			ucMock.spanRepo.EXPECT().GetBySpanDocumentNumber(tt.args.ctx, tt.args.span.DocumentNumber).
				Return(tt.args.span, err)
			newSpan := PatchBankRiau(tt.args.span)
			auditEnt := entities.Audit{}
			ucMock.spanRepo.EXPECT().MakeAuditUpdateWithOldData(tt.args.ctx, tt.args.span, newSpan).Return(auditEnt, err)
			tx := mocksTransactor.NewTransactor(t)
			ucMock.spanRepo.EXPECT().Begin().Return(tx)
			ucMock.auditRepo.EXPECT().Begin().Return(tx)
			ucMock.spanRepo.EXPECT().New(tx).Return(uc.spanRepo)
			ucMock.auditRepo.EXPECT().New(tx).Return(uc.auditRepo)
			ucMock.spanRepo.EXPECT().Update(tt.args.ctx, newSpan).Return(err)
			ucMock.auditRepo.EXPECT().Create(auditEnt).Return(err)
			tx.EXPECT().Commit().Return(&gorm.DB{}).Times(2)
			if err := uc.UpdatePatchBankRiau(tt.args.ctx, tt.args.span); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePatchBankRiau() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
