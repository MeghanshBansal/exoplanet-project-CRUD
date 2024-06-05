package domain

import (
	"NTTData/Models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDomain_CalculateFuel(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDbService := newMockDbServiceCalculateFuel(ctrl)
	mockDomain := InitializeDomain(mockDbService)

	type args struct {
		id       string
		crewSize int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		Resp    float64
	}{
		{
			name:    "tc1",
			args:    args{id: "123", crewSize: 5},
			wantErr: false,
			Resp:    14.234305555555558,
		},
		{

			name:    "tc2",
			args:    args{id: "456", crewSize: 5},
			wantErr: false,
			Resp:    2049.7400000000007,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if value, err := mockDomain.CalculateFuel(tt.args.id, int64(tt.args.crewSize)); (err != nil) != tt.wantErr {
				t.Errorf("AddExoplanet() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Equal(t, tt.Resp, value)
			}
		})
	}
}

func newMockDbServiceCalculateFuel(ctrl *gomock.Controller) *MockDBServices {
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

	mockDb.Expect().GetExoplanetByIdDb("456").Return(Models.Exoplanet{
		Id:                "456",
		Name:              "Proxima Centauri b",
		Description:       "An exoplanet orbiting in the habitable zone of the closest star to the Sun, Proxima Centauri.",
		DistanceFromEarth: 70,
		Radius:            1.1,
		Mass:              0,
		TypeOfExoplanet:   "GasGiant",
	}, nil).AnyTimes()
	return mockDb
}
