package rubricarouters

import (
	"encoding/json"
	"net/http"

	rubricabd "github.com/ascendere/micro-convocatorias/bd/rubrica_bd"
)

func BuscarRubrica(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := rubricabd.BuscoRubrica(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una rubrica ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}