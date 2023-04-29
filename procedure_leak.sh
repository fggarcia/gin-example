echo "PROCESSING FILE: albums.json (60 MB).... 0-499_999"
curl -X GET http://localhost:8080/albums/file/albums_2 > /dev/null

sleep 5

echo "CREATING HEAP (heap_post_file.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_post_file.out

sleep 5

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc > /dev/null

sleep 5

echo "CREATING HEAP (heap_post_file_gc.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_post_file_gc.out

echo "CLEAN KEYS LESS ONE 0-499_998...."
curl -X GET http://localhost:8080/albums/delete/0,499999 > /dev/null

sleep 5

echo "CREATING HEAP (heap_delete_less_one.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_delete_less_one.out

sleep 5

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc > /dev/null

echo "CREATING HEAP (heap_post_file_gc.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_delete_less_one_gc.out

sleep 5
echo "CLEAN KEYS"
curl -X GET http://localhost:8080/albums/delete/0,500000 > /dev/null

sleep 5

echo "CREATING HEAP (heap_delete.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_delete.out

sleep 5

echo "RUNNING GC"
curl -X GET http://localhost:8080/albums/gc > /dev/null

echo "CREATING HEAP (heap_delete_gc.out)...."
curl -X GET http://localhost:8080/debug/pprof/heap > heap_delete_gc.out
