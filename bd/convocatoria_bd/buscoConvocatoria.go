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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoConvocatoria(id string, tk string) (convocatoriamodels.DevuelvoConvocatoria, error, string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("convocatoria")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": objID}

	var resultado convocatoriamodels.DevuelvoConvocatoria
	var convocatoria convocatoriamodels.Convocatoria

	err := col.FindOne(ctx, condicion).Decode(&convocatoria)

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
			return resultado, errAnexo, "error anexos"
		}
		resultado.AnexosConvocatoria = append(resultado.AnexosConvocatoria, resultadoAnexo)
	}

	for _, resultadosEsperados := range convocatoria.ResultadosEsperados {
		objID := resultadosEsperados.String()
		result, errResult := resultadoEsperadobd.BuscoResultadoEsperado(objID)
		if errResult != nil {
			return resultado, errResult, "error resultados"
		}
		resultado.ResultadosEsperados = append(resultado.ResultadosEsperados, result)
	}

	for _, tipos := range convocatoria.TiposProyectos {
		objID := tipos.String()
		resultadoTipo, errTipo := tipoProyectobd.BuscoTipoProyecto(objID)
		if errTipo != nil {
			return resultado, errTipo, "error tipos"
		}
		resultado.TiposProyectos = append(resultado.TiposProyectos, resultadoTipo)
	}

	for _, lineas := range convocatoria.LineasEstrategicas {
		objID := lineas.String()
		resultadoLinea, errLinea := lineaEstrategicabd.BuscoLineaEstrategica(objID)
		if errLinea != nil {
			return resultado, errLinea, "error lineas"
		}
		resultado.LineasEstrategicas = append(resultado.LineasEstrategicas, resultadoLinea)
	}

	for _, rubricas := range convocatoria.RubricasConvocatoria {
		objID := rubricas.String()
		resultadoRubrica, errRubrica := rubricabd.BuscoRubrica(objID)
		if errRubrica != nil {
			return resultado, errRubrica, "error rubricas"
		}
		resultado.RubricasConvocatoria = append(resultado.RubricasConvocatoria, resultadoRubrica)
	}

	for _, recursos := range convocatoria.RecursosConvocatoria {
		objID := recursos.String()
		resultadoRecurso, errRecurso := recursobd.ValidoRecurso(objID, tk)
		if errRecurso != nil {
			return resultado, errRecurso, "error recursos"
		}
		resultado.RecursosConvocatoria = append(resultado.RecursosConvocatoria, resultadoRecurso)
	}

	if err != nil {
		return resultado, err, "No encuentra la convocatoria"
	}
	return resultado, err, ""
}
