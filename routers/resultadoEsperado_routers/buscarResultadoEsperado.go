package resultadoEsperadorouters

import (
	"encoding/json"
	"net/http"

	resultadoEsperadobd "github.com/ascendere/micro-convocatorias/bd/resultadoEsperado_bd"
)

func BuscarResultadoEsperado(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := resultadoEsperadobd.BuscoResultadoEsperado(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un resultado esperado ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}