package convocatoriamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Anexo struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreAnexo   string             `bson:"nombreAnexo" json:"nombreAnexo,omitempty"`
	Link          string             `bson:"link" json:"link,omitempty"`
	FechaCreacion string             `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
}
