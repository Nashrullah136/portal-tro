package audit

import (
	"context"
	"errors"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/modules/audit/mocks"
	csvutils "nashrul-be/crm/utils/csv"
	"nashrul-be/crm/utils/filesystem"
	"reflect"
	"testing"
)

type controllerMock struct {
	useCase *mocks.UseCaseInterface
}

func defaultControllerMock(t *testing.T) controllerMock {
	return controllerMock{useCase: mocks.NewUseCaseInterface(t)}
}

func defaultController(mock controllerMock) ControllerInterface {
	return controller{auditUseCase: mock.useCase}
}

func Test_controller_CreateAudit(t *testing.T) {
	type args struct {
		ctx    context.Context
		action string
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
				ctx:    context.Background(),
				action: "login",
			},
			want:    dto.Success("Success create audit", nil),
			wantErr: false,
		},
		{
			name: "Fail case",
			args: args{
				ctx:    context.Background(),
				action: "login",
			},
			want:    dto.ErrorInternalServerError(),
			wantErr: true,
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
			cMock.useCase.EXPECT().CreateAudit(tt.args.ctx, tt.args.action).Return(err)
			got, err := c.CreateAudit(tt.args.ctx, tt.args.action)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAudit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAudit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_ExportCSV(t *testing.T) {
	defaultFolder, _ := csvutils.NewCSV(filesystem.NewFolder("asdasd"))
	type args struct {
		ctx     context.Context
		request ExportRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *csvutils.FileCsv
		wantErr bool
	}{
		{
			name: "Normal case",
			args: args{
				ctx:     context.Background(),
				request: ExportRequest{},
			},
			want:    defaultFolder,
			wantErr: false,
		},
		{
			name: "Normal case",
			args: args{
				ctx:     context.Background(),
				request: ExportRequest{},
			},
			want:    nil,
			wantErr: true,
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
			auditQuery := mapExportRequestToAuditQuery(tt.args.request)
			cMock.useCase.EXPECT().ExportCsv(tt.args.ctx, auditQuery).Return(tt.want, err)
			got, err := c.ExportCSV(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExportCSV() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_GetAll(t *testing.T) {
	totalRow := 10
	var result []entities.Audit
	type args struct {
		ctx     context.Context
		request GetAllRequest
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
				ctx:     context.Background(),
				request: GetAllRequest{},
			},
			want:    dto.SuccessPagination("Success retrieve audit", 0, 0, totalRow, result),
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
			auditQuery := mapGetAllRequestToAuditQuery(tt.args.request)
			offset := (tt.args.request.Page - 1) * tt.args.request.PerPage
			cMock.useCase.EXPECT().GetAll(tt.args.ctx, auditQuery, tt.args.request.PerPage, offset).
				Return(result, err)
			cMock.useCase.EXPECT().CountAll(tt.args.ctx, auditQuery).Return(totalRow, err)
			got, err := c.GetAll(tt.args.ctx, tt.args.request)
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
