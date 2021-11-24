package tipoProyectorouters

import (
	"encoding/json"
	"net/http"

	tipoProyectobd "github.com/ascendere/micro-convocatorias/bd/tipoProyecto_bd"
)

func BuscarTipoProyecto(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := tipoProyectobd.BuscoTipoProyecto(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un tipo de proyecto ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}