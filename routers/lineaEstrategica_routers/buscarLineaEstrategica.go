package lineaEstrategicarouters

import (
	"encoding/json"
	"net/http"

	lineaEstrategicabd "github.com/ascendere/micro-convocatorias/bd/lineaEstrategica_bd"
)

func BuscarLineaEstrategica(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := lineaEstrategicabd.BuscoLineaEstrategica(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una linea estrategica ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}