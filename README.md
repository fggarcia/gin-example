# EXAMPLE GIN + GO-JSON + HUGE File

PARA CREAR LA DATA A PROCESAR PRIMERO CORRER EL SCRIPT
```
./create_files.sh
```
ESTO CREA LOS 3 ARCHIVOS QUE SE MENCIONAN A CONTINUACION

#### ESTE ARCHIVO CREA EN EL CACHEMAP ALBUMS CON ID DEL 0 AL 999_999_999
```
curl -X GET http://localhost:8080/albums/file/albums
```

#### ESTE CREA DEL 0 AL 499_999 
```
curl -X GET http://localhost:8080/albums/file/albums_2
```

#### ESTE CREA DEL 0 AL 500_000 AL 1_999_999 (PARA PROBAR DE REPISAR KEYS)
```
curl -X GET http://localhost:8080/albums/file/albums_3
```

#### PARA CREAR UN PROFILE
```
curl -X GET http://localhost:8080/debug/pprof/heap > heap.out
```

#### PARA CORRE UN GC MANUALMENTE
```
curl -X GET http://localhost:8080/albums/gc
```


#### OTROS EJEMPLOS /ENDPOINTS
##### CREA UN ARCHIVO CON IDS DEL 0 AL 99_999 POR EJEMPLO
```
curl -X GET http://localhost:8080/albums/create/0,100000 > data.json
```