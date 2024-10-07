package Models

type Exoplanet struct {
	Id                string  `json:"id" gorm:"primary key"`
	Name              string  `json:"name"`
	Description       string  `json:"description,omitempty"`
	DistanceFromEarth int     `json:"distanceFromEarth"`
	Radius            float64 `json:"radius"`
	Mass              float64 `json:"mass,omitempty"`
	TypeOfExoplanet   string  `json:"typeOfExoplanet"`
}

type UpdateMapExoplanet struct {
	InputFields []InputFields `json:"inputFields"`
}

type InputFields struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

type Meta struct {
	StatusCode int
	Message    string
}
