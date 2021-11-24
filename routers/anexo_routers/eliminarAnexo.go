package anexorouters

import (
	"net/http"

	anexobd "github.com/ascendere/micro-convocatorias/bd/anexo_bd"
)

func EliminarAnexo(w http.ResponseWriter, r *http.Request) {

	anexo := r.URL.Query().Get("id")

	if len(anexo) < 1 {
		http.Error(w, "Debe enviar el id", http.StatusBadRequest)
		return
	}

	err := anexobd.EliminoAnexo(anexo)
	if err != nil {
		http.Error(w, "Ocurrio un error"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}