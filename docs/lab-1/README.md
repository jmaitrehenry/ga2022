# Lab 1 - Build and push our containers to Azure Container Registry

## Create an Azure Container Registry

## Build the containers locally

```bash
cd web
docker build -t <registry-name>.azurecr.io/web:1.0.0 .

cd ../words
docker build -t <registry-name>.azurecr.io/words:1.0.0 .
```

## Log into ACR and push our docker images

```bash
docker login <registry-name>.azurecr.io
docker push <registry-name>.azurecr.io/web:1.0.0
docker build -t <registry-name>.azurecr.io/words:1.0.0
```

## Summary

We created a private docker registry, build 2 docker images and push them into our private registry.
Instead of using the private registry of Azure, we could use the [public registry of Docker ](https://hub.docker.com) because our image doesn't have anything priavte or sensible too.