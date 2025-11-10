# REST API con Go

API REST desarrollada con Go, utilizando GORM como ORM y PostgreSQL como base de datos, ejecutándose en un contenedor Docker.

## 🚀 Tecnologías Utilizadas

- **Go 1.25.3** - Lenguaje de programación
- **Gorilla Mux** - Router HTTP para Go
- **GORM** - ORM (Object Relational Mapping) para Go
- **PostgreSQL** - Base de datos relacional
- **Docker** - Contenedor para PostgreSQL
- **Air** - Hot reload para desarrollo en Go

## 📁 Estructura del Proyecto

```
Rest-api-go/
├── db/
│   └── conection.go        # Configuración de conexión a la base de datos
├── models/
│   ├── Task.go            # Modelo de tareas
│   └── User.go            # Modelo de usuarios
├── routes/
│   ├── index.routes.go     # Rutas principales
│   ├── tasks.routes.go     # Rutas de tareas
│   └── users.routes.go     # Rutas de usuarios
├── tmp/                    # Archivos temporales de Air
├── .air.toml              # Configuración de Air para hot reload
├── .http                  # Ejemplos de peticiones HTTP
├── go.mod                 # Dependencias del proyecto
├── go.sum                 # Checksums de dependencias
├── main.go                # Punto de entrada de la aplicación
└── README.md              # Este archivo
```

## 📋 Prerrequisitos

- Go 1.25.3 o superior
- Docker y Docker Compose
- Git

## 🚀 Quick Start

### 1. Configurar PostgreSQL

```bash
docker run --name postgres-gorm -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=gorm -p 5432:5432 -d postgres
```

### 2. Clonar e instalar

```bash
git clone https://github.com/carloss765/RestApiGo.git
cd RestApiGo
go mod download
```

### 3. Ejecutar la aplicación

```bash
go run main.go
# o con hot reload
air
```

### 4. Probar la API

```bash
# Crear un usuario
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"firstname":"John","lastname":"Doe","email":"john@example.com"}'

# Crear una tarea
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Mi tarea","description":"Descripción","user_id":1}'

# Ver todos los usuarios
curl http://localhost:3000/users
```

## 🐳 Configuración de PostgreSQL con Docker

### 1. Crear y ejecutar el contenedor de PostgreSQL

```bash
docker run --name postgres-gorm -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=gorm -p 5432:5432 -d postgres
```

### 2. Verificar que el contenedor está corriendo

```bash
docker ps
```

### 3. Comandos útiles de Docker

```bash
# Detener el contenedor
docker stop postgres-gorm

# Iniciar el contenedor
docker start postgres-gorm

# Ver logs del contenedor
docker logs postgres-gorm

# Eliminar el contenedor
docker rm postgres-gorm
```

## 🔧 Instalación

### 1. Clonar el repositorio

```bash
git clone https://github.com/carloss765/RestApiGo.git
cd RestApiGo
```

### 2. Instalar dependencias

```bash
go mod download
```

### 3. Instalar Air para hot reload (opcional pero recomendado)

```bash
go install github.com/air-verse/air@latest
```

## ▶️ Ejecución

### Opción 1: Ejecutar con Go

```bash
go run main.go
```

### Opción 2: Ejecutar con Air (Hot Reload)

```bash
air
```

La aplicación estará disponible en: `http://localhost:3000`

## 🔌 Configuración de Base de Datos

La configuración de conexión a la base de datos se encuentra en `db/conection.go`:

```go
DSN = "host=localhost user=admin password=admin dbname=gorm port=5432"
```

### Variables de conexión:

- **Host**: localhost
- **Usuario**: admin
- **Contraseña**: admin
- **Base de datos**: gorm
- **Puerto**: 5432

## 📊 Modelos de Base de Datos

### Usuario (User)

```go
type User struct {
    gorm.Model
    FirstName string   // Nombre del usuario
    LastName  string   // Apellido del usuario
    Email     string   // Email único del usuario
    Tasks     []Task   // Tareas asociadas al usuario
}
```

### Tarea (Task)

```go
type Task struct {
    gorm.Model
    Title       string // Título único de la tarea
    Description string // Descripción de la tarea
    Done        bool   // Estado de la tarea (false por defecto)
    UserID      uint   // ID del usuario propietario
}
```

**Relación:** Un usuario puede tener múltiples tareas (One-to-Many).

## 🛣️ Endpoints Disponibles

### Principal

| Método | Endpoint | Descripción                                    |
| ------ | -------- | ---------------------------------------------- |
| GET    | `/`      | Endpoint de prueba que retorna "Hello, World!" |

### Usuarios

| Método | Endpoint      | Descripción                              |
| ------ | ------------- | ---------------------------------------- |
| GET    | `/users`      | Obtiene todos los usuarios               |
| GET    | `/users/{id}` | Obtiene un usuario por ID con sus tareas |
| POST   | `/users`      | Crea un nuevo usuario                    |
| DELETE | `/users/{id}` | Elimina un usuario por ID                |

### Tareas

| Método | Endpoint      | Descripción              |
| ------ | ------------- | ------------------------ |
| GET    | `/tasks`      | Obtiene todas las tareas |
| GET    | `/tasks/{id}` | Obtiene una tarea por ID |
| POST   | `/tasks`      | Crea una nueva tarea     |
| DELETE | `/tasks/{id}` | Elimina una tarea por ID |

## 🧪 Probar la API

### Usando curl

#### Crear un usuario

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"firstname":"John","lastname":"Doe","email":"john.doe@example.com"}'
```

#### Obtener todos los usuarios

```bash
curl http://localhost:3000/users
```

#### Obtener un usuario específico

```bash
curl http://localhost:3000/users/1
```

#### Crear una tarea

```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Mi primera tarea","description":"Descripción de la tarea","user_id":1}'
```

#### Obtener todas las tareas

```bash
curl http://localhost:3000/tasks
```

#### Obtener una tarea específica

```bash
curl http://localhost:3000/tasks/1
```

#### Eliminar una tarea

```bash
curl -X DELETE http://localhost:3000/tasks/1
```

#### Eliminar un usuario

```bash
curl -X DELETE http://localhost:3000/users/1
```

### Usando el archivo .http

El proyecto incluye un archivo `.http` con ejemplos de todas las peticiones. Puedes usar extensiones como REST Client para VS Code para ejecutar estas peticiones directamente desde el editor.

## 📋 Ejemplos de Respuestas

### Crear Usuario (POST /users)

**Request:**

```json
{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@example.com"
}
```

**Response:**

```json
{
  "ID": 1,
  "CreatedAt": "2024-11-10T15:04:05.999Z",
  "UpdatedAt": "2024-11-10T15:04:05.999Z",
  "DeletedAt": null,
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@example.com",
  "tasks": null
}
```

### Obtener Usuario con Tareas (GET /users/1)

**Response:**

```json
{
  "ID": 1,
  "CreatedAt": "2024-11-10T15:04:05.999Z",
  "UpdatedAt": "2024-11-10T15:04:05.999Z",
  "DeletedAt": null,
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@example.com",
  "tasks": [
    {
      "ID": 1,
      "CreatedAt": "2024-11-10T15:10:20.999Z",
      "UpdatedAt": "2024-11-10T15:10:20.999Z",
      "DeletedAt": null,
      "title": "Mi primera tarea",
      "description": "Descripción de la tarea",
      "done": false,
      "user_id": 1
    }
  ]
}
```

### Crear Tarea (POST /tasks)

**Request:**

```json
{
  "title": "Mi primera tarea",
  "description": "Descripción de la tarea",
  "user_id": 1
}
```

**Response:**

```json
{
  "ID": 1,
  "CreatedAt": "2024-11-10T15:10:20.999Z",
  "UpdatedAt": "2024-11-10T15:10:20.999Z",
  "DeletedAt": null,
  "title": "Mi primera tarea",
  "description": "Descripción de la tarea",
  "done": false,
  "user_id": 1
}
```

## 📦 Dependencias Principales

```go
github.com/gorilla/mux v1.8.1           // Router HTTP
gorm.io/gorm v1.31.1                    // ORM
gorm.io/driver/postgres v1.6.0          // Driver PostgreSQL para GORM
```

## ⚙️ Características Técnicas

### Arquitectura

- **Patrón de diseño**: Separación de responsabilidades
  - `models/`: Definición de estructuras de datos
  - `routes/`: Handlers y lógica de endpoints
  - `db/`: Configuración de base de datos
- **ORM**: GORM para interacción con PostgreSQL
- **Router**: Gorilla Mux para manejo de rutas HTTP

### Funcionalidades GORM

- **AutoMigrate**: Creación automática de tablas basadas en modelos
- **Relaciones**: Soporte de relaciones One-to-Many
- **Soft Delete**: Eliminación lógica de registros con `gorm.Model`
- **Tags de validación**: Constraints a nivel de base de datos
- **Associations**: Carga de relaciones con `.Association()`

### Manejo de Errores

- Validación de IDs inexistentes (404 Not Found)
- Manejo de errores de base de datos (400 Bad Request)
- Validación de constraints únicos (títulos y emails)

## 🔄 Hot Reload con Air

El proyecto está configurado con Air para desarrollo con hot reload. Cuando ejecutas `air`, la aplicación se recarga automáticamente cada vez que guardas cambios en los archivos `.go`.

### Configuración de Air (.air.toml)

- Compila el proyecto en `tmp/main.exe`
- Excluye directorios: assets, tmp, vendor, testdata
- Monitorea archivos con extensiones: .go, .tpl, .tmpl, .html

## 🚧 Estado del Proyecto

### ✅ Completado

- Configuración inicial del proyecto
- Conexión a base de datos PostgreSQL con GORM
- Configuración de Gorilla Mux para routing
- Endpoint principal de prueba
- Configuración de hot reload con Air
- **Modelos de base de datos (User y Task)**
- **CRUD completo de usuarios**
- **CRUD completo de tareas**
- **Relación One-to-Many entre usuarios y tareas**
- **Migraciones automáticas con GORM**
- **Manejo básico de errores**

### 🔮 Mejoras Futuras

- Autenticación y autorización con JWT
- Validación de datos con Go Validator
- Middleware de logging y manejo de errores centralizado
- Paginación de resultados
- Filtros y búsquedas avanzadas
- Actualización de registros (endpoints PUT/PATCH)
- Tests unitarios e integración
- Documentación con Swagger/OpenAPI

## 📝 Notas Importantes

### Configuración Inicial

- Asegúrate de que el contenedor de PostgreSQL esté corriendo antes de iniciar la aplicación
- La aplicación se conecta automáticamente a la base de datos al iniciar
- Si hay errores de conexión, verifica que las credenciales en `db/conection.go` coincidan con las de tu contenedor Docker

### Características Implementadas

- **Migraciones Automáticas**: Las tablas se crean automáticamente al iniciar la aplicación
- **Soft Delete**: Los registros eliminados mantienen un campo `DeletedAt` para recuperación
- **Relaciones**: Los usuarios pueden tener múltiples tareas asociadas
- **Validaciones GORM**:
  - Títulos de tareas son únicos
  - Emails de usuarios son únicos
  - Campos obligatorios definidos en los modelos

### Orden de Uso Recomendado

1. Crear usuarios antes de crear tareas
2. Asignar tareas a usuarios existentes usando `user_id`
3. Consultar usuarios para ver sus tareas asociadas

### Content-Type

- Todas las peticiones POST requieren el header: `Content-Type: application/json`
- Las respuestas se devuelven en formato JSON

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Haz fork del proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto es de código abierto y está disponible bajo la licencia MIT.

## 👤 Autor

**Carlos** - [@carloss765](https://github.com/carloss765)
