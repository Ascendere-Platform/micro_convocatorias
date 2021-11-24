package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/ascendere/micro-convocatorias/middlew"
	anexorouters "github.com/ascendere/micro-convocatorias/routers/anexo_routers"
	lineaEstrategicarouters "github.com/ascendere/micro-convocatorias/routers/lineaEstrategica_routers"
	resultadoEsperadorouters "github.com/ascendere/micro-convocatorias/routers/resultadoEsperado_routers"
	rubricarouters "github.com/ascendere/micro-convocatorias/routers/rubrica_routers"
	tipoProyectorouters "github.com/ascendere/micro-convocatorias/routers/tipoProyecto_routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {

	router := mux.NewRouter()

	//Llamada al CRUD de Anexos
	router.HandleFunc("/registrarAnexo", middlew.ChequeoBD(middlew.ValidoJWT(anexorouters.RegistrarAnexo))).Methods("POST")
	router.HandleFunc("/eliminarAnexo", middlew.ChequeoBD(middlew.ValidoJWT(anexorouters.EliminarAnexo))).Methods("DELETE")
	router.HandleFunc("/buscarAnexo", middlew.ChequeoBD(middlew.ValidoJWT(anexorouters.BuscarAnexo))).Methods("GET")
	router.HandleFunc("/listarAnexos", middlew.ChequeoBD(middlew.ValidoJWT(anexorouters.ListarAnexos))).Methods("GET")

	//Llamada al CRUD de Lineas Estrategicas
	router.HandleFunc("/registrarLinea", middlew.ChequeoBD(middlew.ValidoJWT(lineaEstrategicarouters.RegistrarLineaEstrategica))).Methods("POST")
	router.HandleFunc("/eliminarLinea", middlew.ChequeoBD(middlew.ValidoJWT(lineaEstrategicarouters.EliminarLineaEstrategica))).Methods("DELETE")
	router.HandleFunc("/buscarLinea", middlew.ChequeoBD(middlew.ValidoJWT(lineaEstrategicarouters.BuscarLineaEstrategica))).Methods("GET")
	router.HandleFunc("/listarLineas", middlew.ChequeoBD(middlew.ValidoJWT(lineaEstrategicarouters.ListarLineasEstrategicas))).Methods("GET")


	//Llamada al CRUD de Resultados Esperados
	router.HandleFunc("/registrarResultadoEsperado", middlew.ChequeoBD(middlew.ValidoJWT(resultadoEsperadorouters.RegistrarResultadoEsperado))).Methods("POST")
	router.HandleFunc("/eliminarResultadoEsperado", middlew.ChequeoBD(middlew.ValidoJWT(resultadoEsperadorouters.EliminarResultadoEsperado))).Methods("DELETE")
	router.HandleFunc("/buscarResultadoEsperado", middlew.ChequeoBD(middlew.ValidoJWT(resultadoEsperadorouters.BuscarResultadoEsperado))).Methods("GET")
	router.HandleFunc("/listarResultadosEsperados", middlew.ChequeoBD(middlew.ValidoJWT(resultadoEsperadorouters.ListarResultadosEsperados))).Methods("GET")

	//Llamada al CRUD de Rubricas
	router.HandleFunc("/registrarRubrica", middlew.ChequeoBD(middlew.ValidoJWT(rubricarouters.RegistrarRubrica))).Methods("POST")
	router.HandleFunc("/eliminarRubrica", middlew.ChequeoBD(middlew.ValidoJWT(rubricarouters.EliminarRubrica))).Methods("DELETE")
	router.HandleFunc("/buscarRubrica", middlew.ChequeoBD(middlew.ValidoJWT(rubricarouters.BuscarRubrica))).Methods("GET")
	router.HandleFunc("/listarRubricas", middlew.ChequeoBD(middlew.ValidoJWT(rubricarouters.ListarRubricas))).Methods("GET")

	//Llamada al CRUD de Tipos de Proyectos
	router.HandleFunc("/registrarTipoProyecto", middlew.ChequeoBD(middlew.ValidoJWT(tipoProyectorouters.RegistrarTipoProyecto))).Methods("POST")
	router.HandleFunc("/eliminarTipoProyecto", middlew.ChequeoBD(middlew.ValidoJWT(tipoProyectorouters.EliminarTipoProyecto))).Methods("DELETE")
	router.HandleFunc("/buscarTipoProyecto", middlew.ChequeoBD(middlew.ValidoJWT(tipoProyectorouters.BuscarTipoProyecto))).Methods("GET")
	router.HandleFunc("/listarTiposProyectos", middlew.ChequeoBD(middlew.ValidoJWT(tipoProyectorouters.ListarTiposProyectos))).Methods("GET")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8082"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
