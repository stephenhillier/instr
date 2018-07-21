# Earthworks Instrumentation

gRPC server/client implementation for transmitting instrumentation readings to a central data store

## Server

The **instr** server receives incoming connections through gRPC from clients and stores readings in a database. Currently supports PostgreSQL (see `backend/database` folder)

## Clients
### therm

The therm client collects automated readings from a thermistor. A thermistor's resistance varies when it is heated or cooled, and the measured resistance can be converted a temperature. One application of a thermistor for civil engineering is for observing seasonal or annual trends in ground temperature.

## Instructions

### Running the Server

To run the server, a [PostgreSQL](https://www.postgresql.org/) connection is required.  The application is written in Go.

Create a database user and a database. ([PostgreSQL tutorial](https://www.postgresql.org/docs/current/static/tutorial.html))

Create a table with resistance and device columns:
```sql
CREATE TABLE reading(
  id SERIAL PRIMARY KEY,
  resistance FLOAT NOT NULL,
  device TEXT NOT NULL CHECK (char_length(device) < 100)
)
```

Set the following environment variables:

* DBHOST: the database hostname (localhost if running on the same machine)
* DBPORT: the database port (postgres default is 5432)
* DBNAME: the name of your database
* DBUSER: the database user with privileges on the database

Make sure the database is ready to accept connections and start the instrumentation server: from the `backend/server` folder, run `go run main.go`. Alternatively, build the package with `go build main.go` and execute the binary.

### Running the client

The **therm** client can be run on a Raspberry Pi with Docker. An ADS1015 analog to digital converter was used.

Build the executable in the `clients/therm` folder:
```sh
cd clients/therm
GOOS=linux GOARCH=arm go build main.go
```

Build the docker image:
```
docker build -t instr .
```

and run it from the Raspberry Pi (see [Docker documentation](https://docs.docker.com/) for building and pushing your own images). Use the --device flag to allow the container to access `/dev/i2c-1`:
```
docker run --device=/dev/i2c-1 instr
```

When a collection is established, you will begin to see server responses in the client logs and thermistor readings in the server logs and database.
