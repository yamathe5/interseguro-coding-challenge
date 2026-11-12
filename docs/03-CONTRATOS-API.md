# 03 — Contratos REST propuestos (no impuestos por el PDF)

Base pública Go: `/api/v1`. Base Node: `/internal/v1` (solo red interna cuando sea posible).

## POST /api/v1/matrices/qr
Request: `{"matrix":[[1,2],[3,4],[5,6]]}`.
Response: `{"q":[...],"r":[...],"shape":{"rows":3,"cols":2},"convention":"thin"}`.
Convención **QR reducida** para A m×n con m>=n: Q m×n y R n×n. Para m<n, preferir QR completa Q m×m y R m×n, o un QR reducido general Q m×k y R k×n con k=min(m,n); elegir **una** convención consistente y probarla. Recomendación final: QR reducida general m×k y k×n, k=min(m,n), cuando la librería lo permita. No declarar soporte m<n sin verificarlo.

## POST /api/v1/matrices/rotate
Request: `{"matrix":[[1,2,3],[4,5,6]],"direction":"clockwise"}`.
Response: `{"matrix":[[4,1],[5,2],[6,3]],"rotationDegrees":90}`.

## POST /api/v1/matrices/analyze
Request: `{"matrix":[[1,2],[3,4]],"operation":"qr"}` o `"rotate"`.
Respuesta incluye `operation`, `result` (Q/R o matriz rotada), `statistics` con `max`, `min`, `average`, `sum`, `anyDiagonal`, `diagonalMatrices` y `matrixCount`. Estadísticas QR deben calcularse sobre Q y R, no sobre A. Para rotación, sobre la matriz rotada. Dejar esta semántica explícita.

## POST /internal/v1/statistics
Request: `{"matrices":[[[1,0],[0,2]],[[3,4],[5,6]]]}`.
Response ejemplo: `{"max":6,"min":0,"average":2.625,"sum":21,"anyDiagonal":true,"diagonalMatrices":[0],"matrixCount":2}`.
Los valores del ejemplo corresponden a ocho elementos: suma 21, promedio 2.625.

## Validaciones
- Array externo no vacío; filas no vacías de longitud idéntica; elementos números finitos (no null, string, NaN, Infinity).
- Límite configurable (propuesta 100×100 y <=10 000 elementos), cuerpo <=256 KiB; validar antes de QR.
- Matriz diagonal: cuadrada y elementos fuera de diagonal con `abs(x)<=epsilon`; usar epsilon 1e-9 por defecto, documentar. Una matriz no cuadrada no es diagonal.
- Errores JSON: `{"error":{"code":"INVALID_MATRIX","message":"...","requestId":"..."}}`.
- 400 para JSON inválido; 422 para matriz inválida; 502 para dependencia fallida; 503 si servicio no listo; 500 solo error inesperado.
- Cuidado con redondeos QR: no redondear valores internos; serializar con precisión suficiente y tolerancia en tests.

## Requisitos numéricos de QR
- Usar Gonum `mat.QR` si su API y las dimensiones seleccionadas lo permiten; verificar firmas contra versión fijada.
- Verificar `||A-QR||` relativa <= 1e-9 en casos bien condicionados y `||QᵀQ-I||` con tolerancia apropiada.
- Cubrir 2×2, 3×2, 2×3, identidad, ceros, negativos, matrices degeneradas y entrada no válida. Registrar comportamiento para rango deficiente.
