package domain

import (
	"NTTData/Models"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDomain_AddExoplanet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceAddExoplanet(ctrl)
	mockDomain := InitializeDomain(mockDbService)

	type args struct {
		exoplanet Models.Exoplanet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "tc1",
			args: args{exoplanet: Models.Exoplanet{
				Name:              "Proxima Centauri b",
				Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
				DistanceFromEarth: 70,
				Radius:            1.1,
				Mass:              6.0,
				TypeOfExoplanet:   "Terrestrial",
			}},
			wantErr: false,
		},
		{
			name: "tc2",
			args: args{exoplanet: Models.Exoplanet{
				Name:              "Proxima Centauri c",
				Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
				DistanceFromEarth: 9,
				Radius:            1.1,
				Mass:              6.0,
				TypeOfExoplanet:   "Terrestrial",
			}},
			wantErr: true,
		},
		{
			name: "tc3",
			args: args{exoplanet: Models.Exoplanet{
				Name:              "Proxima Centauri d",
				Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
				DistanceFromEarth: 9,
				Radius:            1.1,
				Mass:              0,
				TypeOfExoplanet:   "GasGiant",
			}},
			wantErr: true,
		},
		{
			name: "tc4",
			args: args{exoplanet: Models.Exoplanet{
				Name:              "Proxima Centauri e",
				Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
				DistanceFromEarth: 80,
				Radius:            1.1,
				Mass:              0,
				TypeOfExoplanet:   "abc",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mockDomain.AddExoplanet(tt.args.exoplanet); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func newMockDbServiceAddExoplanet(ctrl *gomock.Controller) *MockDBServices {
	mockDb := NewMockDBServices(ctrl)
	mockDb.Expect().AddExoplanetToDb(Models.Exoplanet{
		Id:                "0",
		Name:              "Proxima Centauri b",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 70,
		Radius:            1.1,
		Mass:              6.0,
		TypeOfExoplanet:   "Terrestrial",
	}).Return(nil).AnyTimes()

	mockDb.Expect().AddExoplanetToDb(Models.Exoplanet{
		Id:                "0",
		Name:              "Proxima Centauri c",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 9,
		Radius:            1.1,
		Mass:              6.0,
		TypeOfExoplanet:   "Terrestrial",
	}).Return(errors.New("")).AnyTimes()

	mockDb.Expect().AddExoplanetToDb(Models.Exoplanet{
		Name:              "Proxima Centauri d",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 9,
		Radius:            1.1,
		Mass:              0,
		TypeOfExoplanet:   "GasGiant",
	}).Return(nil).AnyTimes()

	mockDb.Expect().AddExoplanetToDb(Models.Exoplanet{
		Name:              "Proxima Centauri e",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 9,
		Radius:            1.1,
		Mass:              0,
		TypeOfExoplanet:   "abc",
	}).Return(errors.New("")).AnyTimes()
	return mockDb
}

func TestDomain_DeleteExoplanetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceDeleteExoplanetById(ctrl)
	mockDomain := InitializeDomain(mockDbService)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "tc1",
			args:    args{id: "123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mockDomain.DeleteExoplanetById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func newMockDbServiceDeleteExoplanetById(ctrl *gomock.Controller) *MockDBServices {
	mockDb := NewMockDBServices(ctrl)
	mockDb.Expect().DeleteExoplanetByIdDb("123").Return(nil).AnyTimes()
	return mockDb
}

func TestDomain_GetAllExoplanet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceGetAllExoplanet(ctrl)
	mockDomain := InitializeDomain(mockDbService)

	tests := []struct {
		name string

		wantErr bool
		Resp    []Models.Exoplanet
	}{
		{
			name:    "tc1",
			wantErr: false,
			Resp: []Models.Exoplanet{
				{Id: "0", Name: "Proxima Centauri b", Description: "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.", DistanceFromEarth: 70, Radius: 1.1, Mass: 6, TypeOfExoplanet: "Terrestrial"},
				{Id: "1", Name: "Proxima Centauri 1", Description: "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.", DistanceFromEarth: 60, Radius: 1.1, Mass: 6, TypeOfExoplanet: "Terrestrial"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if resp, err := mockDomain.GetAllExoplanet(); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.Resp, resp)
			}
		})
	}
}
func newMockDbServiceGetAllExoplanet(ctrl *gomock.Controller) *MockDBServices {
	mockDb := NewMockDBServices(ctrl)
	mockDb.Expect().GetAllExoplanetsDb().Return([]Models.Exoplanet{
		{
			Id:                "0",
			Name:              "Proxima Centauri b",
			Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
			DistanceFromEarth: 70,
			Radius:            1.1,
			Mass:              6.0,
			TypeOfExoplanet:   "Terrestrial",
		},
		{
			Id:                "1",
			Name:              "Proxima Centauri 1",
			Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
			DistanceFromEarth: 60,
			Radius:            1.1,
			Mass:              6.0,
			TypeOfExoplanet:   "Terrestrial",
		},
	}, nil).AnyTimes()
	return mockDb
}

func TestDomain_GetExoplanetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceGetExoplanetById(ctrl)
	mockDomain := InitializeDomain(mockDbService)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		Resp    Models.Exoplanet
	}{
		{
			name: "tc1",
			args: args{
				id: "123",
			},
			wantErr: false,
			Resp:    Models.Exoplanet{Id: "123", Name: "Proxima Centauri b", Description: "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.", DistanceFromEarth: 70, Radius: 1.1, Mass: 6, TypeOfExoplanet: "Terrestrial"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if resp, err := mockDomain.GetExoplanetById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.Resp, resp)
			}
		})
	}
}

func newMockDbServiceGetExoplanetById(ctrl *gomock.Controller) *MockDBServices {
	mockDb := NewMockDBServices(ctrl)
	mockDb.Expect().GetExoplanetByIdDb("123").Return(Models.Exoplanet{
		Id:                "123",
		Name:              "Proxima Centauri b",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 70,
		Radius:            1.1,
		Mass:              6.0,
		TypeOfExoplanet:   "Terrestrial",
	}, nil).AnyTimes()
	return mockDb
}

func TestDomain_UpdateExoplanetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceUpdateExoplanetById(ctrl)
	mockDomain := InitializeDomain(mockDbService)

	type args struct {
		id        string
		updateMap Models.UpdateMapExoplanet
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "tc1",
			args: args{id: "123", updateMap: Models.UpdateMapExoplanet{
				InputFields: []Models.InputFields{
					{Key: "name", Value: "New Planet"},
				},
			}},
			wantErr: false,
		},
		{
			name: "tc2",
			args: args{id: "456", updateMap: Models.UpdateMapExoplanet{
				InputFields: []Models.InputFields{
					{Key: "name", Value: "New Planet"},
				},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mockDomain.UpdateExoplanetById(tt.args.id, tt.args.updateMap); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func newMockDbServiceUpdateExoplanetById(ctrl *gomock.Controller) *MockDBServices {
	mockDb := NewMockDBServices(ctrl)
	mockDb.Expect().GetExoplanetByIdDb("123").Return(Models.Exoplanet{
		Id:                "123",
		Name:              "Proxima Centauri b",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 70,
		Radius:            1.1,
		Mass:              6.0,
		TypeOfExoplanet:   "Terrestrial",
	}, nil).AnyTimes()

	mockDb.Expect().UpdateExoplanetByIdDb("123", Models.Exoplanet{
		Id:                "123",
		Name:              "New Planet",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 70,
		Radius:            1.1,
		Mass:              6.0,
		TypeOfExoplanet:   "Terrestrial",
	}).Return(nil).AnyTimes()

	mockDb.Expect().GetExoplanetByIdDb("456").Return(Models.Exoplanet{}, errors.New("")).AnyTimes()
	return mockDb
}
