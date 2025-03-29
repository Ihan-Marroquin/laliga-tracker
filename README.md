# La Liga Tracker API

La Liga Tracker es una API RESTful desarrollada en Go que permite gestionar partidos de La Liga. La API está construida con [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/) y documentada con [Swagger](https://github.com/swaggo/gin-swagger). Además, la aplicación está dockerizada y utiliza PostgreSQL como base de datos.

## Características

- Listar todos los partidos.
- Obtener detalles de un partido.
- Crear, actualizar y eliminar partidos.
- Incrementar contadores de goles, tarjetas amarillas y rojas.
- Marcar partidos con tiempo extra.
- Documentación interactiva con Swagger.

## Requisitos

- [Go 1.24](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/) (versión 2 o superior)

## Instalación y Ejecución

### Usando Docker

1. **Clona el repositorio:**

   ```bash
   git clone https://github.com/Ihan-Marroquin/laliga-tracker.git
   cd laliga-tracker

2. **Instala las dependencias de Go**

Tienes que correr el siguiente comando en la terminal, en la raiz del proyecto
   ```bash
   go mod download
  ```
3. **Construye y levanta los contenedores con Docker Compose**

Tienes que correr el siguiente comando en la terminal, en la raiz del proyecto
Pero antes de correr el comando, tienes que modificar el archivo **docker-compose.yml**, poniendo tu usuario y contraseña de PostGreSQL
   ```bash
   docker-compose up --build
  ```

Esto levantará dos servicios:
- db: Servicio de PostgreSQL (usuario: tu_usuario, contraseña: tu_contraseña, base de datos: laliga).
- backend: La API, accesible en el puerto 8080.

4. **Accede a la API**
   ```bash
    http://localhost:8080

5. **Acceder a la base de datos**

Para poder consultar la base de datos, se debe correr el siguiente comando en el directorio raiz del proyecto
   ```bash
    docker-compose exec db psql -U postgres -d laliga -c "SELECT * FROM matches;"
   ```
6. **Endpoints de la API**
Todos los endpoints se encuentran bajo el prefijo **/api**:
   ```bash
   GET /api/matches
   Lista todos los partidos.

    GET /api/matches/:id
    Obtiene un partido específico por ID.

    POST /api/matches
    Crea un nuevo partido.

    PUT /api/matches/:id
    Actualiza un partido existente.

    DELETE /api/matches/:id
    Elimina un partido.

    PATCH /api/matches/:id/goals
    Incrementa el contador de goles.

    PATCH /api/matches/:id/yellowcards
    Incrementa el contador de tarjetas amarillas.

    PATCH /api/matches/:id/redcards
    Incrementa el contador de tarjetas rojas.

    PATCH /api/matches/:id/extratime
    Marca el partido como que tuvo tiempo extra.

7. **Imagenes del API corriendo**
**Cargar partidos**
![API funcionando](https://github.com/user-attachments/assets/60eb8f0e-0071-4cd0-96ba-fa2372b1ea94)

**Insertar partidos**
![CREAR PARTIDO](https://github.com/user-attachments/assets/4129fcf6-22dc-4029-a970-fa929eb94a35)

**Partido actualizado**
![partido creado](https://github.com/user-attachments/assets/46cefe48-81ea-4226-b257-4953494ec97e)

**Base de datos con goles, tarjetas y partidos**
![tabla actualizada](https://github.com/user-attachments/assets/42cbcd26-eb15-4c47-b0c1-bdd7749b5977)
