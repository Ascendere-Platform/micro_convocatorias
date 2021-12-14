package convocatoriabd

import (
	"time"

	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
)

func UltimaConvocatoria(tk string) (convocatoriamodels.DevuelvoConvocatoria, string) {
	var ultima convocatoriamodels.DevuelvoConvocatoria

	resultConsulta, status := ListoConvocatorias(tk)

	if !status {
		return ultima, "Convocatorias no encontradas"
	}

	var auxTime time.Time
	var pos int

	for i, convocatoria := range resultConsulta {
		if convocatoria.Estado {
			if convocatoria.FechaPublicacion.Before(auxTime) {
				auxTime = convocatoria.FechaPublicacion
				pos = i
			}
		}
	}

	ultima = *resultConsulta[pos]

	return ultima, ""

}
