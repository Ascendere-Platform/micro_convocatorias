package rubricarouters

import (
	"encoding/json"
	"net/http"

	rubricabd "github.com/ascendere/micro-convocatorias/bd/rubrica_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarRubrica(w http.ResponseWriter, r *http.Request) {
	var rubrica convocatoriamodels.Rubrica

	err := json.NewDecoder(r.Body).Decode(&rubrica)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := rubricabd.RegistroRubrica(rubrica)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nueva rubrica", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nueva rubrica", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}