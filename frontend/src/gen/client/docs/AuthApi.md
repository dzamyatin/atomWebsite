# AuthProto.AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**authLogin**](AuthApi.md#authLogin) | **POST** /login | 
[**authRegister**](AuthApi.md#authRegister) | **POST** /register | 



## authLogin

> AuthLoginResponse authLogin(body)



### Example

```javascript
import AuthProto from 'auth_proto';

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## authRegister

> AuthRegisterResponse authRegister(body)



### Example

```javascript
import AuthProto from 'auth_proto';

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

