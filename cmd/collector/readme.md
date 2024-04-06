## Collector
Collector is used to collect requests and packets which are then saved and analysed in other services. At the very end
`Collector` proxies the incoming requests to the next services (i.e. to the proxy, or if request needs to be analysed,
validated by some kind capatcha it proxies it there).