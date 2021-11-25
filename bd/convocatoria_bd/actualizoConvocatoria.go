package convocatoriabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ActualizoConvocatoria(u convocatoriamodels.Convocatoria) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("convocatoria")

	var resultado convocatoriamodels.Convocatoria

	error := col.FindOne(ctx, bson.M{"_id": u.ID}).Decode(&resultado)

	if error != nil {
		return false, error
	}

	registro := make(map[string]interface{})

	if len(u.NombreConvocatoria) > 0 {
		registro["nombreConvocatoria"] = u.NombreConvocatoria
	}

	if len(u.PeriodoConvocatoria) > 0 {
		registro["periodoConvocatoria"] = u.PeriodoConvocatoria
	}

	if len(u.Antecedentes) > 0 {
		registro["antecendentes"] = u.Antecedentes
	}

	if len(u.Objetivos) > 0 {
		registro["objetivos"] = u.Objetivos
	}

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	if len(u.Destinatario) > 0 {
		registro["destinatario"] = u.Destinatario
	}

	if len(u.Reconocimiento) > 0 {
		registro["reconocimiento"] = u.Reconocimiento
	}

	if len(u.Compromisos)>0{
		registro["compromisos"] = u.Compromisos
	}

	if u.Plazo.IsZero() {
		registro["plazo"] = u.Plazo
	}

	if len(u.Contactos) > 0 {
		registro["contactos"] = u.Contactos
	}

	if u.CalificacionPostulacion > 0 {
		registro["calificacionPostulacion"] = u.CalificacionPostulacion
	}

	if u.CalificacionProyecto > 0 {
		registro["calificacionProyecto"] = u.CalificacionProyecto
	}

	if u.Estado {
		registro["estado"] = u.Estado
	}

	if !u.Estado {
		registro["estado"] = u.Estado
	}

	updtString := bson.M{
		"$set": registro,
	}

	filtro := bson.M{"_id": bson.M{"$eq": u.ID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}
