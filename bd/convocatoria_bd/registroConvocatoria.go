package convocatoriabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroConvocatoria(r convocatoriamodels.Convocatoria) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("convocatoria")

	registro := convocatoriamodels.Convocatoria{
		ID:                 primitive.NewObjectID(),
		NombreConvocatoria: r.NombreConvocatoria,
		PeriodoConvocatoria: r.PeriodoConvocatoria,
		Antecedentes: r.Antecedentes,
		Objetivos: r.Objetivos,
		Banner: r.Banner,
		Estado: false,
		Destinatario: r.Destinatario,
		Reconocimiento: r.Reconocimiento,
		Compromisos: r.Compromisos,
		Plazo: r.Plazo,
		Contactos: r.Contactos,
		FechaCreacion: r.FechaCreacion,
		CalificacionPostulacion: r.CalificacionPostulacion,
		CalificacionProyecto: r.CalificacionProyecto,
		CreadorConvocatoria: r.CreadorConvocatoria,
		AnexosConvocatoria: r.AnexosConvocatoria,
		ResultadosEsperados: r.ResultadosEsperados,
		TiposProyectos: r.TiposProyectos,
		LineasEstrategicas: r.LineasEstrategicas,
		RubricasConvocatoria: r.RubricasConvocatoria,
		RecursosConvocatoria: r.RecursosConvocatoria,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}