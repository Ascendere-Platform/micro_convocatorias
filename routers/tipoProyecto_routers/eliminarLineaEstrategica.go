package tipoProyectorouters

import (
	"net/http"

	tipoProyectobd "github.com/ascendere/micro-convocatorias/bd/tipoProyecto_bd"
)

func EliminarTipoProyecto(w http.ResponseWriter, r *http.Request) {

	tipo := r.URL.Query().Get("id")

	if len(tipo) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := tipoProyectobd.EliminoTipoProyecto(tipo)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}