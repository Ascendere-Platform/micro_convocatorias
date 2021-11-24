package resultadoEsperadorouters

import (
	"net/http"

	resultadoEsperadobd "github.com/ascendere/micro-convocatorias/bd/resultadoEsperado_bd"
)

func EliminarResultadoEsperado(w http.ResponseWriter, r *http.Request) {

	resultado := r.URL.Query().Get("id")

	if len(resultado) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := resultadoEsperadobd.EliminoResultadoEsperado(resultado)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}