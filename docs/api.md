# API Reference

This is the api specification of the mailspree http service. All data must be
sent as json objects.

## Endpoints

### Session

This endpoint is used to get an authentication token.

#### /send POST

##### request
| Param    | Type   | Required |
| -------- | ------ | -------- |
| username | string | true     |
| password | string | true     |

##### response
| Param    | Type   | 
| -------- | ------ |
| username | string |
| token    | string |

##### request example
```json
{
  "username": "admin",
  "password": "supersecret"
}
```

##### response example
```json
{
  "username": "admin",
  "token": "letssaythisisatoken"
}
```
