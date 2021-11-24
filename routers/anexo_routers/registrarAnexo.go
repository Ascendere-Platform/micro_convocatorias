package anexorouters

import (
	"encoding/json"
	"net/http"
	"time"

	anexobd "github.com/ascendere/micro-convocatorias/bd/anexo_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarAnexo(w http.ResponseWriter, r *http.Request) {
	var anexo convocatoriamodels.Anexo

	err := json.NewDecoder(r.Body).Decode(&anexo)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	anexo.FechaCreacion = time.Now()

	_, status, err := anexobd.RegistroAnexo(anexo)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar un nuevo Anexo", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar un nuevo Anexo", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}