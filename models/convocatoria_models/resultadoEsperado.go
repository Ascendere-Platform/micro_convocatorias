package convocatoriamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResultadoEsperado struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreResultado      string       `bson:"nombreResultado" json:"nombreResultado,omitempty"`
	DescripcionResultado string       `bson:"descripcionResultado" json:"descripcionResultado,omitempty"`
	PuntajeResultado     float64      `bson:"puntajeResultado" json:"puntajeResultado,omitempty"`
}