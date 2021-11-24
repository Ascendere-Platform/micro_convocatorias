package tipoProyectobd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroTipoProyecto(r convocatoriamodels.TipoProyecto) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("tiposProyectos")

	registro := convocatoriamodels.TipoProyecto{
		ID:                 primitive.NewObjectID(),
		TipoProyecto:      r.TipoProyecto,
		DecripcionTipo:  r.DecripcionTipo,
		Presupuesto: r.Presupuesto,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}