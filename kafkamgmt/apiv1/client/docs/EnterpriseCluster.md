# EnterpriseCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | ocm cluster id of the registered Enterprise cluster | [optional] 
**Status** | Pointer to **string** | status of registered Enterprise cluster | [optional] 
**FleetshardParameters** | Pointer to [**[]FleetshardParameter**](FleetshardParameter.md) |  | [optional] 

## Methods

### NewEnterpriseCluster

`func NewEnterpriseCluster() *EnterpriseCluster`

NewEnterpriseCluster instantiates a new EnterpriseCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseClusterWithDefaults

`func NewEnterpriseClusterWithDefaults() *EnterpriseCluster`

NewEnterpriseClusterWithDefaults instantiates a new EnterpriseCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *EnterpriseCluster) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnterpriseCluster) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnterpriseCluster) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *EnterpriseCluster) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetStatus

`func (o *EnterpriseCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnterpriseCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnterpriseCluster) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnterpriseCluster) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFleetshardParameters

`func (o *EnterpriseCluster) GetFleetshardParameters() []FleetshardParameter`

GetFleetshardParameters returns the FleetshardParameters field if non-nil, zero value otherwise.

### GetFleetshardParametersOk

`func (o *EnterpriseCluster) GetFleetshardParametersOk() (*[]FleetshardParameter, bool)`

GetFleetshardParametersOk returns a tuple with the FleetshardParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetshardParameters

`func (o *EnterpriseCluster) SetFleetshardParameters(v []FleetshardParameter)`

SetFleetshardParameters sets FleetshardParameters field to given value.

### HasFleetshardParameters

`func (o *EnterpriseCluster) HasFleetshardParameters() bool`

HasFleetshardParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

