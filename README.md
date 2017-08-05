# source-server-stats
A micro-service that keeps track of Source server stats in a time series database.

## This project will

* periodically query the SteamPipe microservice over grpc 

* store query results in InfluxDB

* include a front-end application to display charts