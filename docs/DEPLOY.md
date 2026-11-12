# Despliegue para evaluación

La solución usa dos imágenes independientes y una red privada. En un host Linux con Docker Engine, Compose y un proxy HTTPS existente, clona el repositorio y ejecuta `docker compose up --build --wait`. Mantén privado el puerto 3000. Si el proxy está en el mismo host, enlaza el puerto 8080 a loopback mediante un ajuste de Compose. Configura el proxy para TLS y reenvío a 8080. En una plataforma administrada configura el servicio Go con acceso público, Node en la red privada y `STATISTICS_URL` con el nombre interno y puerto de Node.

Render permite servicios web gratuitos pero su plan gratuito no admite recibir tráfico por red privada. Para mantener Node privado, se necesita un host existente con Docker o un plan con red privada. Consulta [las condiciones de Render](https://render.com/docs/free) y los costos reales del proveedor antes de aprovisionar.

Verifica desde otra red con `MATRIX_URL=https://<host> node scripts/smoke.mjs`. El comando valida salud, QR, rotación, estadísticas e IDs de solicitud. Registra la URL y el SHA de la versión desplegada en `PROGRESS.md`. Para desmontar los contenedores de esta aplicación, usa `docker compose down` en el directorio del despliegue.

No hay URL pública asociada a esta entrega local. El costo incremental de la ejecución local es cero en servicios contratados. Un despliegue en la nube depende del host, plan y tráfico seleccionados.
