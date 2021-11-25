package convocatoriabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	anexobd "github.com/ascendere/micro-convocatorias/bd/anexo_bd"
	lineaEstrategicabd "github.com/ascendere/micro-convocatorias/bd/lineaEstrategica_bd"
	recursobd "github.com/ascendere/micro-convocatorias/bd/recurso_bd"
	resultadoEsperadobd "github.com/ascendere/micro-convocatorias/bd/resultadoEsperado_bd"
	rubricabd "github.com/ascendere/micro-convocatorias/bd/rubrica_bd"
	tipoProyectobd "github.com/ascendere/micro-convocatorias/bd/tipoProyecto_bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoConvocatorias(aprobada string, tk string) ([]*convocatoriamodels.DevuelvoConvocatoria, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("anexosConvocatoria")

	var convocatorias []*convocatoriamodels.Convocatoria
	var results []*convocatoriamodels.DevuelvoConvocatoria

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var s convocatoriamodels.Convocatoria
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}
		convocatorias = append(convocatorias, &s)

	}

	err = cur.Err()
	if err != nil {
		return results, false
	}

	for _, convocatoria := range convocatorias {
		var resultado convocatoriamodels.DevuelvoConvocatoria

		resultado.ID = convocatoria.ID
		resultado.NombreConvocatoria = convocatoria.NombreConvocatoria
		resultado.PeriodoConvocatoria = convocatoria.PeriodoConvocatoria
		resultado.Antecedentes = convocatoria.Antecedentes
		resultado.Banner = convocatoria.Banner
		resultado.Objetivos = convocatoria.Objetivos
		resultado.Estado = convocatoria.Estado
		resultado.Reconocimiento = convocatoria.Reconocimiento
		resultado.Destinatario = convocatoria.Destinatario
		resultado.Compromisos = convocatoria.Compromisos
		resultado.Plazo = convocatoria.Plazo
		resultado.Contactos = convocatoria.Contactos
		resultado.FechaCreacion = convocatoria.FechaCreacion
		resultado.CalificacionPostulacion = convocatoria.CalificacionPostulacion
		resultado.CalificacionProyecto = convocatoria.CalificacionProyecto
		resultado.CreadorConvocatoria = convocatoria.CreadorConvocatoria

		for _, anexo := range convocatoria.AnexosConvocatoria {
			objID := anexo.String()
			resultadoAnexo, errAnexo := anexobd.BuscoAnexo(objID)
			if errAnexo != nil {
				return results, false
			}
			resultado.AnexosConvocatoria = append(resultado.AnexosConvocatoria, resultadoAnexo)
		}

		for _, resultadosEsperados := range convocatoria.ResultadosEsperados {
			objID := resultadosEsperados.String()
			result, errResult := resultadoEsperadobd.BuscoResultadoEsperado(objID)
			if errResult != nil {
				return results, false
			}
			resultado.ResultadosEsperados = append(resultado.ResultadosEsperados, result)
		}

		for _, tipos := range convocatoria.TiposProyectos {
			objID := tipos.String()
			resultadoTipo, errTipo := tipoProyectobd.BuscoTipoProyecto(objID)
			if errTipo != nil {
				return results, false
			}
			resultado.TiposProyectos = append(resultado.TiposProyectos, resultadoTipo)
		}

		for _, lineas := range convocatoria.LineasEstrategicas {
			objID := lineas.String()
			resultadoLinea, errLinea := lineaEstrategicabd.BuscoLineaEstrategica(objID)
			if errLinea != nil {
				return results, false
			}
			resultado.LineasEstrategicas = append(resultado.LineasEstrategicas, resultadoLinea)
		}

		for _, rubricas := range convocatoria.RubricasConvocatoria {
			objID := rubricas.String()
			resultadoRubrica, errRubrica := rubricabd.BuscoRubrica(objID)
			if errRubrica != nil {
				return results, false
			}
			resultado.RubricasConvocatoria = append(resultado.RubricasConvocatoria, resultadoRubrica)
		}

		for _, recursos := range convocatoria.RecursosConvocatoria {
			objID := recursos.String()
			resultadoRecurso, errRecurso := recursobd.ValidoRecurso(objID, tk)
			if errRecurso != nil {
				return results, false
			}
			resultado.RecursosConvocatoria = append(resultado.RecursosConvocatoria, resultadoRecurso)
		}
		results = append(results, &resultado)
	}

	cur.Close(ctx)
	return results, true

}
