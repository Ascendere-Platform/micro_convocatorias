package convocatoriarouters

import (
	"encoding/json"
	"net/http"
	"time"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarConvocatoria (w http.ResponseWriter, r *http.Request) {
	var convocatoria convocatoriamodels.Convocatoria

	err := json.NewDecoder(r.Body).Decode(&convocatoria)

	convocatoria.FechaCreacion = time.Now()

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := convocatoriabd.RegistroConvocatoria(convocatoria)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nueva convocatoria ", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nueva convocatoria ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}