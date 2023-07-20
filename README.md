# Halo Lab Test Task

## Task

The task is to implement the backend service which will continuously generate the (fake) data; and will provide an API to access this data. The service should be implemented using Go language (you are free to choose the web framework for the task). The data should be stored in PostgreSQL database. It is expected that the database for running the project locally will be described in a Docker Compose file — so the project could be started locally via the npm / docker commands only. It is also expected that API will have a Swagger spec.

Each sensor has following properties:

- codename
- 3D-coordinates (x, y, z)
- data output rate (in seconds)

Sensors are organised into the groups; each group has a name (we will use the greek letter names for these: alpha, beta, gamma, etc.). Sensor’s codename consists of the group name and the sensor’s index (integer number, monotonically increasing within a group), for example: gamma 3.
Sensors provide the following data:

- temperature (Celsius, floating point number)
- transparency (%, integer number)
- the list of fish species detected near the sensor since the previous measurement; for each fish specie sensor returns the name (string) & the count (integer), for example:
  - Atlantic Cod: 12
  - Sailfish: 4

## Prerequisites

Before running the project, configure environment variables in .env file in root directory.
Example of .env file:

```
# Example .env file

# Postgres configuration
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=qwerty
POSTGRES_DB=test_db
POSTGRES_SSLMODE=disable

# Redis configuration
REDIS_PORT=6379
REDIS_PASSWORD=qwerty
REDIS_ADDR="cache:${REDIS_PORT}"
REDIS_DB=0

# App configuration
PORT=8080
```

## Installation

To build and run the project, you will need Docker Compose installed on your system.

To build the project run:

`make build && make run`

## Usage

To start only API run:

`make api`

To start only PostgreSQL run:

`make postgres`

To start only Redis run:

`make cache`

To run migrations run:

`make migrate`

## API Endpoints

### Get Average Transparency in Group

Get the current average transparency inside the specified group.

**URL**: `/group/{groupName}/transparency/average`

**Method**: `GET`

**Path Parameters**:

- `groupName` (string, required): Name of the group for which to retrieve the average transparency.

**Response**:

- `200 OK`: Returns the average transparency value as a float64.
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---

### Get Average Temperature in Group

Get the current average temperature inside the specified group.

**URL**: `/group/{groupName}/temperature/average`

**Method**: `GET`

**Path Parameters**:

- `groupName` (string, required): Name of the group for which to retrieve the average temperature.

**Response**:

- `200 OK`: Returns the average temperature value as a float64.
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---

### Get List of Fish Species in Group

Get the full list of fish species (with counts) currently detected inside the specified group.

**URL**: `/group/{groupName}/species`

**Method**: `GET`

**Path Parameters**:

- `groupName` (string, required): Name of the group for which to retrieve the list of detected fish species.

**Response**:

- `200 OK`: Returns a map of fish species names (string) with their respective counts (int).
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---

### Get Top N Fish Species in Group

Get the list of top N fish species (with counts) currently detected inside the specified group.

**URL**: `/group/{groupName}/species/top/{N}`

**Method**: `GET`

**Path Parameters**:

- `groupName` (string, required): Name of the group for which to retrieve the top N detected fish species.
- `N` (int, required): Number of top fish species to retrieve.

**Response**:

- `200 OK`: Returns a map of the top N fish species names (string) with their respective counts (int).
- `400 Bad Request`: If the `N` parameter is invalid or not provided, an error response is returned.
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---

### Get Minimum Temperature in Region

Retrieve the current minimum temperature inside the specified region.

**URL**: `/region/temperature/min?xMin=<xMin>&xMax=<xMax>&yMin=<yMin>&yMax=<yMax>&zMin=<zMin>&zMax=<zMax>`

**Method**: `GET`

**Query Parameters**:

- `xMin` (int, required): Minimum value for the x-coordinate of the region.
- `yMin` (int, required): Minimum value for the y-coordinate of the region.
- `zMin` (int, required): Minimum value for the z-coordinate of the region.
- `xMax` (int, required): Maximum value for the x-coordinate of the region.
- `xMin` (int, required): Maximum value for the y-coordinate of the region.
- `xMin` (int, required): Maximum value for the z-coordinate of the region.

**Response**:

- `200 OK`: Returns the minimal temperature inside the specified region
- `400 Bad Request`: If the query parameters are invalid or not provided, an error response is returned.
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---

### Get Maximum Temperature in Region

Retrieve the current maximum temperature inside the specified region.

**URL**: `/region/temperature/max?xMin=<xMin>&xMax=<xMax>&yMin=<yMin>&yMax=<yMax>&zMin=<zMin>&zMax=<zMax>`

**Method**: `GET`

**Query Parameters**:

- `xMin` (int, required): Minimum value for the x-coordinate of the region.
- `yMin` (int, required): Minimum value for the y-coordinate of the region.
- `zMin` (int, required): Minimum value for the z-coordinate of the region.
- `xMax` (int, required): Maximum value for the x-coordinate of the region.
- `xMin` (int, required): Maximum value for the y-coordinate of the region.
- `xMin` (int, required): Maximum value for the z-coordinate of the region.

**Response**:

- `200 OK`: Returns the maximum temperature inside the specified region
- `400 Bad Request`: If the query parameters are invalid or not provided, an error response is returned.
- `500 Internal Server Error`: If there is an error while processing the request, an error response with the error message is returned.

---
