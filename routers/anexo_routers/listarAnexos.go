package anexorouters

import (
	"encoding/json"
	"net/http"

	anexobd "github.com/ascendere/micro-convocatorias/bd/anexo_bd"
)

func ListarAnexos(w http.ResponseWriter, r *http.Request) {

	result, status := anexobd.ListoAnexos()
	if !status {
		http.Error(w, "Error al leer los anexos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}