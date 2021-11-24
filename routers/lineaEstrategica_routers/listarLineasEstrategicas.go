package lineaEstrategicarouters

import (
	"encoding/json"
	"net/http"

	lineaEstrategicabd "github.com/ascendere/micro-convocatorias/bd/lineaEstrategica_bd"
)

func ListarLineasEstrategicas(w http.ResponseWriter, r *http.Request) {

	result, status := lineaEstrategicabd.ListoLineasEstrategicas()
	if !status {
		http.Error(w, "Error al leer las lineas estrategicas", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}