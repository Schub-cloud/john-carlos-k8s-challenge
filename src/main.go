package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/gorilla/mux"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/pods", GetPods).Methods("GET")
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/readyz", Readyz).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func Readyz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func GetPods(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get Kubernetes in-cluster config: %v", err), http.StatusInternalServerError)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create Kubernetes client: %v", err), http.StatusInternalServerError)
		return
	}

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		http.Error(w, "Namespace not specified. Please set the NAMESPACE environment variable.", http.StatusBadRequest)
		return
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(r.Context(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list pods: %v", err), http.StatusInternalServerError)
		return
	}

    if len(pods.Items) == 0 {
        fmt.Fprintf(w, "No pods in namespace %s\n", namespace)
        return
    }

	for _, pod := range pods.Items {
		fmt.Fprintf(w, "%s\n", pod.Name)
	}
}