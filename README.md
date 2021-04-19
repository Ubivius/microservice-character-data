# character-data-microservice
Character data microservice for our online game framework. 

## Character data endpoints

`GET` `/characters` Returns json data about every character.

`GET` `/characters/{id}` Returns json data about a specific character. `id=[string]`

`GET` `/characters/user/{user_id}` Returns the list of characters of a user. `user_id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/characters` Add new character with specific data.

__Data Params__
```json
{
  "user_id": "string, required",
  "name":    "string, required",
}
```

`PUT` `/characters` Update character data </br>
__Data Params__
```json
{
  "id":      "string, required",
  "user_id": "string",
  "name":    "string",
}
```

`DELETE` `/characters/{id}` Delete a character.  `id=[string]`
