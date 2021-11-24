package recursobd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	recursosmodels "github.com/ascendere/micro-convocatorias/models/recursos_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidoRecurso(id primitive.ObjectID, tk string) (recursosmodels.ConsultoRecurso, error) {
	var recurso recursosmodels.ConsultoRecurso

	client := &http.Client{}

	objID := id.String()

	req, err := http.NewRequest("GET", "http://34.121.243.107/buscarRecurso?recurso="+objID, nil)

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

	log.Println(recurso)

	return recurso, err
}
