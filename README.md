#### elliepotato-api

### About
This a public / personal API server to do random things.

As to what it will do yet, I'm not sure, but things will be added overtime.
It will be hosted on elliepotato.de

## Making a request
| Field     | Description                                     | Required?                      |
| :-------: | :---------------------------------------------: | :----------------------------: |
| session   | Session key, for more than 1 request sessions.  | Optional, depends on end point |
| endpoint  | The endpoint the request is intended for.       | Yes                            |
| method    | The method for the endpoint to read.            | Yes                            |
| payload   | Payload information for the endpoint to process.| No                             |

If there is an issue decoding the request, an ambigious 400 error will be returned.
If a required field is missing, an ambigious error 400 error will be returned.

### End-points:
*There are limited right now and not at all useful to anyone but here they are.*

| Field     | Description                                          | Requires auth? |
| :-------: | :--------------------------------------------------: | :------------: |
| identify  | For creating and invalidating self-session.          | No             |
| plugin    | For getting up to date version stats about a plugin. | No             |

If the end-point doesn't exist, an ambigious 400 error will be returned.
If you try to access an endpoint requring authentication without a valid session (will be further specified later when necessary), will return a 403 error.

If a method is given for an end-point which is not recognized an ambigious 400 error will be returned.

## Identify methods

# NewSession
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |

Will send back your session ID with a 200 OK code.

# InvalidateSession
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |

Will invalidate the "session" bundled with your request. 
An 200 OK request will be returned regarldess of the outcome.

## Plugin methods

# GetVersion
| Parameter | Description          | Required? |
| :-------: | :------------------: | :-------: |
| plugin-id | ID of plugin to get. | Yes       |

If "plugin-id" is not a registered plugin, will return 404 error.

*More will be added when new things happen!!*
