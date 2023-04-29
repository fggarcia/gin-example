#/bin/bash

curl -X GET http://localhost:8080/albums/create/0,1000000 > albums.json
curl -X GET http://localhost:8080/albums/create/0,500000 > albums_2.json
curl -X GET http://localhost:8080/albums/create/500000,2000000 > albums_3.json
