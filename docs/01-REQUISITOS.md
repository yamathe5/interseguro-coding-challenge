# 01 — Requisitos y trazabilidad

## Fuente: PDF «Coding Challenge», División TI, Interseguro (5 páginas)
- Pág. 2: dos APIs, Go/Fiber y Node.js/Express, código documentado, Docker, comunicación HTTP y despliegue cloud. Arquitectura describe **rotación** en Go seguida de estadísticas en Node.
- Pág. 3: API Go recibe array de arrays rectangular y devuelve **factorización QR**; Node recibe matrices resultantes. También exige implementar rotación. Frontend, JWT y pruebas unitarias/integración se enumeran como opcionales.
- Pág. 4: Node calcula máximo, mínimo, promedio, suma y si alguna matriz es diagonal. Se valoran decisiones justificadas, eficiencia y claridad.
- Correo del proceso (transcrito por candidato): entregar repositorio, despliegue, accesos y credenciales necesarias para revisión, en 24 horas.

## Ambigüedad y resolución propuesta (NO es texto literal del PDF)
QR y rotación son operaciones diferentes. Implementar ambos endpoints y un endpoint de orquestación que calcula QR y pide estadísticas sobre Q y R. Para rotación, permitir enviar la matriz rotada al servicio de estadísticas. Documentar esta interpretación como ADR. No afirmar que la rotación es un paso matemático de QR.

## Obligatorio
1. Matrices rectangulares, validación y errores JSON claros.
2. Factorización QR correcta, con convención de dimensiones documentada.
3. Rotación 90° horario, dirección explícita; permitir variantes solo si son simples.
4. Estadísticas globales sobre todas las matrices recibidas y bandera de alguna diagonal.
5. Comunicación HTTP real Go → Node, con timeout y gestión de errores.
6. Dockerfiles, Compose, documentación y despliegue cloud verificable.
7. Repositorio y acceso a revisión, sin exponer secretos.

## Criterios de calidad añadidos por ingeniería
- Tests de reconstrucción QR, ortogonalidad, matrices rectangulares y rangos deficientes.
- Tests de estadísticas y diagonal con tolerancia; errores 400/422/502/503 según contrato.
- Logs estructurados sin volcar matrices ni tokens; request ID; healthchecks.
- CI que ejecute tests y builds reproducibles.
