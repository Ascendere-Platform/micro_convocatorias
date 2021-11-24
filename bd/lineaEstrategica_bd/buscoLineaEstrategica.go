package lineaEstrategicabd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoLineaEstrategica(id string) (convocatoriamodels.LineaEstrategica, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("lineaEstrategica")

	objID, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id":objID}
	
	var resultado convocatoriamodels.LineaEstrategica

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil{
		return resultado, err
	}
	return resultado, err
}