# K8S Challenge


A very simple and minimalist REST API written in Go using Gorilla Mux HTTP router package to access pods inside a given namespace with the GET /pods endpoint.

Also include `/healthz` and `/readyz` for kubernetes liveness and readiness probes, respectively.

# Required envs

* `NAMESPACE` : str

# How to run locally

* Navigate to `src` directory

* `docker build -t go-check-pods .`

* `docker run -p 80:80 -e NAMESPACE=default go-check-pods`

# Minikube Deployment Steps:

Required tools:

- Minikube
- Kubectl, Helm CLI

* Minikube namespace creation inside default minikube cluster:

Option A: Provision namespace via Terraform. Navigate to `provision` directory and run:

`terraform init`

`terraform plan -var-file="env/default.tfvars"`

`terraform apply -var-file="env/default.tfvars"`

Default namespace name is: go-check-pods-ns. This can be overriden in env/default.tfvars

Option B: Create namespace via helm install command (--create-namespace) - Known caveat is that namespace will not be removed and tracked when uninstalling helm deployments in new namespaces(helm uninstall)

* Build helm dependencies. Navigate to `go-check-pods` and run: `helm dependency build`. This should create chart.lock file and include the reloader templates inside the helm charts. Validate by running `helm template .`

* Run helm install: `helm install go-check-pods-release go-check-pods --namespace go-check-pods-ns --create-namespace`


* Deployments, configmaps, and services should be created following the command above. To access service from localhost, use minikube service built-in command.

`minikube service <app-name> -n <namespace>`
`minikube service go-check-pods -n go-check-pods-ns`


# CI/CD

The app utilizes GitHub Actions to build and push the image to Docker Hub upon code push to `main` branch. This will be followed by a `helm upgrade` command to force a helm chart upgrade with the new image version (using mutable `latest` image tagging for simplicity)

CI/CD is currently not tested and working as user is unable to post GitHub Actions secrets due to permission limitations.


# Description

This challenge is meant to evaluate software engineering and kubernetes skills, mind to apply all the best practices you know.

Your challenge will consist in developing an application that is able to interact with Kubernetes.

# Requirements

Setup a Minikube cluster

Develop a simple restful API with:

* health check endpoint
* GET `/pods` endpoint:
  * Retrieve all pods running in a given namespace.
  * The namespace should be configured using an environment variable.

Deploy the API in Minikube:

* The API should have permission only to read pods in all namespaces.
* The API should be reachable from localhost (without using port forwarding).
* The env variable that controls which namespace should be listed should be loaded from a ConfigMap.
* Use [Reloader](https://github.com/stakater/Reloader) to auto-restart the pod when the ConfigMap is updated.
* The container image should be in a public registry (like DockerHub or Github Registry).

# Considerations

* All code and configuration should be delivered in a GitHub repo containing the steps to run the project, from the cluster creation to the API testing and validation.
* Code organization and documentation clarity will be evaluated.
* Project enhancements and extrapolations are also welcome and will be a bonus (like ci/cd pipelines). Don't be shy. Show off as many skills as possible, and if implementing them is either expensive or not reasonable to do in the time you have, leave a section in the documentation explaining what you would do and how you would implement it.
* From our first interview, we learned which programming languages you are proficient in. Another bonus will be counted if you deliver code in a language you've never used (Golang is our preferred language). Our main interest is to understand how fast you learn.
* Loave the challenge description in the README, place your documentation above the Description header.
