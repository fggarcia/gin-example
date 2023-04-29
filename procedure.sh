echo "PROCESSING FILE: albums.json (120 MB)...."
curl -X GET http://localhost:8080/albums/file/albums

sleep 5

echo "CREATING FIRST HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_first.out

sleep 5

echo "PROCESSING FILE: albums_2.json (60 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_2

sleep 5

echo "CREATING SECOND HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_not_alloc.out

sleep 5

echo "PROCESSING FILE: albums_3.json (180 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_3

sleep 5

echo "CREATING THIRD HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_alloc.out

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc

sleep 5

echo "CREATING FORTH HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_gc.out
