# Go API client for connectormgmtclient

Connector Service Fleet Manager is a Rest API to manage connectors.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./connectormgmtclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.openshift.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConnectorClustersApi* | [**CreateConnectorCluster**](docs/ConnectorClustersApi.md#createconnectorcluster) | **Post** /api/connector_mgmt/v1/kafka_connector_clusters | Create a new connector cluster
*ConnectorClustersApi* | [**DeleteConnectorCluster**](docs/ConnectorClustersApi.md#deleteconnectorcluster) | **Delete** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Delete a connector cluster
*ConnectorClustersApi* | [**GetConnectorCluster**](docs/ConnectorClustersApi.md#getconnectorcluster) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id} | Get a connector cluster
*ConnectorClustersApi* | [**GetConnectorClusterAddonParameters**](docs/ConnectorClustersApi.md#getconnectorclusteraddonparameters) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters/{connector_cluster_id}/addon_parameters | Get a connector cluster&#39;s addon parameters
*ConnectorClustersApi* | [**ListConnectorClusters**](docs/ConnectorClustersApi.md#listconnectorclusters) | **Get** /api/connector_mgmt/v1/kafka_connector_clusters | Returns a list of connector clusters
*ConnectorTypesApi* | [**GetConnectorTypeByID**](docs/ConnectorTypesApi.md#getconnectortypebyid) | **Get** /api/connector_mgmt/v1/kafka_connector_types/{connector_type_id} | Get a connector type by id
*ConnectorTypesApi* | [**ListConnectorTypes**](docs/ConnectorTypesApi.md#listconnectortypes) | **Get** /api/connector_mgmt/v1/kafka_connector_types | Returns a list of connector types
*ConnectorsApi* | [**CreateConnector**](docs/ConnectorsApi.md#createconnector) | **Post** /api/connector_mgmt/v1/kafka_connectors | Create a new connector
*ConnectorsApi* | [**DeleteConnector**](docs/ConnectorsApi.md#deleteconnector) | **Delete** /api/connector_mgmt/v1/kafkas_connectors/{id} | Delete a connector
*ConnectorsApi* | [**GetConnector**](docs/ConnectorsApi.md#getconnector) | **Get** /api/connector_mgmt/v1/kafkas_connectors/{id} | Get a connector
*ConnectorsApi* | [**ListConnectors**](docs/ConnectorsApi.md#listconnectors) | **Get** /api/connector_mgmt/v1/kafka_connectors | Returns a list of connector types
*ConnectorsApi* | [**PatchConnector**](docs/ConnectorsApi.md#patchconnector) | **Patch** /api/connector_mgmt/v1/kafkas_connectors/{id} | patch a connector


## Documentation For Models

 - [AddonClusterTarget](docs/AddonClusterTarget.md)
 - [AddonParameter](docs/AddonParameter.md)
 - [CloudProviderClusterTarget](docs/CloudProviderClusterTarget.md)
 - [ClusterTarget](docs/ClusterTarget.md)
 - [Connector](docs/Connector.md)
 - [ConnectorAllOf](docs/ConnectorAllOf.md)
 - [ConnectorAllOfMetadata](docs/ConnectorAllOfMetadata.md)
 - [ConnectorCluster](docs/ConnectorCluster.md)
 - [ConnectorClusterAllOf](docs/ConnectorClusterAllOf.md)
 - [ConnectorClusterAllOfMetadata](docs/ConnectorClusterAllOfMetadata.md)
 - [ConnectorClusterList](docs/ConnectorClusterList.md)
 - [ConnectorClusterListAllOf](docs/ConnectorClusterListAllOf.md)
 - [ConnectorList](docs/ConnectorList.md)
 - [ConnectorListAllOf](docs/ConnectorListAllOf.md)
 - [ConnectorType](docs/ConnectorType.md)
 - [ConnectorTypeAllOf](docs/ConnectorTypeAllOf.md)
 - [ConnectorTypeList](docs/ConnectorTypeList.md)
 - [ConnectorTypeListAllOf](docs/ConnectorTypeListAllOf.md)
 - [Error](docs/Error.md)
 - [ErrorAllOf](docs/ErrorAllOf.md)
 - [KafkaConnectionSettings](docs/KafkaConnectionSettings.md)
 - [List](docs/List.md)
 - [ObjectReference](docs/ObjectReference.md)


## Documentation For Authorization



### Bearer

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARERTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


