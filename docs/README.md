# Introduction

Series of labs and instructions to introduce you to containers, Github Action and Azure Container Apps.

## Description

The goal is to deploy a frontend application and an API packaged into Docker containers, containers pushed into Azure Container Registry and deployed into Azure Container Apps automatically each time we push new change into our github repository.

## Agenda

|  |  |
| :--- | :--- |
| [Lab 0](lab-0/README.md) | Pre-work - Install Docker |
| [Lab 1](lab-1/README.md) | Lab 1 - Build and push our containers to Azure Container Registry |
| [Lab 2](lab-2/README.md) | Lab 2 - Deploy our app using Azure Container Apps |
| [Lab 3](lab-3/README.md) | Lab 3 - Automate app building and deployment using Github Action to Azure Container Apps |

<!-- | [Lab 4](lab-4/README.md) | Lab 4 - Creating a Kubernetes Cluster using Azure Kubernetes Service (AKS) |
| [Lab 5](lab-5/README.md) | Lab 5 - Configuring an Ingress Controller using nginx and Let's Encrypt |
| [Lab 6](lab-6/README.md) | Lab 6 - Deploy our application into AKS using Github Action | -->

## Pre-requirements

For this workshop you must have:

* Docker
* A Github Account
* An Azure Account
* [Fork the current repository inside your Github account](https://github.com/jmaitrehenry/ga2022/fork){:target="_blank"}

## Technology Used

* [Docker CLI](https://docs.docker.com/engine/reference/commandline/cli/)
* [Docker Compose](https://docs.docker.com/compose/)
* [Github Actions](https://docs.github.com/en/actions)
* [Azure Container Apps](https://docs.microsoft.com/en-us/azure/container-apps/)

## Application to deploy

The demo app runs across two containers:

* words - a Java REST API which serves words read from a static list
* web - a Go web application which calls the API and builds words into sentences

### Run the application locally

You can start the application using docker compose: `docker compose up -d` and test the application inside your browser on [http://localhost:8080](http://localhost:8080){:target="_blank"}.
