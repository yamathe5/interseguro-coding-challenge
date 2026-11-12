# Decisiones de arquitectura

1. Dos servicios en un monorepo para mantener contratos y pruebas de integración en el mismo historial.
2. El dominio y los casos de uso quedan separados del transporte y de las librerías matemáticas.
3. Go llama a Node por HTTP síncrono con timeout y sin reintentos de POST.
4. QR y rotación son operaciones independientes. Analyze calcula estadísticas sobre Q y R o sobre la matriz rotada.
5. Las operaciones son deterministas y no necesitan persistencia ni caché.
6. Se empaquetan servicios Docker independientes. Compose publica Go y conecta ambos mediante una red privada de proyecto.
