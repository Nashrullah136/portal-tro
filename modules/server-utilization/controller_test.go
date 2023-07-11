package server_utilization

import (
	"errors"
	"nashrul-be/crm/dto"
	"nashrul-be/crm/entities"
	"nashrul-be/crm/utils/zabbix/mocks"

	"nashrul-be/crm/utils/zabbix"
	"reflect"
	"testing"
)

type controllerMock struct {
	cache     *mocks.Cache
	zabbixApi *mocks.API
}

func defaultControllerMock(t *testing.T) controllerMock {
	return controllerMock{
		cache:     mocks.NewCache(t),
		zabbixApi: mocks.NewAPI(t),
	}
}

func defaultController(mock controllerMock) ControllerInterface {
	return controller{
		cache:     mock.cache,
		zabbixApi: mock.zabbixApi,
	}
}

func Test_controller_GetLastData(t *testing.T) {
	var (
		safe      []entities.ServerUtilization
		threshold []entities.ServerUtilization
	)
	result := map[string][]entities.ServerUtilization{
		"threshold": threshold,
		"safe":      safe,
	}
	tests := []struct {
		name    string
		want    dto.BaseResponse
		wantErr bool
	}{
		{
			name:    "Normal case",
			want:    dto.Success("Success get latest data", result),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cMock := defaultControllerMock(t)
			uc := defaultController(cMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			cMock.cache.EXPECT().GetLastValue().Return([]entities.ServerUtilization{}, err)
			got, err := uc.GetLastData()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLastData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_controller_RefreshHostList(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Normal case",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cMock := defaultControllerMock(t)
			uc := defaultController(cMock)
			err := errors.New("mock")
			if !tt.wantErr {
				err = nil
			}
			cMock.zabbixApi.EXPECT().GetAllHost().Return([]zabbix.Host{}, err)
			cMock.zabbixApi.EXPECT().GetItemFromHosts([]string{}).Return([]zabbix.Item{}, err)
			cMock.cache.EXPECT().SetTemplate([]entities.ServerUtilization{}).Return(err)
			cMock.cache.EXPECT().SetItemIds([]string{}).Return(err)
			if err := uc.RefreshHostList(); (err != nil) != tt.wantErr {
				t.Errorf("RefreshHostList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
