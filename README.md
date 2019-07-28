#### elliepotato-api

### About
This a public / personal API server to do random things.

As to what it will do yet, I'm not sure, but things will be added overtime.
It will be hosted on elliepotato.de

## Making a request
| Field         | Description                                     | Required?                      |
| :-------:     | :---------------------------------------------: | :----------------------------: |
| ~~session~~   | ~~Session key, for more than 1 request sessions.~~  | ~~Optional, depends on end point~~ |
| endpoint      | The endpoint the request is intended for.       | Yes                            |
| method        | The method for the endpoint to read.            | Yes                            |
| payload       | Payload information for the endpoint to process.| No                             |

If there is an issue decoding the request, an ambigious 400 error will be returned.
If a required field is missing, an ambigious error 400 error will be returned.

### End-points:
*There are limited right now and not at all useful to anyone but here they are.*

| Field     | Description                                          | Requires auth? |
| :-------: | :--------------------------------------------------: |  :------------: |
| identify  | ~~For creating and invalidating self-session.~~ **No longer in use until further notice.**          | No             |
| plugin    | For getting up to date version stats about a plugin. | No             |
| service    | Subscribing and get statuses about internal services | Yes             |


If the end-point doesn't exist, an ambigious 400 error will be returned.
If you try to access an endpoint requring authentication without a valid session (will be further specified later when necessary), will return a 403 error.

If a method is given for an end-point which is not recognized an ambigious 400 error will be returned.
A full list of API responses can be found [Here](https://github.com/literallyEllie/elliepotato-api/blob/master/src/main/api_response.go)

## Plugin methods

### GetVersion
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |
| plugin-id | ID of plugin to get. | Yes       |

### Responses:
* If "plugin-id" is not a registered plugin, will return 404 error.
* A 200 OK response will be returned with no additional information.

## Service methods

### Subscribe
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |
| login_key | Login key to access the system. | Yes       |
| to | The service to subscribe to. | Yes       |

### Responses:
* If "to" isn't a valid service, will return 404 error.
* A 200 OK response wil be returned with no additional information.

### GetStatus
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |
| login_key | Login key to access the system. | Yes       |
| to | The service to get the status of. | Yes       |

### Responses:
* If "to" isn't a valid service, will return 404 error.
* If the service is NOT active, will return 200 with "not_active"
* A 200 OK response will be returned allong with a JSON escaped string of [ep_service#ActiveService](https://github.com/literallyEllie/elliepotato-api/blob/master/src/main/ep_service.go#L17) 

### Status
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |
| id | ID of service which will be set. | Yes       |
| access_key | Access key to manage the ID | Yes       |
| type | Either "alive" or "exit". Alive is recognized as a string and exit is recognized as the service going offline. | Yes       |

### Responses:
* If "id" isn't a valid service, will return a 404 error.
* If the access_key cannot manage the ID, will return a 400 error.
* A 200 OK response will be returned with no additional information.

___
## --TRASH BIN--

## Identify methods

### NewSession
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |

Will send back your session ID with a 200 OK code.

### InvalidateSession
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |

### Responses:
* Will invalidate the "session" bundled with your request. 
* An 200 OK response will be returned regarldess of the outcome.


*More will be added when new things happen!!*
