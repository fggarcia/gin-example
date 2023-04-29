# EXAMPLE GIN + GO-JSON + HUGE File

PARA CREAR LA DATA A PROCESAR PRIMERO CORRER EL SCRIPT
```
./create_files.sh
```

SE AUTOMATIZO LOS PASOS DE ABAJO EN UN SCRIPT
```
./procedure.sh
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
##### DELETE IDS DEL MAPA CON RANGO DE IDS
```
curl -X GET http://localhost:8080/albums/delete/0,2000000
```
##### VER TAMAÑO DEL MAPA
```
curl -X GET http://localhost:8080/albums/cache
```

# PROCEDURE LEAK
Para poder probar los distintos escenarios de los serializadores
en el archivo main.go se encuentra la siguiente config:
```
"github.com/goccy/go-json"
//"encoding/json"
//"github.com/go-json-experiment/json"
```

PD: sino estan creados los archivos albums.json, albums_2.json y 
albums_3.json primero correr
```
./create_files.sh
```

Para crear el escenario correr:
```
./procedure_leak.sh
```

se crean los siguientes profiles:

* heap_post_file.out -> profile apenas se terminan de insertar los 500_000 albums
* heap_post_file_gc.out -> profile despues de correr el GC
* heap_delete_less_one.out -> profile despues de eliminar casi todos los albums salvo 1
* heap_delete_less_one_gc.out -> profile despues de correr el GC
* heap_delete.out -> profile despues de eliminar todos los albums
* heap_delete_gc.out -> profile despues de correr el GC

ver cada uno con
```
go tool pprof -http=:<numero_de_puerto> <nombre_del_archivo>.out
```