# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/CountryFriedCoders/OpenTerps/0.0.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTerpene**](TerpeneApi.md#AddTerpene) | **Post** /terpene | Add a new terpene to the api
[**DeleteTerpene**](TerpeneApi.md#DeleteTerpene) | **Delete** /terpene/{terpeneId} | Deletes a terpene
[**FindTerpenesByEffect**](TerpeneApi.md#FindTerpenesByEffect) | **Get** /terpene/findByEffect | Finds Terpenes by effect
[**GetTerpeneById**](TerpeneApi.md#GetTerpeneById) | **Get** /terpene/{terpeneId} | Find terpene by ID
[**UpdateTerpene**](TerpeneApi.md#UpdateTerpene) | **Put** /terpene | Update an existing terpene
[**UpdateTerpeneByID**](TerpeneApi.md#UpdateTerpeneByID) | **Post** /terpene/{terpeneId} | Updates a terpene in the api

# **AddTerpene**
> AddTerpene(ctx, body)
Add a new terpene to the api

Add new terpene to the api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Terpene**](Terpene.md)| Terpene object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTerpene**
> DeleteTerpene(ctx, apiKey, terpeneId)
Deletes a terpene

Deletes a terpene from the API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**|  | 
  **terpeneId** | **int64**| Terpene id to delete | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FindTerpenesByEffect**
> []Terpene FindTerpenesByEffect(ctx, effect)
Finds Terpenes by effect

Multiple status values can be provided with comma separated strings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **effect** | [**[]string**](string.md)| Effect values that need to be considered for filter | 

### Return type

[**[]Terpene**](Terpene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTerpeneById**
> Terpene GetTerpeneById(ctx, terpeneId)
Find terpene by ID

Returns a single terpene

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terpeneId** | **int64**| ID of terpene to return | 

### Return type

[**Terpene**](Terpene.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/xml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTerpene**
> UpdateTerpene(ctx, body)
Update an existing terpene

Update terpene in the api

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Terpene**](Terpene.md)| Terpene object | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTerpeneByID**
> UpdateTerpeneByID(ctx, terpeneId, optional)
Updates a terpene in the api

Update a single terpene in the API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terpeneId** | **int64**| ID of terpene that needs to be updated | 
 **optional** | ***TerpeneApiUpdateTerpeneByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TerpeneApiUpdateTerpeneByIDOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **effects** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

