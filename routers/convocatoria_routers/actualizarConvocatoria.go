package convocatoriarouters

import (
	"encoding/json"
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func ActualizarConvocatoria(w http.ResponseWriter, r *http.Request) {

	var t convocatoriamodels.Convocatoria

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = convocatoriabd.ActualizoConvocatoria(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar la convocatoria "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar la convocatoria ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}