package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func init(){
	prometheus.MustRegister(jobsInQueueOne)
	prometheus.MustRegister(jobsInQueueTwo)
}

var jobsInQueueOne = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "jobs_in_queue_one",
		Help: "Current number of jobs in the queue one",
	},
)

var jobsInQueueTwo = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "jobs_in_queue_two",
		Help: "Current number of jobs in the queue two",
	},
	[]string{"job_type"},
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	var count int
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		count++
		fmt.Println(count)
		jobsInQueueOne.Inc()
		jobsInQueueTwo.WithLabelValues("get_method").Inc()
	})
	panic(http.ListenAndServe(":8080", nil))
}




