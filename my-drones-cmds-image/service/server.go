package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/cloudnativego/drones-cmds/fakes"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer(appEnv *cfenv.App) *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	positionDispatcher := buildDispatcher("positions", appEnv)
	telemetryDispatcher := buildDispatcher("telemetry", appEnv)
	alertDispatcher := buildDispatcher("alerts", appEnv)

	initRoutes(mx, formatter, telemetryDispatcher, alertDispatcher, positionDispatcher)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, telemetryDispatcher queueDispatcher, alertDispatcher queueDispatcher, positionDispatcher queueDispatcher) {
	mx.HandleFunc("/api/cmds/telemetry", addTelemetryHandler(formatter, telemetryDispatcher)).Methods("POST")
	mx.HandleFunc("/api/cmds/alerts", addAlertHandler(formatter, alertDispatcher)).Methods("POST")
	mx.HandleFunc("/api/cmds/positions", addPositionHandler(formatter, positionDispatcher)).Methods("POST")
}

func resolveAMQPURL(appEnv *cfenv.App) string {
	//url, err := cftools.GetVCAPServiceProperty("rabbit", "url", appEnv)
	//if err != nil {
		//fmt.Println("Failed to detect bound service for rabbit. Falling back to in-memory dispatcher (fake)")
		//return "fake://foo"
	//}
	//if len(url) < 10 {
		//fmt.Printf("URL detected for bound rabbit service not valid, was '%s'. Falling back to in-memory fake.\n", url)
		//return "fake://foo"
	//}
  url := "amqp://guest:guest@rabbit:5672/"
	return url
}

func buildDispatcher(queueName string, appEnv *cfenv.App) queueDispatcher {
	url := resolveAMQPURL(appEnv)
	if strings.Compare(url, "fake://foo") == 0 {
		fmt.Printf("Building fake dispatcher for queue '%s'", queueName)
		return fakes.NewFakeQueueDispatcher()
	}
	return createAMQPDispatcher(queueName, url)
}

func createAMQPDispatcher(queueName string, url string) queueDispatcher {
	fmt.Printf("\nUsed URL (%s) for Rabbit.\n", url)

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	dispatcher := NewAMQPDispatcher(ch, q.Name, false)
	return dispatcher
}

// If we did detect a bound service, failing to connect to it should be a fatal,
// crash-inducing error.
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
