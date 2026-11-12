# 06 — Docker, nube y costos

## Local
- Dockerfile Go multi-stage: compilar binario en imagen builder, ejecutar en runtime mínimo sin root.
- Dockerfile Node multi-stage: instalar con `npm ci`, compilar TS, runtime con dependencias de producción y usuario sin privilegios.
- Compose: red interna para ambos, publicar **solo Go** (p.ej. 8080), Node por nombre de servicio `http://statistics-api:3000`, healthchecks, variables de entorno.
- `.env.example` sin secretos. `.dockerignore` por servicio.

## Cloud (decisión sujeta a revisión actual)
- Preferir PaaS que permita dos servicios Docker y comunicación privada (p.ej. Render o Railway, según plan vigente). Verificar condiciones y costos actuales antes de crear recursos.
- Go público con HTTPS; Node interno. Si no existe red privada gratuita, evaluar despliegue de ambos contenedores en una sola VM gratuita o exponer Node con controles y límites, documentando compromisos; nunca fingir privacidad.
- No desplegar dos VMs por defecto; no Kubernetes, API Gateway ni base de datos.
- No elegir región, plan o instancia de pago sin aprobación.
- Si no hay opción gratuita viable, dejar Docker/Compose y guía reproducible, y solicitar decisión de costo antes de activar cloud. El correo sí exige despliegue: priorizar obtener autorización a tiempo.

## Costos y límites
- Docker local, Git y repositorio GitHub pueden usarse sin cargo dentro de sus límites.
- Cloud **puede cobrar** por CPU/RAM, horas activas, tráfico, IP pública, builds o almacenamiento. Un plan gratuito puede dormir, expirar o requerir tarjeta.
- Configurar límites y alertas de presupuesto cuando el proveedor lo permita; las alertas no siempre impiden cobros.
- No guardar tarjetas ni tokens en repo. No comprar dominio: URL de plataforma es suficiente para revisión.
- No borrar otros recursos de la cuenta. Identificar los recursos del challenge con nombres exclusivos y documentar cómo apagarlos tras evaluación.

## Entrega
README con URL pública real, endpoints, ejemplo curl, cómo iniciar con Compose, health, arquitectura, pruebas, decisiones, limitaciones y estado del despliegue. Accesos solo mediante canal seguro; no incluir credenciales en repositorio ni correo en texto plano si puede evitarse.
