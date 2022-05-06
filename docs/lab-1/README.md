# Lab 1 - Build and Push Containers to Azure Container Registry

## Create an Azure Container Registry

### Resource Group

If needed, a new `Resource Group` can be created. It is recommended to create a new resource group for this tutorial, after which cleaning up will be easier.

In Azure Portal's [Resource groups](https://portal.azure.com/#blade/HubsExtension/BrowseResourceGroups), create a new resource group under your subscription.

### Container Registry

In Azure Portal's [Container registries](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.ContainerRegistry%2Fregistries), create a new registry.

After the registry is created, in that registry's Settings > Access Key, enable `Admin user`. This may not always be adviseable, but this enables the manual deployment in the `lab-2` tutorial.

## Build and Push the Containers

This step builds the two applications locally:

```bash
docker build -t <registry-name>.azurecr.io/words:1.0.0 ./words
docker build -t <registry-name>.azurecr.io/web:1.0.0 ./web
```

Then the following will log into Azure, and push the containers.

During the login, use the registry's admin user's name and password.

```bash
docker login <registry-name>.azurecr.io
docker push <registry-name>.azurecr.io/words:1.0.0
docker push <registry-name>.azurecr.io/web:1.0.0
```

Finally, in the registry's Services > Repositories, validate that the two Docker images were pushed successfully.

## Summary

In this step, we created a private docker registry, built two docker images and pushed them into the cloud.

Note on public and private registries: When a project contains private or sensitive data, it is important to use a private registry. Other projects, such as open-source applications, may be pushed to public registries so that everyone can pull them. [Docker Hub](https://hub.docker.com) is a popular service that offers both public and private registries.
