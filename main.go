package main

import (
	"log"

	"github.com/ascendere/micro-convocatorias/bd"
	"github.com/ascendere/micro-convocatorias/handlers"
)

func main (){
	if bd.ChequeoConnection() == 0{
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}