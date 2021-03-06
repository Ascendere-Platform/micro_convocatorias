package convocatoriarouters

import (
	"encoding/json"
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	"github.com/ascendere/micro-convocatorias/routers"
)

func ListarConvocatorias(w http.ResponseWriter, r *http.Request) {

	result, status := convocatoriabd.ListoConvocatorias(routers.Tk)
	if !status {
		http.Error(w, "Error al leer las convocatorias", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}