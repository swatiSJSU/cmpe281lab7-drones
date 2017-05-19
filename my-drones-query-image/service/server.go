package service

import (
	"github.com/cloudnativego/cfmgo"
	"github.com/cloudnativego/drones-events/mongo"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	repo := initRepository()
	initRoutes(mx, formatter, repo)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo eventRepository) {
	mx.HandleFunc("/drones/{droneId}/lastTelemetry", lastTelemetryHandler(formatter, repo)).Methods("GET")
	mx.HandleFunc("/drones/{droneId}/lastAlert", lastAlertHandler(formatter, repo)).Methods("GET")
	mx.HandleFunc("/drones/{droneId}/lastPosition", lastPositionHandler(formatter, repo)).Methods("GET")
}

func initRepository() (repo eventRepository) {
	//appEnv, _ := cfenv.Current()
	//dbServiceURI, err := cftools.GetVCAPServiceProperty("mongo-eventrollup", "url", appEnv)
	//if err != nil || len(dbServiceURI) == 0 {
		//if err != nil {
			//fmt.Printf("\nError retreieving database configuration: %v\n", err)
		//}
		//fmt.Println("MongoDB was not detected, using fake repository THIS IS BAD...")
		////repo = NewFakeRepository()
	//} else {
    dbServiceURI := "mongodb://mongo:27017/events-database"
		telemetryCollection := cfmgo.Connect(cfmgo.NewCollectionDialer, dbServiceURI, "telemetry")
		positionsCollection := cfmgo.Connect(cfmgo.NewCollectionDialer, dbServiceURI, "positions")
		alertsCollection := cfmgo.Connect(cfmgo.NewCollectionDialer, dbServiceURI, "alerts")
		repo = mongo.NewEventRollupRepository(positionsCollection, alertsCollection, telemetryCollection)
	//}

	return
}
