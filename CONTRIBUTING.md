# Contribution Guide
## Autogeneration of SDKs

1. All SDKs are autogenerated using OpenAPI Generator project.
2. Bot accounts creates PR with the new versions of the SDK that needs be reviewed by the team
3. Once PR is merged developers are responsible for creating tagged release 


## Release process

1. Check what SDK changed
2. Create tag
3. Generate changelog and put it into release description

## Contributing to App Services SDK

## Generating an API client locally

You may want to generate an SDK locally. This can be done in two steps.

1. Fetch the OpenAPI specification file.

```shell
# Set the arguments
remote_openapi_url="https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml"

# List of client IDs can be found in `.config/api-client-metadata.json`
client_id="kafka-mgmt/v1"

./scripts/fetch_api.sh $remote_openapi_url
```

2. Generate the OpenAPI client

```shell
./scripts/generate_go.sh $client_id $remote_openapi_url
```

## Adding a new service

_Follow these steps if you wish to add an API client to the SDK for your API service._

> The current workflow for API client generation is limited to the services which are part of the **bf2fc6cc711aee1a0c2a** GitHub organization. If you wish to onboard a service which is in a different GitHub organization please [open an issue](https://github.com/redhat-developer/app-services-sdk-go/issues/new) so that we can discuss options.

#### 1. Adding an access token to your project

In order to send a [`repository_dispatch`](https://docs.github.com/en/actions/reference/events-that-trigger-workflows#repository_dispatch) event to the SDK repositories, you need to use a `public_repo` scoped personal access token. This token has already been created - ask your organization administrator to add this as a repository secret called `APP_SERVICES_CI_TOKEN`.

#### 2. Define your GitHub Action

When a change to the OpenAPI document is merged to the main branch, an event will be sent to each of the SDK repositories, where an API client will be generated using the latest OpenAPI specification document.

Please use the following template to create a GitHub Action in `.github/workflows` of your service's repository.

```yaml
name: Dispatch OpenAPI File Change event
on:
  push:
    branches: 
      - main # use whatever your primary Git branch is called.
    paths:
      - <path/to/openapi/your-service-name.yaml>

jobs:
  dispatch:
    if: github.repository == 'bf2fc6cc711aee1a0c2a/<your-service-name>'
    env:
    strategy:
      matrix:
        repo: ["redhat-developer/app-services-sdk-go", "redhat-developer/app-services-sdk-js", "redhat-developer/app-services-sdk-java"]
    runs-on: ubuntu-latest
    steps:
      - name: Repository Dispatch
        uses: peter-evans/repository-dispatch@v1
        with:
          token: ${{ secrets.APP_SERVICES_CI_TOKEN }}
          repository: ${{ matrix.repo }}
          event-type: openapi-spec-change
          client-payload: '{ "id": "<unique-service-id>/<v1alpha>", "download_url": "<openai_file_download_url>"}'
```

See [a working example here](https://github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/blob/main/.github/workflows/openapi_update.yaml).

- `on.push.branches`: Replace this with the primary branch of your Git repository.
- `on.push.paths`: Replace this with the relative path to the OpenAPI specification document. This ensures that _only_ changes to this file will trigger the event.
- `jobs.dispatch.if`: Replace `<your-service-name>` with the name of the GitHub repository. This condition ensures that changes on forks do not trigger this event.
- `jobs.steps[0].with.client-payload`:
    - Replace `<unique-service-id>` with a clear, unique identifer for your service. Previous examples include `kafka-mgmt` and `srs-mgmt`.
    - Replace `<v1alpha>` with the major API version of your service.
    - Replace `<openai_file_download_url>` with the direct URL to the raw OpenAPI document. Example: [https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml](https://raw.githubusercontent.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/main/openapi/kas-fleet-manager.yaml)

#### 3. Define API client metadata

The Go SDK relies on additional metadata for each API in order to generate code in a "Go" way. Add your API client config to [.config/api-client-metadata.json](.config/api-client-metadata.json) like so:

```json
"foo-mgmt/v1beta": {
  "apiGroup": "foomgmt",
  "apiVersion": "v1beta",
  "release_level": "beta",
}
```

The `"foo-mgmt/v1beta"` key should match the ID from the received client payload which we defined in step 2. The generated client is output to `foomgmt/v1beta/client`.

#### 4. Define custom API constructor

Every API client should have a custom constructor defined one level above the generated client code in `foomgmt/v1beta`.

Copy [kafkamgmt/apiv1/api_client.go](kafkamgmt/apiv1/api_client.go) into the `foomgmt/v1beta` folder and fix the imports. If you find that that the config is too restrictive for the requirements, please update the APIConfig in [internal/api_config.go](internal/api_config.go) and start a discussion on a pull request, or open an issue.

#### 5. Add custom code (optional)

If your SDK requires custom code, define them in your desired structure in `foomgmt/v1beta`. Please do not add custom code to `foomgmt/v1beta/client` as this could be overwritten.

## Linting

Run `make lint` to perform a lint check of the Go code.