package convocatoriamodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type LineaEstrategica struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	NombreLinea      string             `bson:"nombreLinea" json:"nombreLinea,omitempty"`
	DescripcionLinea string             `bson:"descripcionLinea" json:"descripcionLinea,omitempty"`
}
