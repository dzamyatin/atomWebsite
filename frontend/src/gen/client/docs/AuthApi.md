# AuthProto.AuthApi

All URIs are relative to *http://localhost:8503*

Method | HTTP request | Description
------------- | ------------- | -------------
[**authChangePassword**](AuthApi.md#authChangePassword) | **POST** /change-password | 
[**authConfirmEmail**](AuthApi.md#authConfirmEmail) | **POST** /confirm-email | 
[**authConfirmPhone**](AuthApi.md#authConfirmPhone) | **POST** /confirm-phone | 
[**authLogin**](AuthApi.md#authLogin) | **POST** /login | 
[**authRegister**](AuthApi.md#authRegister) | **POST** /register | 
[**authRememberPassword**](AuthApi.md#authRememberPassword) | **POST** /remember-password | 
[**authSendEmailConfirmation**](AuthApi.md#authSendEmailConfirmation) | **POST** /send-email-confirmation | 
[**authSendPhoneConfirmation**](AuthApi.md#authSendPhoneConfirmation) | **POST** /send-phone-confirmation | 



## authChangePassword

> Object authChangePassword(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthChangePasswordRequest(); // AuthChangePasswordRequest | 
apiInstance.authChangePassword(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthChangePasswordRequest**](AuthChangePasswordRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authConfirmEmail

> Object authConfirmEmail(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthConfirmEmailRequest(); // AuthConfirmEmailRequest | 
apiInstance.authConfirmEmail(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthConfirmEmailRequest**](AuthConfirmEmailRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authConfirmPhone

> Object authConfirmPhone(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthConfirmPhoneRequest(); // AuthConfirmPhoneRequest | 
apiInstance.authConfirmPhone(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthConfirmPhoneRequest**](AuthConfirmPhoneRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authLogin

> AuthLoginResponse authLogin(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthLoginRequest(); // AuthLoginRequest | 
apiInstance.authLogin(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthLoginRequest**](AuthLoginRequest.md)|  | 

### Return type

[**AuthLoginResponse**](AuthLoginResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authRegister

> AuthRegisterResponse authRegister(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthRegisterRequest(); // AuthRegisterRequest | 
apiInstance.authRegister(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthRegisterRequest**](AuthRegisterRequest.md)|  | 

### Return type

[**AuthRegisterResponse**](AuthRegisterResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authRememberPassword

> Object authRememberPassword(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthRememberPasswordRequest(); // AuthRememberPasswordRequest | 
apiInstance.authRememberPassword(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthRememberPasswordRequest**](AuthRememberPasswordRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authSendEmailConfirmation

> Object authSendEmailConfirmation(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthSendEmailConfirmationRequest(); // AuthSendEmailConfirmationRequest | 
apiInstance.authSendEmailConfirmation(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthSendEmailConfirmationRequest**](AuthSendEmailConfirmationRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authSendPhoneConfirmation

> Object authSendPhoneConfirmation(body)



### Example

```javascript
import AuthProto from 'auth_proto';
let defaultClient = AuthProto.ApiClient.instance;
// Configure OAuth2 access token for authorization: Bearer
let Bearer = defaultClient.authentications['Bearer'];
Bearer.accessToken = 'YOUR ACCESS TOKEN';

let apiInstance = new AuthProto.AuthApi();
let body = new AuthProto.AuthSendPhoneConfirmationRequest(); // AuthSendPhoneConfirmationRequest | 
apiInstance.authSendPhoneConfirmation(body, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthSendPhoneConfirmationRequest**](AuthSendPhoneConfirmationRequest.md)|  | 

### Return type

**Object**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

