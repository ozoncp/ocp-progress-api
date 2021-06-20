package metrics

import "github.com/prometheus/client_golang/prometheus"

var createCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "create_counter",
		Help: "Number of created classrooms",
	},
)

var updateCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "update_counter",
		Help: "Number of updated classrooms",
	},
)

var removeCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "remove_counter",
		Help: "Number of removed classrooms",
	},
)

func RegisterMetrics() {
	prometheus.MustRegister(createCounter)
	prometheus.MustRegister(updateCounter)
	prometheus.MustRegister(removeCounter)
}

func IncCreateCounter() {
	createCounter.Inc()
}

func IncUpdateCounter() {
	updateCounter.Inc()
}

func IncRemoveCounter() {
	removeCounter.Inc()
}
