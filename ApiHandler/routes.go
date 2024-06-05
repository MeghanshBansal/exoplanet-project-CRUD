package ApiHandler

import (
	"NTTData/Domain"
	"NTTData/Models"
	"NTTData/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Webservice struct {
	Domain Domain.Service
	Server Server
}

func NewWebservices(servicePort string, d Domain.Service) *Webservice {
	server := NewServer(servicePort)
	return &Webservice{
		Server: server,
		Domain: d,
	}
}

func (s *Webservice) Start() error {
	s.addRoutes()
	err := s.Server.Start()
	return err
}

func (s *Webservice) addRoutes() {
	http.HandleFunc("/ping", s.ping)
	http.HandleFunc("/add-exoplanet", s.AddExoplanets)
	http.HandleFunc("/list-exoplanet", s.ListExoplanet)
	http.HandleFunc("/get-exoplanet", s.GetExoplanetById)
	http.HandleFunc("/update-exoplanet", s.UpdateExoplanet)
	http.HandleFunc("/delete-exoplanet", s.DeleteExoplanet)
	http.HandleFunc("/get-fuel-estimation", s.GetFuelEstimation)
}

func (s *Webservice) ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		return
	}
	return
}

func (s *Webservice) AddExoplanets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var exoplanet Models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, "failed to decode request")
		return
	}

	err = s.Domain.AddExoplanet(exoplanet)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to add exoplanet")
		return
	}
	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanet Added Successfully",
		},
	})
	return
}

func (s *Webservice) ListExoplanet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	resp, err := s.Domain.GetAllExoplanet()
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to get all exoplanet")
		return
	}

	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanets Fetched Successfully",
		},
		Data: resp,
	})
	return
}

func (s *Webservice) GetExoplanetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ReturnResponse(w, http.StatusBadRequest, "id is empty")
		return
	}
	resp, err := s.Domain.GetExoplanetById(id)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to get exoplanet")
		return
	}
	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanet Fetched Successfully",
		},
		Data: resp,
	})
	return
}

func (s *Webservice) UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ReturnResponse(w, http.StatusBadRequest, "id is empty")
		return
	}

	var updatePlanet Models.UpdateMapExoplanet
	err := json.NewDecoder(r.Body).Decode(&updatePlanet)
	if err != nil {
		utils.ReturnResponse(w, http.StatusBadRequest, "failed to decode request")
		return
	}

	err = s.Domain.UpdateExoplanetById(id, updatePlanet)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to update exoplanet")
		return
	}
	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanet Updated Successfully",
		},
	})
	return

}

func (s *Webservice) DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		utils.ReturnResponse(w, http.StatusBadRequest, "id is empty")
		return
	}
	err := s.Domain.DeleteExoplanetById(id)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to delete exoplanet")
		return
	}
	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanet Deleted Successfully",
		},
	})
	return
}

func (s *Webservice) GetFuelEstimation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id := r.URL.Query().Get("id")
	crewSize := r.URL.Query().Get("crewSize")
	if id == "" || crewSize == "" || crewSize == "0" {
		utils.ReturnResponse(w, http.StatusBadRequest, "given values are incorrect")
		return
	}
	crewSizeInt, err := strconv.ParseInt(crewSize, 10, 64)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to parse request")
		return
	}

	resp, err := s.Domain.CalculateFuel(id, crewSizeInt)
	if err != nil {
		utils.ReturnResponse(w, http.StatusInternalServerError, "failed to get fuel estimation")
		return
	}
	utils.ReturnResponse(w, http.StatusOK, Models.Response{
		Meta: Models.Meta{
			StatusCode: http.StatusOK,
			Message:    "Exoplanet Deleted Successfully",
		},
		Data: fmt.Sprintf("%g units", resp),
	})
}
