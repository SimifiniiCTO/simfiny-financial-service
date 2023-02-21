# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [simfiny/proto/service/api/v1/common/common.proto](#simfiny_proto_service_api_v1_common_common-proto)
- [simfiny/proto/service/api/v1/response/response.proto](#simfiny_proto_service_api_v1_response_response-proto)
    - [CreateAccountResponse](#response-CreateAccountResponse)
    - [EmptyResponse](#response-EmptyResponse)
    - [PlaidLinkTokenExchangeResponse](#response-PlaidLinkTokenExchangeResponse)
    - [PlaidLinkUpdateResponse](#response-PlaidLinkUpdateResponse)
    - [PlaidLinktTokenResponse](#response-PlaidLinktTokenResponse)
    - [ServiceHealthResponse](#response-ServiceHealthResponse)
    - [ServiceReadyResponse](#response-ServiceReadyResponse)
  
- [simfiny/proto/service/api/v1/request/request.proto](#simfiny_proto_service_api_v1_request_request-proto)
    - [CreateAccountRequest](#request-CreateAccountRequest)
    - [DeleteAccountRequest](#request-DeleteAccountRequest)
    - [EmptyRequest](#request-EmptyRequest)
    - [PlaidLinkTokenExchangeRequest](#request-PlaidLinkTokenExchangeRequest)
    - [PlaidLinkTokenRequest](#request-PlaidLinkTokenRequest)
    - [PlaidLinkUpdateRequest](#request-PlaidLinkUpdateRequest)
    - [PlaidManualLinkResyncRequest](#request-PlaidManualLinkResyncRequest)
  
- [simfiny/proto/service/api/v1/service.proto](#simfiny_proto_service_api_v1_service-proto)
    - [FinancialService](#endpoint-service-FinancialService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="simfiny_proto_service_api_v1_common_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## simfiny/proto/service/api/v1/common/common.proto


 

 

 

 



<a name="simfiny_proto_service_api_v1_response_response-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## simfiny/proto/service/api/v1/response/response.proto



<a name="response-CreateAccountResponse"></a>

### CreateAccountResponse
CreateAccountResponse encompasses the response from a create account call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| AccountID | [uint64](#uint64) |  |  |






<a name="response-EmptyResponse"></a>

### EmptyResponse
EmptyResponse is a response type used in scenarios when a response is not
expected






<a name="response-PlaidLinkTokenExchangeResponse"></a>

### PlaidLinkTokenExchangeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linkId | [int64](#int64) |  | The ID of the newly created link for the Plaid connection. The Link ID is not a unique identifier for Plaid, it is simfiny&#39;s internal unique identifier. |






<a name="response-PlaidLinkUpdateResponse"></a>

### PlaidLinkUpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linkToken | [string](#string) |  | If the link is able to be put into update mode, then a Link Token is returned to the client. This can then be used with the Plaid Link library to allow the user to update their account selection or re - authenticate their link with their bank. |






<a name="response-PlaidLinktTokenResponse"></a>

### PlaidLinktTokenResponse
The link token created from Plaid, this can be used with the Plaid SDK to
authenticate a user&#39;s bank account.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| LinkToken | [string](#string) |  |  |






<a name="response-ServiceHealthResponse"></a>

### ServiceHealthResponse
ServiceHealthResponse: Represent the object returned as a response to the
service health api invocation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Healthy | [bool](#bool) |  |  |






<a name="response-ServiceReadyResponse"></a>

### ServiceReadyResponse
ServiceReadyResponse: Represent the object returned as a response to the
service readyness api invocation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Ready | [bool](#bool) |  |  |





 

 

 

 



<a name="simfiny_proto_service_api_v1_request_request-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## simfiny/proto/service/api/v1/request/request.proto



<a name="request-CreateAccountRequest"></a>

### CreateAccountRequest
CreateAccountRequest witholds parameters required to properly create an
account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| globalUserID | [uint64](#uint64) |  |  |
| email | [string](#string) |  |  |






<a name="request-DeleteAccountRequest"></a>

### DeleteAccountRequest
DeleteAccountRequest witholds the parameters required to delete an account


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| globalUserID | [uint64](#uint64) |  |  |






<a name="request-EmptyRequest"></a>

### EmptyRequest
EmptyRequest: Represents a request object invoked against the social service
that is empty






<a name="request-PlaidLinkTokenExchangeRequest"></a>

### PlaidLinkTokenExchangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| publicToken | [string](#string) |  | The public_token you received from the Plaid Link flow via the SDK. |
| institutionId | [string](#string) |  | The Plaid institution_id, the API will not return an error if this is not provided (TODO). But it is required in order for the API to function properly. |
| institutionName | [string](#string) |  | The name of the institution as provided by the Plaid SDK. This is used as the initial name for the link until the user decides to rename it. |
| accountIds | [string](#string) | repeated | An array of Plaid bank account IDs, this should be the list of IDs that the user selected when they were going through the link process. |






<a name="request-PlaidLinkTokenRequest"></a>

### PlaidLinkTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| useCache | [bool](#bool) |  | If true then the endpoint will not try to create another link token if one has already been created for the current user and has not been completed. This is used to reduce the number of API calls made to Plaid. It is recommended to use true for this. The endpoint will never return a Link token that has already been used. |
| globalUserID | [uint64](#uint64) |  | The global user on whose behalf we are making the request |






<a name="request-PlaidLinkUpdateRequest"></a>

### PlaidLinkUpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linkId | [int64](#int64) |  | A link ID must be provided in order to put that link into update mode. The link must also be a Plaid link, a manual link will result in a bad request. |
| updateAccountSelection | [bool](#bool) |  | This parameter is used to specify whether you want to put this link into an update mode that will allow you to add/remove accounts that are visible to monetr. This will change the Plaid Link dialog behavior slightly when it is presented to the client. |






<a name="request-PlaidManualLinkResyncRequest"></a>

### PlaidManualLinkResyncRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| linkId | [int64](#int64) |  | The link ID that you want to manually resync. This must be a Plaid link that is in a status of Setup or Error. Other link statuses will result in a bad request. |





 

 

 

 



<a name="simfiny_proto_service_api_v1_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## simfiny/proto/service/api/v1/service.proto


 

 

 


<a name="endpoint-service-FinancialService"></a>

### FinancialService
FinancialService: Represents the financial service of the simifinii platform.

The service serves as the entrypoint for all financial interactions on the
simfiny platform

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ServiceHealth | [.request.EmptyRequest](#request-EmptyRequest) | [.response.ServiceHealthResponse](#response-ServiceHealthResponse) | ServiceHealth: Used to get the service health status of a service |
| ServiceReady | [.request.EmptyRequest](#request-EmptyRequest) | [.response.ServiceReadyResponse](#response-ServiceReadyResponse) | ServiceReady: Used to deduce wether service is ready to serve request |
| CreateAccount | [.request.CreateAccountRequest](#request-CreateAccountRequest) | [.response.CreateAccountResponse](#response-CreateAccountResponse) | CreateAccount: Used to create an account record |
| DeleteAccount | [.request.DeleteAccountRequest](#request-DeleteAccountRequest) | [.response.EmptyResponse](#response-EmptyResponse) | DeleteAccount: Used to perform a delete account request |
| GetNewLinkToken | [.request.PlaidLinkTokenRequest](#request-PlaidLinkTokenRequest) | [.response.PlaidLinktTokenResponse](#response-PlaidLinktTokenResponse) | In order to connect a user&#39;s account with Plaid and thus their bank account, you need to create a link token. This token is then used by the Plaid SDK in the UI in order to provide an authentication flow for that user and their bank |
| PlaidTokenCallback | [.request.PlaidLinkTokenExchangeRequest](#request-PlaidLinkTokenExchangeRequest) | [.response.PlaidLinkTokenExchangeResponse](#response-PlaidLinkTokenExchangeResponse) | Once the user has completed authenticating their bank account using the Plaid SDK, you will be given a public_token. This new token must be exchanged with Plaid&#39;s API in order to allow simfiny to access the bank data. The bank account should not be considered linked until this token is successfully exchanged. |
| UpdatePlaidLink | [.request.PlaidLinkUpdateRequest](#request-PlaidLinkUpdateRequest) | [.response.PlaidLinkUpdateResponse](#response-PlaidLinkUpdateResponse) | Plaid links can be updated after they have been established. This can be used as a method of re-authenticating a link if it ends up in an error state (though sometimes a link can end up in an error state without indicating [2]). This can also be used as a way to add additional accounts to a link if those accounts were not originally granted access when the link was created. |
| ManuallyResyncPlaidLink | [.request.PlaidManualLinkResyncRequest](#request-PlaidManualLinkResyncRequest) | [.response.EmptyResponse](#response-EmptyResponse) | Sometimes you might need to manually trigger a resync with Plaid; this can happen if there were issues where a webhook was not properly received. By triggering a manual resync, transactions for the last 14 days and balances for all bank accounts within the specified link will be checked.

Links can be manually synced once every 30 minutes, an error will be returned if you try to resync a link too quickly.

NOTE - This will not send a &#34;sync&#34; request to Plaid. This will only retrieve data already available via Plaid&#39;s API and update simfiny&#39;s data accordingly. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

