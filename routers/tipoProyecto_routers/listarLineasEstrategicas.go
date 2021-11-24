package tipoProyectorouters

import (
	"encoding/json"
	"net/http"

	tipoProyectobd "github.com/ascendere/micro-convocatorias/bd/tipoProyecto_bd"
)

func ListarTiposProyectos(w http.ResponseWriter, r *http.Request) {

	result, status := tipoProyectobd.ListoTipoProyecto()
	if !status {
		http.Error(w, "Error al leer los tipos de proyectos", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}