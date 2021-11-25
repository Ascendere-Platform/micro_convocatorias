package convocatoriarouters

import (
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
)

func EliminarConvocatoria(w http.ResponseWriter, r *http.Request) {

	convocatoria := r.URL.Query().Get("id")

	if len(convocatoria) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := convocatoriabd.EliminoConvocatoria(convocatoria)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}