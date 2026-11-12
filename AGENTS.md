# Instrucciones para Codex — Interseguro Coding Challenge

Lee primero todos los archivos de `docs/` y el PDF original del reto si está disponible. Implementa el proyecto de extremo a extremo conforme a estos documentos. No inventes resultados de pruebas, commits, despliegues ni URLs. Si el PDF no está disponible, utiliza los requisitos transcritos en estos documentos y declara esa limitación.

## Reglas de ejecución
- Monorepo con dos servicios desplegables: `matrix-api` (Go + Fiber) y `statistics-api` (Node.js + Express + TypeScript).
- Implementa por etapas, realizando **un commit semántico por etapa** según `docs/04-PLAN-COMMITS.md`. No fabriques historial ni fechas; cada commit debe corresponder a cambios reales y verificados.
- Ejecuta formateo, lint y pruebas aplicables antes de cada commit. Si fallan, corrige antes de avanzar.
- Mantén `main` estable. Usa rama `feat/coding-challenge` y fusiona mediante PR cuando sea viable; si el repositorio está vacío y no hay remoto, documenta el flujo real.
- No introduzcas Redis, base de datos, Kubernetes, colas, API Gateway ni frontend salvo justificación y aprobación explícita: no aportan al alcance obligatorio.
- **Full access no autoriza gastos**: no aprovisiones recursos de pago, no agregues tarjeta ni actives servicios con cobro sin autorización del usuario. Prefiere planes gratuitos y confirma sus condiciones actuales antes de utilizarlos.
- Nunca expongas secretos ni credenciales en código, logs, commits o documentación. No borres recursos remotos existentes ni fuerces push.
- No hagas afirmaciones de cumplimiento productivo sin evidencias. Registra decisiones, riesgos, resultados de pruebas y limitaciones.
- Si la evaluación restringe el uso de IA, respeta sus reglas; facilita una declaración transparente de asistencia por IA si corresponde.

## Criterio de finalización
Ambas APIs funcionan localmente con Docker Compose; QR, rotación y estadísticas están probadas; integración HTTP real; documentación reproducible; CI funcional; despliegue público verificado si hay autorización y recursos disponibles; reporte final con SHA de commits, comandos ejecutados, resultados, URLs y pendientes.
