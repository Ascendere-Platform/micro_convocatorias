package anexobd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminoAnexo(tipoID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("anexosConvocatoria")

	objID, _ := primitive.ObjectIDFromHex(tipoID)

	condicion := bson.M{
		"_id":objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}