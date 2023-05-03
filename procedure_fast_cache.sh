echo "PROCESSING FILE: albums_0-1_000_000.json (120 MB)...."
curl -X GET http://localhost:8080/albums/file_fast_cache/albums_0-1_000_000

sleep 1

echo "CREATING FAST CACHE HEAP...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_fast_cache.out

sleep 1

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc

sleep 1

echo "CREATING FAST CACHE HEAP POST GC...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_fast_cache_gc.out
