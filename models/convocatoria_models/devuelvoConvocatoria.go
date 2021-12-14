package convocatoriamodels

import (
	"time"

	recursosmodels "github.com/ascendere/micro-convocatorias/models/recursos_models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoConvocatoria struct {
	ID                      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	NombreConvocatoria      string             `bson:"nombreConvocatoria" json:"nombreConvocatoria,omitempty"`
	PeriodoConvocatoria     string             `bson:"periodoConvocatoria" json:"periodoConvocatoria,omitempty"`
	Antecedentes            string             `bson:"antecedentes" json:"antecedentes,omitempty"`
	Objetivos               string             `bson:"objetivos" json:"objetivos,omitempty"`
	Banner                  string             `bson:"banner" json:"banner,omitempty"`
	Estado                  bool               `bson:"estado" json:"estado,omitempty"`
	Destinatario            string             `bson:"destinatario" json:"destinatario,omitempty"`
	Reconocimiento          string             `bson:"reconocimiento" json:"reconocimiento,omitempty"`
	Compromisos             string             `bson:"compromisos" json:"compromisos,omitempty"`
	Plazo                   time.Time          `bson:"plazo" json:"plazo,omitempty"`
	Contactos               string             `bson:"contactos" json:"contactos,omitempty"`
	FechaCreacion           time.Time          `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
	CalificacionPostulacion float64            `bson:"calificacionPostulacion" json:"calificacionPostulacion,omitempty"`
	CalificacionProyecto    float64            `bson:"calificacionProyecto" json:"calificacionProyecto,omitempty"`
	FechaPublicacion        time.Time          `bson:"fechaPublicacion" json:"fechaPublicacion,omitempty"`

	CreadorConvocatoria struct {
		Email  string             `json:"email"`
		ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	}
	AnexosConvocatoria   []Anexo                          `bson:"anexosConvocatoria" json:"anexosConvocatoria,omitempty"`
	ResultadosEsperados  []ResultadoEsperado              `bson:"resultadosEsperados" json:"resultadosEsperados,omitempty"`
	TiposProyectos       []TipoProyecto                   `bson:"tiposProyectos" json:"tiposProyectos,omitempty"`
	LineasEstrategicas   []LineaEstrategica               `bson:"lineasEstrategicas" json:"lineasEstrategicas,omitempty"`
	RubricasConvocatoria []Rubrica                        `bson:"rubricasConvocatoria" json:"rubricasConvocatoria,omitempty"`
	RecursosConvocatoria []recursosmodels.ConsultoRecurso `bson:"recursosConvocatoria" json:"recursosConvocatoria,omitempty"`
}
