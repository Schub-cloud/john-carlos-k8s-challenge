# K8S Challenge

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
