# Earthworks Instrumentation

gRPC server/client implementation for transmitting instrumentation readings to a central data store

## Services
### therm

The therm service allows an "thermistor" type instrument to send automated readings in Ohms. A thermistor's resistance varies when it is heated or cooled, and the measured resistance can be converted a temperature. On a successful request, the server records the reading and returns an empty response.

  * resistance: ohms
  * device: an ID for the sending device (e.g. thermistor string number)
