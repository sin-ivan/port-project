# port-project

How to rebuild proto files:
1) protoc --go_out=plugins=grpc:. *.proto

How to run test:
1) Locate to project root
2) Run "go test ./..."

How to build application:
1) Locate to project root
2) run "docker-compose up"

How to launched services:
1) Type in browser "localhost:3000/parse" to start file parsing and submit parsed data from api-service to domain-service for storing
2) Type in browser "localhost:3000/getAll" to print saved data in JSON format to screen

P.S.
Deletion and Update methods are implemented, but not available through API/
