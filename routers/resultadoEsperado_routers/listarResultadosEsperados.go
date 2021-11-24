package resultadoEsperadorouters

import (
	"encoding/json"
	"net/http"

	resultadoEsperadobd "github.com/ascendere/micro-convocatorias/bd/resultadoEsperado_bd"
)

func ListarResultadosEsperados(w http.ResponseWriter, r *http.Request) {

	result, status := resultadoEsperadobd.ListoResultadosEsperados()
	if !status {
		http.Error(w, "Error al leer los resultados esperados", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}