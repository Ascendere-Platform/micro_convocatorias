package convocatoriarouters

import (
	"encoding/json"
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	"github.com/ascendere/micro-convocatorias/routers"
)

func BuscarUltimaConvocatoria(w http.ResponseWriter, r *http.Request) {

	result, status := convocatoriabd.UltimaConvocatoria(routers.Tk)
	if len(status) != 0{
		http.Error(w, "Error al leer las convocatorias" + status ,http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}