# Pruebas manuales

Abre [Swagger UI público](https://interseguro-matrix-api.onrender.com/docs) o [Swagger UI local](http://localhost:8080/docs). Para levantar la versión local, sigue el [README](../README.md#ejecutar-en-tu-equipo).

En Swagger, abre la ruta de cada caso, pulsa **Try it out**, pega el JSON y pulsa **Execute**. Compara el código y el resultado con **Server response**. En Postman, selecciona **POST**, combina la URL base con la ruta del caso y pega el JSON en **Body → raw → JSON**. La API pública se puede probar sin credenciales.

## Rotación y estadísticas

`POST /api/v1/matrices/analyze`

```json
{"matrix":[[1,2,3],[4,5,6]],"operation":"rotate"}
```

**HTTP 200.** Matriz rotada `[[4,1],[5,2],[6,3]]`, suma `21` y promedio `3.5`. Este caso comprueba que Go llama a Node.

## QR y estadísticas

`POST /api/v1/matrices/analyze`

```json
{"matrix":[[1,2],[3,4],[5,6]],"operation":"qr"}
```

**HTTP 200.** El resultado incluye Q de 3×2, R de 2×2 y `statistics.matrixCount: 2`. Los valores de Q y R se expresan en punto flotante.

## Solo rotación

`POST /api/v1/matrices/rotate`

```json
{"matrix":[[1,2,3],[4,5,6]],"direction":"clockwise"}
```

**HTTP 200.** Devuelve `matrix: [[4,1],[5,2],[6,3]]` y `rotationDegrees: 90`.

## Solo QR

`POST /api/v1/matrices/qr`

```json
{"matrix":[[1,2],[3,4],[5,6]]}
```

**HTTP 200.** Devuelve Q de 3×2, R de 2×2 y `convention: "thin"`.

## Matriz inválida

`POST /api/v1/matrices/analyze`

```json
{"matrix":[[1,2],[3]],"operation":"rotate"}
```

**HTTP 422.** Devuelve `error.code: "INVALID_MATRIX"` porque las filas tienen diferente longitud.
