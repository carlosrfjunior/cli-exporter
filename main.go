package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/carlosrfjunior/cli-exporter/providers"
	"github.com/carlosrfjunior/cli-exporter/providers/awscli"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddr  = flag.String("web.listen-addres", ":9100", "The address to listen on for HTTP requests.")
	metricsPath = flag.String("web.telemetry-path", "/metrics", "The address to listen on for HTTP requests.")
	configFile  = flag.String("config.file", "clis.yaml", "Provider configuration file and its functions.")
)

func main() {

	p := awscli.Aws{
		Name: "AWS Provider",
	}

	flag.Parse()

	providers.Start(p)

	http.Handle(*metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
