package rubricarouters

import (
	"net/http"

	rubricabd "github.com/ascendere/micro-convocatorias/bd/rubrica_bd"
)

func EliminarRubrica(w http.ResponseWriter, r *http.Request) {

	rubrica := r.URL.Query().Get("id")

	if len(rubrica) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := rubricabd.EliminoRubrica(rubrica)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}