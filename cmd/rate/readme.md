# Rate
Rate is an GRPC server, that stores limits for requests from given ip-address, based on the yaml configuration.
Other services are calling it, to get information about limits for given ip address (ideally limits should be cached in 
the other services in order to limit requests to this single instance).