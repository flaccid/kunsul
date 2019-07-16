package kunsul

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

var (
	code int = http.StatusTemporaryRedirect
)

func Serve(kubeConfig *rest.Config, configDir string, templateFile string, listenPort int, accessLog bool, ingressLabels []string) {
	log.Info("initialize kunsul")
	log.Debug("debug logging enabled")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/", healthCheckHandler)

	log.WithFields(log.Fields{
		"configDir":     configDir,
		"port":          listenPort,
		"templateFile":  templateFile,
		"ingressLabels": ingressLabels,
	}).Info("listening for requests")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, kubeConfig, configDir, templateFile, ingressLabels)
	})

	http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), nil)
}
