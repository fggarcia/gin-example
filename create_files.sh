#/bin/bash

curl -X GET http://localhost:8080/albums/create/0,1000000 > albums_0-1_000_000.json
curl -X GET http://localhost:8080/albums/create/0,500000 > albums_0-500_000.json
curl -X GET http://localhost:8080/albums/create/500000,2000000 > albums_500_000-2_000_000.json
