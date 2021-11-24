package lineaEstrategicabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroLineaEstrategica(r convocatoriamodels.LineaEstrategica) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("lineaEstrategica")


	registro := convocatoriamodels.LineaEstrategica{
		ID:                 primitive.NewObjectID(),
		NombreLinea:      r.NombreLinea,
		DescripcionLinea:  r.DescripcionLinea,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}