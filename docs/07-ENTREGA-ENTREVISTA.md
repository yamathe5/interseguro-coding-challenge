# 07 — Entrega y defensa técnica

## Checklist final
- [ ] Repositorio accesible para evaluadores (público o invitación privada).
- [ ] README reproducible desde clon limpio.
- [ ] URLs reales verificadas desde fuera del entorno.
- [ ] API Go pública, Node privada si infraestructura lo permite.
- [ ] Ejemplos JSON y colección Postman o curl.
- [ ] Pruebas Go, Node, integración y Docker con resultados reales.
- [ ] CI en verde y commits progresivos auténticos.
- [ ] No hay secretos, tarjetas ni credenciales en Git.
- [ ] Notas de limitaciones, ambigüedad QR/rotación y decisiones.
- [ ] Acceso y credenciales de revisión, si se requieren, por canal seguro.

## Cómo justificar decisiones
- «Elegí dos servicios porque el reto lo exige; mantuve separación de responsabilidades sin infraestructura extra».
- «Apliqué arquitectura hexagonal ligera para aislar lógica matemática, transporte HTTP y librerías externas, sin capas ceremoniales».
- «No usé Redis: no hay lecturas repetidas de estado ni persistencia, y el costo de invalidación supera el beneficio».
- «No usé Kubernetes ni API Gateway: serían sobreingeniería para dos servicios y un plazo de 24 horas».
- «Implementé QR y rotación por separado porque el PDF describe ambas; la interpretación está documentada».
- «Verifiqué la factorización mediante reconstrucción y ortogonalidad con tolerancia numérica».
- «Los errores de red tienen timeout, trazabilidad y respuesta estable; la API estadística puede ejecutarse y probarse independientemente».

## Uso de IA
No atribuir a trabajo manual código generado automáticamente. Revisar políticas del proceso; si se permite IA, explicar con transparencia la supervisión, decisiones, pruebas y cambios propios. Debes poder explicar y modificar las partes críticas.
