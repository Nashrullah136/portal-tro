package export_csv

import (
	"context"
	"errors"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/repositories"
	"nashrul-be/crm/repositories/mocks"
	"nashrul-be/crm/utils/filesystem"
	mocksFs "nashrul-be/crm/utils/filesystem/mocks"
	"reflect"
	"testing"
)

type useCaseMock struct {
	exportCsvRepo *mocks.ExportCsvRepositoryInterface
	auditRepo     *mocks.AuditRepositoryInterface
	folder        *mocksFs.Folder
}

func defaultUseCaseMock(t *testing.T) useCaseMock {
	return useCaseMock{
		exportCsvRepo: mocks.NewExportCsvRepositoryInterface(t),
		auditRepo:     mocks.NewAuditRepositoryInterface(t),
		folder:        mocksFs.NewFolder(t),
	}
}

func defaultUseCase(mock useCaseMock) useCase {
	return useCase{
		exportCsvRepo: mock.exportCsvRepo,
		auditRepo:     mock.auditRepo,
		folder:        mock.folder,
	}
}

func Test_useCase_CountAll(t *testing.T) {
	type args struct {
		ctx   context.Context
		query repositories.ExportCsvQuery
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx:   context.Background(),
				query: repositories.ExportCsvQuery{},
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultUseCaseMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.exportCsvRepo.EXPECT().CountAll(tt.args.ctx, tt.args.query).Return(tt.want, err)
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

func Test_useCase_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultUseCaseMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			exportCsv := entities.ExportCsv{}
			ucMock.exportCsvRepo.EXPECT().GetById(tt.args.id).Return(exportCsv, err)
			file := mocksFs.NewFile(t)
			ucMock.folder.EXPECT().GetFile(exportCsv.Filename).Return(file, err)
			file.EXPECT().Remove().Return(err)
			ucMock.exportCsvRepo.EXPECT().Delete(tt.args.ctx, tt.args.id).Return(err)
			if err := uc.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCase_DownloadCsv(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		args    args
		want    filesystem.File
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultUseCaseMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			exportCsv := entities.ExportCsv{}
			ucMock.exportCsvRepo.EXPECT().GetById(tt.args.id).Return(exportCsv, err)
			tt.want = filesystem.NewFile(exportCsv.Filename, uc.folder)
			ucMock.folder.EXPECT().GetFile(exportCsv.Filename).Return(tt.want, err)
			got, err := uc.DownloadCsv(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadCsv() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCase_GetAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		query  repositories.ExportCsvQuery
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []entities.ExportCsv
		wantErr bool
	}{
		{
			name: "Normal Scenario",
			args: args{
				ctx:    context.Background(),
				query:  repositories.ExportCsvQuery{},
				limit:  10,
				offset: 0,
			},
			want:    []entities.ExportCsv{{}, {}, {}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ucMock := defaultUseCaseMock(t)
			uc := defaultUseCase(ucMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			ucMock.exportCsvRepo.EXPECT().GetAll(tt.args.ctx, tt.args.query, tt.args.limit, tt.args.offset).Return(tt.want, err)
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
