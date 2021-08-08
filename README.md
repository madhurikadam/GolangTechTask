# Golang Tech Test

As part of the recruitment process we want to know how you think, code and structure your work.
In order to do that, we're going to ask you to complete this coding challenge.

## Task

In the `api/` folder of this repo there is a basic `grpc` service definition for a voting service.
This service contains RPCs for creating, listing and voting on `voteable` items.

We need you to:
- Implemented the Grpc services with clean/ hexagoanal architecutre 
- Added the `DynamoDB` support as DB, with clean architecture, it is very easy to replace the dynamoDB with any other database
- Seprated the service layer from transport protocol(`grpc`), with this if want to support the http(REST) protocol no change in 
  Service layer required
- Structure of project is like reusable libraries added in the pkg and other packages are like service, repo and model 
- Dependencies of services are injectable, hence it is easy to mock the external dependcies for unit test case 
- Added the unit test cases for service layer and gernated mock for db layer via `mockery` library 


## How to impress us

There are a few optional tasks you can complete if you really want to show off.

1. Adding Observability 

    - Added the zap logger support 
    - using middleware and grpc interceptor(`go-grpc-middleware`) observability can be added like promothis or opentelemery 

2. Adding Configuration and Secrets management
    - Managed the configuration via viper library which help to manage configuration and enviornment variable
3. Added the simple Dockerfile to create the container of service
4. To make build, deployment and easy added makefile which contains all commands to build, deploy and destroy the deployment
5. Inside the docker-compose, added the golangtechtask service 
6. For local debugging added the launch.json with required setting 



## How to setup 
 - User the make commands to start the service
 - Update the compose.env to update enviorment variable
 - steps 
      ```
      $ make binary  # build binary
      $ make docker-build
      $ make docker-run
      $ make docker-down #stop the docker container

## How to use it 
- grpcurl example to test the service
  > ```$grpcurl -v --plaintext -d '{"question":"Since which of the following year Winter Olympics are held ?", "answers":["1896","1900","1888"]}' -proto api/service.proto localhost:9090 VotingService/CreateVoteable ```

  > ```$grpcurl -v --plaintext -d '{"question":"Who had proposed the Olympic motto? ", "answers":["Andrew Parsons","Henri Didon","Carl Jung"]}' -proto api/service.proto localhost:9090 VotingService/CreateVoteable ```
  
  > ``` $ grpcurl -v --plaintext -d '{"pageSize" : 1}' -proto api/service.proto localhost:9090 VotingService/ListVoteables ```

  > ``` grpcurl -v --plaintext -d '{"pageSize" : 1,"nextPageToken": "eyJJRCI6eyJCIjpudWxsLCJCT09MIjpudWxsLCJCUyI6bnVsbCwiTCI6bnVsbCwiTSI6bnVsbCwiTiI6bnVsbCwiTlMiOm51bGwsIk5VTEwiOm51bGwsIlMiOiJiYmZmYmQzYi1jN2VmLTQyYWUtYmQzOS1jNWZhZjNhZDk3NWQiLCJTUyI6bnVsbH19"}' -proto api/service.proto localhost:9090 VotingService/ListVoteables ```

  > ```grpcurl -v --plaintext -d '{"uuid":"cd4eed3e-7aff-4805-a7f4-5bf9a270786a","answer_index":0}' -proto api/service.proto localhost:9090 VotingService/CastVote ```

