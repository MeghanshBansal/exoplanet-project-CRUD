package domain

import (
	"Exoplanet/Models"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockDBServices struct {
	ctrl     *gomock.Controller
	recorder *MockDBServiceMockRecorder
}

func (p *MockDBServices) AddExoplanetToDb(exoplanet Models.Exoplanet) error {
	p.ctrl.T.Helper()
	exoplanet.Id = "0"
	ret := p.ctrl.Call(p, "AddExoplanetToDb", exoplanet)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDBServiceMockRecorder) AddExoplanetToDb(exoplanet Models.Exoplanet) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddExoplanetToDb", reflect.TypeOf((*MockDBServices)(nil).AddExoplanetToDb), exoplanet)
}

func (p *MockDBServices) GetAllExoplanetsDb() ([]Models.Exoplanet, error) {
	p.ctrl.T.Helper()
	ret := p.ctrl.Call(p, "GetAllExoplanetsDb")
	ret0, _ := ret[0].([]Models.Exoplanet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDBServiceMockRecorder) GetAllExoplanetsDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllExoplanetsDb", reflect.TypeOf((*MockDBServices)(nil).GetAllExoplanetsDb))
}

func (p *MockDBServices) GetExoplanetByIdDb(id string) (Models.Exoplanet, error) {
	p.ctrl.T.Helper()
	ret := p.ctrl.Call(p, "GetExoplanetByIdDb", id)
	ret0, _ := ret[0].(Models.Exoplanet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDBServiceMockRecorder) GetExoplanetByIdDb(id string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExoplanetByIdDb", reflect.TypeOf((*MockDBServices)(nil).GetExoplanetByIdDb), id)
}

func (p *MockDBServices) UpdateExoplanetByIdDb(id string, exoplanet Models.Exoplanet) error {
	p.ctrl.T.Helper()
	ret := p.ctrl.Call(p, "UpdateExoplanetByIdDb", id, exoplanet)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDBServiceMockRecorder) UpdateExoplanetByIdDb(id string, exoplanet Models.Exoplanet) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExoplanetByIdDb", reflect.TypeOf((*MockDBServices)(nil).UpdateExoplanetByIdDb), id, exoplanet)
}

func (p *MockDBServices) DeleteExoplanetByIdDb(id string) error {
	p.ctrl.T.Helper()
	ret := p.ctrl.Call(p, "DeleteExoplanetByIdDb", id)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDBServiceMockRecorder) DeleteExoplanetByIdDb(id string) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExoplanetByIdDb", reflect.TypeOf((*MockDBServices)(nil).DeleteExoplanetByIdDb), id)
}

type MockDBServiceMockRecorder struct {
	mock *MockDBServices
}

func NewMockDBServices(ctrl *gomock.Controller) *MockDBServices {
	mock := &MockDBServices{
		ctrl: ctrl,
	}
	mock.recorder = &MockDBServiceMockRecorder{mock}

	return mock
}

func (p *MockDBServices) Expect() *MockDBServiceMockRecorder {
	return p.recorder
}
