## This is just a project to have like a portfolio to show
### (UNDER HEAVY DEVELOPMENT)


</br>
</br>
</br>
</br>

This project does not have any goal or anything. It's just using technologies that I have not used or dont have alot of expirence combined with things I have experience with and know very well. It's here to show what I know and also if I find something interesting I'm going to implemented as part of this project. It's going to be just another service or something like that depends on what the technology is

**TODO**:
- add grpc streaming between secrvices 
- add TLS and authentication for grpc 
- add TLS and authentication for http
- add postgres migrations (https://github.com/golang-migrate/migrate) for todo service
- add swagger for user service (like the one for user service)
- maybe add swagger for grpc service with grpc to REST conversion
- setup config file or enviroment variables instead of hardcoded strings (this is low priority because it's the easiest thing to do but just requires time that I want to spend elsewhere right now)
- replace the terrible error handling in user service and add logging
- clean up mongo collections so they are not hardcoded
- finish tests inside todo service and add test to the user service
- implement just RPC between services (currently there is only gRPC)
- setup docker compose for easier local development when more services are created
- create kubernetes files for services 
- setup prometheus (prometheus it's self and the middleware inside the services so prometheus can collect the metrics)
- setup CI/CD with Github actions (or maybe CircleCI have not jet decided)
- when everything is setup for services setup the make file for ease of life 
- setup helm and terraform when actual servers come into play
