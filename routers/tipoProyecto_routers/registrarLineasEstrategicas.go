package tipoProyectorouters

import (
	"encoding/json"
	"net/http"

	tipoProyectobd "github.com/ascendere/micro-convocatorias/bd/tipoProyecto_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarTipoProyecto(w http.ResponseWriter, r *http.Request) {
	var tipo convocatoriamodels.TipoProyecto

	err := json.NewDecoder(r.Body).Decode(&tipo)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := tipoProyectobd.RegistroTipoProyecto(tipo)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nueva tipo proyecto ", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nueva tipo proyecto ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}