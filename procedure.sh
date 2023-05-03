echo "PROCESSING FILE: albums_0-1_000_000.json (120 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_0-1_000_000

sleep 1

echo "CREATING FIRST HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_first.out

sleep 1

echo "PROCESSING FILE: albums_0-500_000.json (60 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_0-500_000

sleep 1

echo "CREATING SECOND HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_not_alloc.out

sleep 1

echo "PROCESSING FILE: albums_500_000-2_000_000.json (180 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_500_000-2_000_000

sleep 1

echo "CREATING THIRD HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_alloc.out

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc

sleep 1

echo "CREATING FORTH HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_gc.out
