package rubricarouters

import (
	"encoding/json"
	"net/http"

	rubricabd "github.com/ascendere/micro-convocatorias/bd/rubrica_bd"
)

func ListarRubricas(w http.ResponseWriter, r *http.Request) {

	result, status := rubricabd.ListoRubricas()
	if !status {
		http.Error(w, "Error al leer las lineas estrategicas", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}