# cache-memory
This is a project for storing a key-value store and accessing it via API. The data is also stored persistently to the given file path as JSON within given interval.

It provides thread-safe ```map[string]string``` storage.

# Dependencies
It is developed by using go standard packages.
### Used packages
- fmt
- log
- net/http
- encoding/json
- os
- io/ioutil
- sync
- time
- strconv

# Installation
You must have golang installed on your computer to run the project.

Clone the project
```shell
git clone git@github.com:canberksinangil/cache-memory.git
```

You can set the following environmental variables. If you do not so, default values will be used. You can see them in the ```config.go```
```
DEFAULT_FILE_PATH
DEFAULT_SAVING_FREQUENCY
DEFAULT_PORT
```

Navigate to the project directory in your terminal and run
```shell
go run main.go
```

# API DOC

| Endoints | Methods |
| :------------ | :------------: |
| /healtz | GET |
| /cache | GET  |
| /cache | POST |
| /cache | DELETE |
| /flush | DELETE |

---
### /healtz

#### GET
##### Summary

Check if API up and running

##### Description

Controls the server and returns 200 if API runs properly in the specified port. Can be used for health check.

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Status OK |
| 405 | Method Not Allowed  |
| 500 | Internal Server Error |

##### Example
```
curl --location --request GET 'localhost:3333/healthz'
```

### /cache

#### GET
##### Summary

Get the value of given key.

##### Description

Checks the cache and returns the value of the given key if exist. If it is not exist it returns the error. `Key` the is required parameter.

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Status OK |
| 400 | Bad Request |
| 405 | Method Not Allowed  |
| 500 | Internal Server Error |

##### Example
```
curl --location --request GET 'localhost:3333/cache?key=a'
```

##### Response Examples
```
200

{
    "value": "aa"
}
```

```
200

{
    "error": "The key '1' could not be found."
}
```

```
400

{
    "error": "The 'key' is required."
}
```

### /cache

#### POST
##### Summary

Set key-value.

##### Description

Sets the given value for the given key. It accepts only strings. 'Key' and 'Value' are the required parameters. If the key exist, it overrides the value.

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Status OK |
| 400 | Bad Request |
| 405 | Method Not Allowed  |
| 500 | Internal Server Error |

##### Example
```
curl --location --request POST 'localhost:3333/cache' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "test_key",
    "value": "test_value"
}'
```

##### Response Examples

```
400

{
    "error": "The 'key' is required."
}
```

```
400

{
    "error": "The 'value' is required."
}
```

```
400

{
    "error": "json: cannot unmarshal number into Go struct field setRequest.key of type string"
}
```

### /cache

#### DELETE
##### Summary

Delete the given key

##### Description

Checks the cache for the given key. If the key exist, it deletes the key.

##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Status OK |
| 405 | Method Not Allowed  |
| 500 | Internal Server Error |

##### Example
```
curl --location --request DELETE 'localhost:3333/cache' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "test_key"
}'
```

### /flush

#### DELETE
##### Summary

Flush the cache

##### Description

Flush basically creates a new cache and tight after it truncates the JSON file.
##### Responses

| Code | Description |
| ---- | ----------- |
| 200 | Status OK |
| 405 | Method Not Allowed  |
| 500 | Internal Server Error |

##### Example
```
curl --location --request DELETE 'localhost:3333/flush'
```