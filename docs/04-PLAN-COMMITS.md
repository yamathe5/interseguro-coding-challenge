# 04 — Plan de ejecución y commits reales

Trabaja en `feat/coding-challenge`. Un commit por hito cuando los checks estén verdes; si un hito requiere corrección, crear commit adicional `fix:`. No simular commits ni introducir fechas falsas.

| # | Commit sugerido | Entregable verificable |
|---|---|---|
| 01 | `chore: bootstrap monorepo and architecture docs` | estructura, README, ADR, .gitignore, env.example |
| 02 | `feat(go): initialize fiber server and health endpoints` | Go compila, endpoints health probados |
| 03 | `feat(go): validate matrix input` | DTO y tests casos inválidos |
| 04 | `feat(go): implement matrix rotation` | endpoint + tests rectangulares |
| 05 | `feat(go): implement QR factorization with gonum` | QR + tests matemáticos |
| 06 | `feat(node): implement statistics service` | Express TS, stats y tests |
| 07 | `feat(go): integrate statistics HTTP client` | orquestación y tests con stub HTTP |
| 08 | `build: containerize services with compose` | multi-stage Dockerfiles, Compose ejecutable |
| 09 | `test: add end-to-end and integration checks` | prueba real de Go→Node y errores |
| 10 | `ci: add lint test and docker build workflows` | GitHub Actions funcional |
| 11 | `docs: add OpenAPI deployment and decision records` | Swagger/OpenAPI, README, ejemplos |
| 12 | `chore: harden configuration and prepare delivery` | revisión seguridad, despliegue verificado si autorizado |

## Flujo autónomo
1. Verificar estado del repositorio, ramas y remoto antes de tocar nada.
2. Escribir/ajustar tests y código del hito.
3. Ejecutar checks, revisar diff, registrar commit con mensaje indicado.
4. Actualizar `docs/PROGRESS.md` con SHA real, evidencia, pendientes y decisiones.
5. Si una etapa bloquea otra, corregir; si requiere aprobación de gasto o credenciales, detener únicamente esa acción y continuar trabajo local.
6. Al final preparar PR o dejar rama lista; no forzar merge si hay políticas de revisión.

## Prioridad si se acaba el tiempo
QR correcto > estadísticas > integración HTTP > Docker > despliegue > pruebas y README > CI > extras. Rotación también es exigencia textual: implementarla temprano por su bajo costo. No añadir JWT/frontend salvo que todo lo anterior esté validado.
