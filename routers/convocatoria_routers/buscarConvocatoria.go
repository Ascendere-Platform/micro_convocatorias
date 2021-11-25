package convocatoriarouters

import (
	"encoding/json"
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	"github.com/ascendere/micro-convocatorias/routers"
)

func BuscarConvocatoria(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err := convocatoriabd.BuscoConvocatoria(id, routers.Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una convocatoria "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}