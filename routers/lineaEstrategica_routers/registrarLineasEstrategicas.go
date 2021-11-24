package lineaEstrategicarouters

import (
	"encoding/json"
	"net/http"

	lineaEstrategicabd "github.com/ascendere/micro-convocatorias/bd/lineaEstrategica_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarLineaEstrategica(w http.ResponseWriter, r *http.Request) {
	var linea convocatoriamodels.LineaEstrategica

	err := json.NewDecoder(r.Body).Decode(&linea)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := lineaEstrategicabd.RegistroLineaEstrategica(linea)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nueva linea estrategica", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nueva linea estrategica", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}