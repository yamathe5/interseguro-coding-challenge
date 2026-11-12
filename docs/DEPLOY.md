# Despliegue para evaluación

La configuración local con Compose usa dos imágenes independientes y una red privada. En un host Linux con Docker Engine, Compose y un proxy HTTPS existente, clona el repositorio y ejecuta `docker compose up --build --wait`. Mantén privado el puerto 3000. Si el proxy está en el mismo host, enlaza el puerto 8080 a loopback mediante un ajuste de Compose. Configura el proxy para TLS y reenvío a 8080. En una plataforma administrada configura el servicio Go con acceso público, Node en la red privada y `STATISTICS_URL` con el nombre interno y puerto de Node.

Render permite servicios web gratuitos pero su plan gratuito no admite recibir tráfico por red privada. Para mantener Node privado, se necesita un host existente con Docker o un plan con red privada. Consulta [las condiciones de Render](https://render.com/docs/free) y los costos reales del proveedor antes de aprovisionar.

Verifica desde otra red con `MATRIX_URL=https://<host> node scripts/smoke.mjs`. El comando valida salud, QR, rotación, estadísticas e IDs de solicitud. Registra la URL y el SHA de la versión desplegada en `PROGRESS.md`. Para desmontar los contenedores de esta aplicación, usa `docker compose down` en el directorio del despliegue.

La demo está desplegada en Render como dos servicios web gratuitos. Go se comunica con Node por HTTPS. Abre [Swagger UI público](https://interseguro-matrix-api.onrender.com/docs) y sigue la [guía de pruebas manuales](PRUEBAS-MANUALES.md). Si Node está en reposo, abre su [ruta de disponibilidad](https://interseguro-coding-challenge-b623.onrender.com/health/ready) y vuelve a probar.
