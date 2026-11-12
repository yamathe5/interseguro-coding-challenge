# 05 — Pruebas y calidad

## Go
- Tests table-driven de validación (vacía, filas desiguales, números inválidos, límites).
- Rotación 90° de 2×3 y 3×2, 1×N, Nx1, identidad y cuatro rotaciones sucesivas.
- QR con `A≈Q·R`, `QᵀQ≈I` y R triangular superior, para matrices altas, anchas, cuadradas, con ceros y rango deficiente. Probar semántica real de Gonum; no aceptar tests que solo comparan valores hardcodeados.
- HTTP tests de status, schema, timeout, error de Node y request ID.

## Node
- Estadísticas de múltiples matrices, negativos, ceros, decimales, una diagonal, ninguna diagonal, matrices no cuadradas y matrices inválidas.
- Diagonal con tolerancia; prueba que `anyDiagonal` no marque matrices rectangulares.
- Tests de Express con supertest o equivalente; TypeScript strict, ESLint.

## Integración
- `docker compose up --build -d` y smoke test desde host contra Go.
- Confirmar que Go llama realmente a Node; probar caída de Node y error 502/503.
- Revisar que `docker compose down` limpia recursos creados por el proyecto.

## CI
- Workflow `pull_request`/`push`: Go fmt/vet/test; npm ci/lint/test/build; docker build de ambos.
- Dependencias lockeadas; sin secretos en CI para tests locales.
- No afirmar cobertura específica sin medirla.

## Evidencia final
Guardar comandos exactos, salida resumida, fecha real, commit SHA y URLs. No inventar `passed` si no se ejecutó. Ejecutar un smoke test contra la URL desplegada y guardar resultado sin credenciales.
