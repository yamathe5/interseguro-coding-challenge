# Evidencia de entrega — 12 de noviembre de 2026 (Lima)

Rama local: `feat/coding-challenge`. El repositorio se inicializó en la carpeta del reto. `main` apunta al arranque y la rama de trabajo contiene la implementación. El PDF original no estaba en la carpeta. Se usó la transcripción de `docs/01-REQUISITOS.md`.

| SHA | Etapa |
|---|---|
| `7784a03` | Arranque del monorepo y decisiones |
| `7e11a7c` | Fiber y rutas de salud |
| `cf0a433` | Validación y rotación, verificadas juntas |
| `42ea4b9` | QR con Gonum |
| `3570b68` | Estadísticas en Node |
| `c9569d1` | Cliente HTTP e integración de Go |
| `4d44ee5` | Dockerfiles y Compose |
| `1748490` | Pruebas de integración |
| `7e846a9` | GitHub Actions |
| `b967e97` | OpenAPI y guía de despliegue |

El commit que incluye este documento registra el cierre y los ajustes de IDs, manejo de errores y dependencias de desarrollo. Su SHA se consulta con `git rev-parse HEAD`. La validación y la rotación quedaron juntas porque Windows bloqueó transitoriamente el ejecutable de pruebas durante el primer intento de la etapa de validación. Se volvió a ejecutar la suite y ambos cambios pasaron antes de registrarse.

## Estructura

- `matrix-api/cmd/api` e `internal/{domain,application,adapters}` contienen Go, las operaciones y el transporte HTTP.
- `statistics-api/src/{domain,application,http}` y `statistics-api/test` contienen Node y su suite.
- `compose.yaml` y los Dockerfiles ejecutan ambos servicios en una red privada para Node.
- `scripts/` contiene las pruebas HTTP con y sin contenedores. `.github/workflows/ci.yml` ejecuta los checks en GitHub Actions.
- `docs/` reúne requisitos, ADR, OpenAPI, guía y evidencias.

## Verificación local

| Comando | Resultado |
|---|---|
| `gofmt -l .`, `go mod verify`, `go vet ./...`, `go test ./...`, `go build ./...` (en matrix-api) | Sin archivos por formatear. Módulos verificados y pruebas Go aprobadas. |
| `npm ci`, `npm run lint`, `npm test`, `npm run build`, `npm audit` (en statistics-api) | Instalación reproducible, lint y compilación correctos, dos pruebas aprobadas y cero hallazgos de npm audit. |
| `node scripts/smoke-native.mjs matrix-api/matrix-api.exe` | HTTP real Go → Node. QR, rotación, estadísticas e IDs verificados. Con Node detenido, QR respondió 200, análisis 502 y readiness 503. |
| `node --check scripts/smoke.mjs` y `bash -n scripts/smoke-compose.sh` | Sintaxis válida. |

## Verificación con Docker Compose

Docker Engine 29.1.3 y Compose 2.40.3 se ejecutaron en Ubuntu sobre WSL 2. La primera instalación de Docker Desktop para todos los usuarios no se completó por la solicitud de elevación. El reto solo necesita Docker y Compose.

Desde la raíz del repo, en PowerShell:

```powershell
wsl -d Ubuntu -u root -- env PUBLIC_PORT=18080 docker compose -p interseguro-verificacion config -q
wsl -d Ubuntu -u root -- env PUBLIC_PORT=18080 docker compose -p interseguro-verificacion up --build --wait --wait-timeout 90
$env:MATRIX_URL = 'http://127.0.0.1:18080'
node scripts/smoke.mjs
wsl -d Ubuntu -u root -- docker compose -p interseguro-verificacion stop statistics-api
node scripts/smoke.mjs --dependency-down
wsl -d Ubuntu -u root -- docker compose -p interseguro-verificacion down
```

Las dos imágenes se construyeron y los dos contenedores quedaron saludables. Compose publicó solo Go. Las pruebas HTTP pasaron con los dos servicios activos y con Node detenido. Las consultas de contenedores y redes filtradas por el proyecto quedaron vacías después de `down`.

## Estado externo

El workflow quedó versionado y los checks equivalentes se ejecutaron localmente. No hay ejecución remota de GitHub Actions, PR ni repositorio publicado porque no se configuró un remoto de entrega. Tampoco hay despliegue público ni URL verificable porque no hay un destino cloud configurado. La guía describe cómo desplegar y verificar cuando haya un destino. No se aprovisionaron recursos con cobro.

Codex asistió en la implementación, documentación y ejecución de pruebas. El historial y los resultados locales permiten revisar esa contribución.
