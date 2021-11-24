package routers

import (
	"encoding/json"
	"net/http"

	recursobd "github.com/ascendere/micro-convocatorias/bd/recurso_bd"
)

func TestearRecurso(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	recurso, err := recursobd.ValidoRecurso(id, Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error " + err.Error(), 400)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(recurso)
}