package lineaEstrategicarouters

import (
	"net/http"

	lineaEstrategicabd "github.com/ascendere/micro-convocatorias/bd/lineaEstrategica_bd"
)

func EliminarLineaEstrategica(w http.ResponseWriter, r *http.Request) {

	anexo := r.URL.Query().Get("id")

	if len(anexo) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := lineaEstrategicabd.EliminoLineaEstrategica(anexo)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}