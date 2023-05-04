echo "PROCESSING FILE: albums_0-1_000_000.json (120 MB)...."
curl -X GET http://localhost:8080/albums/file/albums_0-1_000_000

sleep 1

echo "CREATING FAST CACHE HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_cache_map.out

sleep 1

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc

sleep 5

echo "CREATING FAST CACHE HEAP POST GC...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_cache_map_gc.out
