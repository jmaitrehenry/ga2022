# Lab 3 - Automate Build and Deployment using Github Action to Azure Container Apps

The demo application will be deployed automatically using GitHub actions in this step.

## Create a Service Principal for GitHub

Instead of the main Azure account, it is recommended to create a `Service Principal`, which will be used by GitHub to connect to Azure.

The following will describe the steps necessary to create that service principal, and grant it roles at the Container Registry and Container Apps level.

### From the Terminal

Using the `az` command, create a service principal:

```bash
az ad sp create-for-rbac \
  --name "github-bot"
```

Will produce an output like this:
```json
{
  "appId": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
  "displayName": "github-bot",
  "password": "<random string>",
  "tenant": "<tenant-id>"
}
```
The output will contain the service principal's ID (`appId`) and its password. Take note of them.

### Add role to ACR

Back in Azure Portal's [Container Registry](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.ContainerRegistry%2Fregistries){:target="_blank"}, select your registry and got to Access control > Add > Role Assignment.

Choose `AcrPush` role, and, in the Member tab, select your Service Principal.

> __Note__: You need to search for your Service Principal name as, by default, Azure only show you a list of Azure Active Directory members

![type:video](../assets/add-acrpush-role-to-github-bot-sp.mp4)

#### Add role to Ressource

Go in Azure Portal's [Resource Group](https://portal.azure.com/#blade/HubsExtension/BrowseResourceGroups){:target="_blank"}, select the web app and got to Access control > Add > Role Assignment.

This time, choose `Contributor` role, and, in the Member tab, select your Service Principal.

> __Note__: Usually, you should never add the contributor role to the resource group directly, we try to give more granularity. But, actually, for Container App, we don't have other choice.

## Add GitHub Secrets

In your GitHub repository, go to Settings > Secrets, and add the following Action secrets:

* `AZURE_CLIENT_ID`: the Service Principale ID (`appId` from the previous outpu)
* `AZURE_CLIENT_SECRET`: the Service Principale password
* `AZURE_CREDENTIALS`: a JSON with the following value (placeholders replaced with actual values):

```json
{
    "tenantId": "<tenant ID>",
    "subscriptionId": "<subscription ID>",
    "clientId": "<Service Principale ID>",
    "clientSecret": "<Service Principale password>"
}
```

While you are in the GitHub interface, visit the Actions section. If prompted, enable GitHub workflows, otherwise the next step will not trigger any actions.

## Add Github Actions

The `azure-deploy.yml` file contains the GitHub workflow that is triggered when the `main` branch is pushed.

Fill the `jobs.deploy.env` variables with their respective values:

* `DOCKER_REGISTRY`: The container registry's name, followed by `.azurecr.io`
* `AZURE_RESOURCE_GROUP_NAME`: The resource group's name
* `AZURE_APP_NAME_WORDS`: The container app's name for the Words application
* `AZURE_APP_NAME_WEB`: The container app's name for the Web application

Now, create a directory for the Github Action:
```bash
mkdir -p .github/workflows
```

And copy or move the github action into it:
```bash
mv azure-deploy.yml .github/workflows/
```

## Push GitHub Workflow

The changes to the GitHub workflow can now be commited and pushed to the remote repository. The workflow should start automatically, and can be found in the Actions section.

## Make a change in the web application

For example, you can edit the title of the page. Open the `web/static/index.html` file and change `<title>Global Azure 2022</title>` for adding your name.
Commit and push your change and check if the web application build and deploy. Once the Github Action finish, check your application to validate the change.

## Only build and deploy the application with a change

It's fine to rebuild the two apps every time the repo have a change but we can do something better: only build and deploy if the code op the app was change.
For that, duplicate the Github Action file, and rename the other one:

```bash
cp .github/workflows/azure-deploy.yml .github/workflows/azure-deploy-web.yml
mv .github/workflows/azure-deploy.yml .github/workflows/azure-deploy-words.yml
```

Now, on each file, change the `on` block to add a path like this:
```diff
on:
  push:
    branches: 
      - main
+   paths:
+     - 'web/**'
```

Commit the change and it's done!

## Summary

We created a limited account that could interact with our docker registry and our container applications. After that, we built and deployed our application only when they have some changes.

You have now a full Continuous Deployment Pipeline running into Github!
