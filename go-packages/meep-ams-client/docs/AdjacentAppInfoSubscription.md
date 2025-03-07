# AdjacentAppInfoSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [***AdjacentAppInfoSubscriptionLinks**](AdjacentAppInfoSubscriptionLinks.md) |  | [optional] [default to null]
**CallbackReference** | **string** | URI selected by the service consumer to receive notifications on the subscribed Application Mobility Service. This shall be included both in the request and in response. | [default to null]
**RequestTestNotification** | **bool** | Shall be set to TRUE by the service consumer to request a test notification via HTTP on the callbackReference URI, specified in ETSI GS MEC 009, as described in clause 6.12a. | [optional] [default to null]
**WebsockNotifConfig** | [***WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] [default to null]
**ExpiryDeadline** | [***TimeStamp**](TimeStamp.md) |  | [optional] [default to null]
**FilterCriteria** | [***AdjacentAppInfoSubscriptionFilterCriteria**](AdjacentAppInfoSubscriptionFilterCriteria.md) |  | [default to null]
**SubscriptionType** | [***SubscriptionType**](SubscriptionType.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


