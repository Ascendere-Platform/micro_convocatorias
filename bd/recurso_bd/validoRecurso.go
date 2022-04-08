package recursobd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	recursosmodels "github.com/ascendere/micro-convocatorias/models/recursos_models"
)

func ValidoRecurso(id string, tk string) (recursosmodels.ConsultoRecurso, error) {
	var recurso recursosmodels.ConsultoRecurso

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://35.239.170.70/buscarRecurso?recurso="+id, nil)

	if err != nil {
		return recurso, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tk)

	resp, error := client.Do(req)

	if error != nil {
		return recurso, error
	}

	defer resp.Body.Close()

	bodyBytes, errorBytes := ioutil.ReadAll(resp.Body)

	if errorBytes != nil {
		return recurso, errorBytes
	}

	json.Unmarshal(bodyBytes, &recurso)



	return recurso, err
}
