# microservice-user

Microservice to manage users.

## User endpoints

`GET` `/users` Returns json data about every user.

`GET` `/users/{id}` Returns json data about a specific user. `id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/users` Add new user with specific data.</br>
__Data Params__
```json
{
  "id":          "string",
  "username":    "string, required",
  "email":       "string, required",
  "dateofbirth": "string, required",
  "firstname":   "string",
  "lastname":    "string",
  "gender":      "string",
  "address":     "string",
  "bio":         "string"
}
```

`PUT` `/users` Update user data. </br>
__Data Params__
```json
{
  "id":           "string, required",
  "username":     "string",
  "email":        "string",
  "dateofbirth":  "string",
  "firstname":    "string",
  "lastname":     "string",
  "gender":       "string",
  "address":      "string",
  "bio":          "string",
  "achievements": ["string"]
}
```

`DELETE` `/users/{id}` Delete user. `id=[string]`
