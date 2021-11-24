package resultadoEsperadorouters

import (
	"encoding/json"
	"net/http"

	resultadoEsperadobd "github.com/ascendere/micro-convocatorias/bd/resultadoEsperado_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func RegistrarResultadoEsperado(w http.ResponseWriter, r *http.Request) {
	var resultado convocatoriamodels.ResultadoEsperado

	err := json.NewDecoder(r.Body).Decode(&resultado)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	_, status, err := resultadoEsperadobd.RegistroResultadoEsperado(resultado)

	if err != nil {
		http.Error(w, "Ocurrio un error al insertar una nuevo resultado esperado", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logaro registrar una nuevo resultado esperado", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}