# Lab 3 - Automate Build and Deployment using Github Action to Azure Container Apps

The demo application will be deployed automatically using GitHub actions in this step.

## Create a Service Principal for GitHub

Instead of the main Azure account, it is recommended to create a `Service Principal`, which will be used by GitHub to connect to Azure.

The following will describe the steps necessary to create that service principal, and grant it the `Contributor` role at the resource group level. This role will be inherited by all resources under the resource group.

Usually, the permissions granted would be more fine-grained. For example, the service principal would instead have the `AcrPush` role in the registry itself, but granting a wide-scoped role is easier during a demo.

### From the Terminal

Using the `az` command, create a service principal:

```bash
az ad sp create-for-rbac \
  --name "github-bot" \
  --role "contributor" \
  --scopes /subscriptions/c235daef-a49b-4feb-b432-9ea0b7cbbc6b/resourceGroups/064b73c3
```

The output will contain the service principal's ID (`appId`) and its password. Take note of them.

### From Azure Portal

In Azure Portal's [App registrations](https://portal.azure.com/#blade/Microsoft_AAD_RegisteredApps/ApplicationsListBlade), create a new registration.

Then, in that service principal's Manage > Certificates and secrets, create a new client secret, and take note of the generated secret value.

In Azure Portal's [Resource groups](https://portal.azure.com/#blade/HubsExtension/BrowseResourceGroups), select the resource group used for this demo. In Access Control (IAM), add a new role assignment with the `Contributor` role to the service principal created earlier.

## Add GitHub Secrets

In your GitHub repository, go to Settings > Secrets, and add the following Action secrets:

* `AZURE_CLIENT_ID`: the Application (client) ID
* `AZURE_CLIENT_SECRET`: the generated secret value
* `AZURE_CREDENTIALS`: a JSON with the following value (placeholders replaced with actual values):

```json
{
    "tenantId": "<Tenant ID>",
    "subscriptionId": "<Subscription ID>",
    "clientId": "<Service Principal ID>",
    "clientSecret": "<Service Principal Secret>",
    "resourceManagerEndpointUrl": "https://management.azure.com/"
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

## Summary

We created a limited account that could interact with our docker registry and our container applications. After that, we built and deployed our application only when they have some changes.

You have now a full Continuous Deployment Pipeline running into Github!
