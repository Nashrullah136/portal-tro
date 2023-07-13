package span

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"reflect"
	"testing"
)

type controllerMock struct {
	useCase *MockUseCaseInterface
}

func defaultControllerMock(t *testing.T) controllerMock {
	return controllerMock{useCase: NewMockUseCaseInterface(t)}
}

func defaultController(mock controllerMock) ControllerInterface {
	return controller{spanUseCase: mock.useCase}
}

func Test_controller_GetByDocumentNumber(t *testing.T) {
	span := entities.SPAN{StatusCode: "0002"}
	type args struct {
		ctx            context.Context
		documentNumber string
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
				ctx:            context.Background(),
				documentNumber: "12122",
			},
			want:    dto.Success("Success retrieve span", mapSpanToPresentation(span)),
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
			cMock.useCase.EXPECT().GetByDocumentNumberPatchBankRiau(tt.args.ctx, tt.args.documentNumber).
				Return(span, err)
			got, err := c.GetByDocumentNumber(tt.args.ctx, tt.args.documentNumber)
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

func Test_controller_UpdateBankRiau(t *testing.T) {
	type args struct {
		ctx     context.Context
		request UpdateRequest
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
				request: UpdateRequest{},
			},
			want:    dto.Success("Success update SPAN", nil),
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
			span := mapUpdateRequestToSpan(tt.args.request)
			cMock.useCase.EXPECT().ValidateSpan(span, mock.Anything).Return(err, err)
			cMock.useCase.EXPECT().UpdatePatchBankRiau(tt.args.ctx, span).Return(err)
			got, err := c.UpdateBankRiau(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateBankRiau() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateBankRiau() got = %v, want %v", got, tt.want)
			}
		})
	}
}
