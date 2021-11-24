package resultadoEsperadobd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegistroRubrica(r convocatoriamodels.ResultadoEsperado) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("resultadosEsperados")


	registro := convocatoriamodels.ResultadoEsperado{
		ID:                 primitive.NewObjectID(),
		NombreResultado:      r.NombreResultado,
		DescripcionResultado:  r.DescripcionResultado,
		PuntajeResultado: r.PuntajeResultado,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}