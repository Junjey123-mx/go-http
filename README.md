# Ejercicio 4 - API JSON en Go

**Tema:** canciones favoritas

Esta API fue desarrollada en Go usando únicamente la librería estándar.  
Permite listar, buscar, crear, actualizar y eliminar canciones favoritas mediante una API JSON.

## Estructura del proyecto

```text
Ejercicio_4/
├── data/
│   └── items.json
├── evidencias/
│   ├── 01-get-all.png
│   ├── 02-post-create.png
│   ├── 03-put-update.png
│   ├── 04-patch-update.png
│   └── 05-delete-item.png
├── main.go
├── Dockerfile
├── docker-compose.yml
└── README.md

Sí, buena observación.

Corregido así:

* **quito las instrucciones de ejecución** del `README`
* **dejo las evidencias apuntando explícitamente a la carpeta `evidencias/`**
* uso los nombres de archivos que creamos:

  * `01-get-all.png`
  * `02-post-create.png`
  * `03-put-update.png`
  * `04-patch-update.png`
  * `05-delete-item.png`

## Comandos para crear la carpeta y archivos de evidencias

Dentro de `Ejercicio_4`:

```bash
mkdir -p evidencias
touch evidencias/01-get-all.png
touch evidencias/02-post-create.png
touch evidencias/03-put-update.png
touch evidencias/04-patch-update.png
touch evidencias/05-delete-item.png
```

Si tus capturas ya existen en otra carpeta, muévelas así:

```bash
mv RUTA_DE_TU_CAPTURA_GET evidencias/01-get-all.png
mv RUTA_DE_TU_CAPTURA_POST evidencias/02-post-create.png
mv RUTA_DE_TU_CAPTURA_PUT evidencias/03-put-update.png
mv RUTA_DE_TU_CAPTURA_PATCH evidencias/04-patch-update.png
mv RUTA_DE_TU_CAPTURA_DELETE evidencias/05-delete-item.png
```

## README.md corregido

````md
# Ejercicio 4 - API JSON en Go

**Tema:** canciones favoritas

Esta API fue desarrollada en Go usando únicamente la librería estándar.  
Permite listar, buscar, crear, actualizar y eliminar canciones favoritas mediante una API JSON.

## Estructura del proyecto

```text
Ejercicio_4/
├── data/
│   └── items.json
├── evidencias/
│   ├── 01-get-all.png
│   ├── 02-post-create.png
│   ├── 03-put-update.png
│   ├── 04-patch-update.png
│   └── 05-delete-item.png
├── main.go
├── Dockerfile
├── docker-compose.yml
└── README.md
````

## Estructura de cada item

Cada canción tiene los siguientes campos:

* `id`
* `cancion`
* `album`
* `autor`
* `genero`
* `anio`
* `sello`

## Endpoints

### GET /api/items

Obtiene todas las canciones registradas.

### GET /api/items?id=1

Obtiene una canción por su id.

### GET /api/items?autor=Eminem

Filtra canciones por autor.

### GET /api/items?genero=Hip Hop

Filtra canciones por género.

### POST /api/items

Crea una nueva canción.

Ejemplo de body:

```json
{
  "cancion": "Numb",
  "album": "Meteora",
  "autor": "Linkin Park",
  "genero": "Nu Metal",
  "anio": 2003,
  "sello": "Warner Bros."
}
```

### PUT /api/items/1

Reemplaza completamente una canción existente.

### PATCH /api/items/1

Actualiza parcialmente una canción existente.

Ejemplo de body:

```json
{
  "genero": "Rock"
}
```

### DELETE /api/items/10

Elimina una canción por id.

## Métodos HTTP implementados

En esta API se implementaron los siguientes métodos:

* `GET`
* `POST`
* `PUT`
* `PATCH`
* `DELETE`

## Evidencias del funcionamiento de la API

### 1. GET /api/items

Se realizó una petición `GET` al endpoint `/api/items` y la API respondió con estado `200 OK`, devolviendo la lista completa de canciones en formato JSON.

![GET de todos los elementos](evidencias/01-get-all.png)

---

### 2. POST /api/items

Se realizó una petición `POST` enviando un nuevo registro en formato JSON.
La API respondió con `201 Created`, confirmando que el recurso fue creado correctamente.

![POST exitoso](evidencias/02-post-create.png)

---

### 3. PUT /api/items/1

Se realizó una petición `PUT` para reemplazar completamente el recurso con id `1`.
La API respondió con `200 OK` y devolvió el objeto actualizado.

![PUT exitoso](evidencias/03-put-update.png)

---

### 4. PATCH /api/items/1

Se realizó una petición `PATCH` para modificar parcialmente el recurso con id `1`.
La API respondió con `200 OK`, confirmando la actualización del registro.

![PATCH exitoso](evidencias/04-patch-update.png)

---

### 5. DELETE /api/items/10

Se realizó una petición `DELETE` sobre el recurso con id `10`.
La API respondió con `200 OK` y el mensaje `item deleted`, indicando que el registro fue eliminado correctamente.

![DELETE exitoso](evidencias/05-delete-item.png)

## Conclusión

Con estas evidencias queda demostrado el funcionamiento correcto de los cinco métodos HTTP principales solicitados:

* `GET`
* `POST`
* `PUT`
* `PATCH`
* `DELETE`

````

Después de eso:

```bash
git add .
git commit -m "Agregar evidencias y actualizar README de Ejercicio 4"
git push
````

Si quieres, ahora te dejo el bloque para agregar la **captura del query param** y la **captura del error 4xx**.
