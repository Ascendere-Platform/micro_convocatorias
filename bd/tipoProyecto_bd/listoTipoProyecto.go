package tipoProyectobd

import (
	"context"
	"time"

	"github.com/ascendere/micro-convocatorias/bd"
	convocatoriamodels "github.com/ascendere/micro-convocatorias/models/convocatoria_models"
	"go.mongodb.org/mongo-driver/bson"
)

func ListoTipoRecurso() ([]*convocatoriamodels.TipoProyecto, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := bd.MongoCN.Database("Convocatoria")
	col := db.Collection("tiposProyectos")

	var results []*convocatoriamodels.TipoProyecto

	query := bson.M{}

	cur, err := col.Find(ctx, query)
	if err != nil {
		return results, false
	}

	for cur.Next(ctx) {
		var s convocatoriamodels.TipoProyecto
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}
			results = append(results, &s)

	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true

}