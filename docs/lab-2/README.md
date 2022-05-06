# Lab 2 - Deploy Applications using Azure Container Apps

## Create a Container App

The demo application will be manually deployed in this step.

### words:1.0.0

In Azure Portal's [Container Apps](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.App%2FcontainerApps), create a new application.

In the `Basics` section, name the application `words`. A new container app environment can also be created.

In the `App settings` section, uncheck the Use quickstart image checkbox, name the container `words`, and select your `words:1.0.0` image from the registry. Enable the HTTP ingress, accepting traffic from anywhere, and set the port to `8080`.

After the application is deployed, in its Overview, visit the application's URL, appending a `/verbs` at the end of the path. This should result in a JSON with a `word` attribute.

### web:1.0.0

Back in Azure Portal's [Container Apps](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.App%2FcontainerApps), create another new application.

In the `Basics` section, name the application `web`, and select the same container app environment as the one that was created in the previous step.

In the `App settings` section, uncheck the Use quickstart image checkbox, name the container `web`, and select your `web:1.0.0` image from the registry. Enable the HTTP ingress, accepting traffic from anywhere, and set the port to `80`.

After the application is deployed, in its Application > Containers, edit the container and add an environment variable `WORD_API_URL` with value set to the `words`'s application's URL.

After the `web` application is redeployed, visit its application URL. Lego blocks should appear on screen with words on them.

## Summary

In this step, we created two container apps based on the images that were previously built.
