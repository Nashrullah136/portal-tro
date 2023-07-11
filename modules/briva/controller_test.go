package briva

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
	return controller{brivaUseCase: mock.useCase}
}

func Test_controller_GetByBrivaNo(t *testing.T) {
	briva := entities.Briva{}
	type args struct {
		ctx     context.Context
		brivano string
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
				brivano: "121",
			},
			want:    dto.Success("Brivano has been found", briva),
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
			cMock.useCase.EXPECT().GetByBrivaNo(tt.args.ctx, tt.args.brivano).Return(briva, err)
			got, err := c.GetByBrivaNo(tt.args.ctx, tt.args.brivano)
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

func Test_controller_Update(t *testing.T) {
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
			want:    dto.Success("Success update briva", nil),
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
			briva := mapUpdateRequestToBriva(tt.args.request)
			cMock.useCase.EXPECT().ValidateBriva(briva, mock.Anything).Return(err, err)
			cMock.useCase.EXPECT().Update(tt.args.ctx, briva).Return(err)
			got, err := c.Update(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
