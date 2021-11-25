package convocatoriarouters

import (
	"encoding/json"
	"log"
	"net/http"

	convocatoriabd "github.com/ascendere/micro-convocatorias/bd/convocatoria_bd"
	"github.com/ascendere/micro-convocatorias/routers"
)

func BuscarConvocatoria(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	informacion, err, mensaje := convocatoriabd.BuscoConvocatoria(id, routers.Tk)

	if err != nil {
		http.Error(w, "Ocurrio un error al buscar una convocatoria "+err.Error()+" " + mensaje, 400)
		return
	}

	log.Print(informacion)

	log.Print(routers.Tk)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(informacion)

}