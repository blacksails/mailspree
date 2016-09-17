# API Reference

This is the api specification of the mailspree http service. All data must be
sent as json objects.

## Types

### Email
The Email type represents an email

| Param   | Type   | Required |
| ------- | ------ | -------- |
| name    | string | false    |
| address | string | true     |

## Endpoints

### Session

The session endpoint is used to get an authentication token.

#### /session POST

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

### Send

The send endpoint is used for sending email.

#### /send POST

##### request
| Param   | Type    | Required |
| ------- | ------- | -------- |
| from    | Email   | true     |
| to      | []Email | true     |
| subject | string  | true     |
| body    | string  | true     |

##### response
Responds with 200 and an empty body

##### request example
```json
{
  "from": {
    "name": "Santa Claus",
    "address": "santa@christmas.com"
  },
  "to": [
    {
      "name": "Happy Elf",
      "address": "happy@christmas.com"
    },
    {
      "name": "Sad Elf",
      "address": "sad@christmas.com"
    }
  ],
  "subject": "Chirstmas is delayed",
  "body": "Unfortunately christmas has been delayed"
}
```
