package audit

import (
	"context"
	"errors"
	"github.com/adjust/rmq/v5"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/repositories/mocks"
	"reflect"
	"testing"
)

type useCaseMock struct {
	auditRepo     *mocks.AuditRepositoryInterface
	exportCsvRepo *mocks.ExportCsvRepositoryInterface
	queue         rmq.Queue
}

func defaultMock(t *testing.T) useCaseMock {
	return useCaseMock{
		auditRepo:     mocks.NewAuditRepositoryInterface(t),
		exportCsvRepo: mocks.NewExportCsvRepositoryInterface(t),
		queue:         rmq.NewTestQueue("mock-queue"),
	}
}

func defaultUseCase(ucMock useCaseMock) useCase {
	return useCase{
		auditRepo:     ucMock.auditRepo,
		exportCsvRepo: ucMock.exportCsvRepo,
		queue:         ucMock.queue,
	}
}

func Test_useCase_CountAll(t *testing.T) {
	type args struct {
		ctx   context.Context
		query repositories.AuditQuery
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Normal scenario",
			args: args{
				ctx:   context.Background(),
				query: repositories.AuditQuery{},
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
			ucMock.auditRepo.EXPECT().CountGetAll(tt.args.ctx, tt.args.query).Return(tt.want, err)
			got, err := uc.CountAll(tt.args.ctx, tt.args.query)
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

func Test_useCase_CreateAudit(t *testing.T) {
	type args struct {
		ctx    context.Context
		action string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx:    context.Background(),
				action: "Open menu BRIVA",
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
			ucMock.auditRepo.EXPECT().CreateAudit(tt.args.ctx, tt.args.action).Return(err)
			if err := uc.CreateAudit(tt.args.ctx, tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("CreateAudit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_ExportCSV(t *testing.T) {
	type args struct {
		ctx   context.Context
		query repositories.AuditQuery
	}

	account := entities.User{
		Username: "admin",
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal scenario",
			args: args{
				ctx:   context.WithValue(context.Background(), "user", account),
				query: repositories.AuditQuery{},
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
			csvReq := entities.InitExportCsv(account.Username)
			csvReqAfter := csvReq
			csvReqAfter.ID = 1
			ucMock.exportCsvRepo.EXPECT().Create(csvReq).Return(csvReqAfter, err)
			if err := uc.ExportCsvAsync(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("ExportCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		query  repositories.AuditQuery
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []entities.Audit
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx:    context.Background(),
				query:  repositories.AuditQuery{},
				limit:  10,
				offset: 0,
			},
			want:    []entities.Audit{{}, {}},
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
			ucMock.auditRepo.EXPECT().GetAll(tt.args.ctx, tt.args.query, tt.args.limit, tt.args.offset).
				Return(tt.want, err)
			got, err := uc.GetAll(tt.args.ctx, tt.args.query, tt.args.limit, tt.args.offset)
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
