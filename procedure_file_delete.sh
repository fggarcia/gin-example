echo "CREATING FILE..."

./create_files.sh

sleep 1

echo "PROCESSING FILE albums_2.json..."

curl -X GET http://localhost:8080/albums/file/albums_2

sleep 1

echo "DELETING LESS THAN 1 ELEMENT..."

curl -X GET http://localhost:8080/albums/delete/0,499999 > /dev/null

sleep 1

echo "RUNNING GC..."

curl -X GET http://localhost:8080/albums/gc > /dev/null

sleep 1

echo "GENERATING HEAP..."

curl -X GET http://localhost:8080/debug/pprof/heap > heap.out