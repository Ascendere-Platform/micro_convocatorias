package anexorouters

import (
	"encoding/json"
	"net/http"

	anexobd "github.com/ascendere/micro-convocatorias/bd/anexo_bd"
)

func BuscarAnexo(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := anexobd.BuscoAnexo(id)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar un anexo ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}