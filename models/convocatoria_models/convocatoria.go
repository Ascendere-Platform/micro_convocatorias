package convocatoriamodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Convocatoria struct {
	ID                  primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	NombreConvocatoria  string             `bson:"nombreConvocatoria" json:"nombreConvocatoria,omitempty"`
	PeriodoConvocatoria string             `bson:"periodoConvocatoria" json:"periodoConvocatoria,omitempty"`
	Antecedentes        string             `bson:"antecedentes" json:"antecedentes,omitempty"`
	Objetivos           string             `bson:"objetivos" json:"objetivos,omitempty"`
	Banner              string             `bson:"banner" json:"banner,omitempty"`
	Estado              bool               `bson:"estado" json:"estado,omitempty"`
	Destinatario        string             `bson:"destinatario" json:"destinatario,omitempty"`
	Reconocimiento      string             `bson:"reconocimiento" json:"reconocimiento,omitempty"`
	Compromisos         string             `bson:"compromisos" json:"compromisos,omitempty"`
	Plazo               time.Time          `bson:"plazo" json:"plazo,omitempty"`
	Contactos           string             `bson:"contactos" json:"contactos,omitempty"`
	FechaCreacion       time.Time          `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
	CreadorConvocatoria struct {
		Email  string             `json:"email"`
		ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	}
	AnexosConvocatoria   []primitive.ObjectID `bson:"anexosConvocatoria" json:"anexosConvocatoria,omitempty"`
	ResultadosEsperados  []primitive.ObjectID `bson:"resultadosEsperados" json:"resultadosEsperados,omitempty"`
	TiposProyectos       []primitive.ObjectID `bson:"tiposProyectos" json:"tiposProyectos,omitempty"`
	LineasEstrategicas   []primitive.ObjectID `bson:"lineasEstrategicas" json:"lineasEstrategicas,omitempty"`
	RubricasConvocatoria []primitive.ObjectID `bson:"rubricasConvocatoria" json:"rubricasConvocatoria,omitempty"`
	RecursosConvocatoria []primitive.ObjectID `bson:"recursosConvocatoria" json:"recursosConvocatoria,omitempty"`
}
