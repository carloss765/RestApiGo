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
├── routes/
│   ├── index.routes.go     # Rutas principales
│   ├── tasks.routes.go     # Rutas de tareas
│   └── users.routes.go     # Rutas de usuarios
├── tmp/                    # Archivos temporales de Air
├── .air.toml              # Configuración de Air para hot reload
├── go.mod                 # Dependencias del proyecto
├── go.sum                 # Checksums de dependencias
├── main.go                # Punto de entrada de la aplicación
└── README.md              # Este archivo
```

## 📋 Prerrequisitos

- Go 1.25.3 o superior
- Docker y Docker Compose
- Git

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

## 🛣️ Endpoints Disponibles

### Principal

| Método | Endpoint | Descripción                                    |
| ------ | -------- | ---------------------------------------------- |
| GET    | `/`      | Endpoint de prueba que retorna "Hello, World!" |

### Tareas (En desarrollo)

| Método | Endpoint | Descripción     |
| ------ | -------- | --------------- |
| -      | -        | Por implementar |

### Usuarios (En desarrollo)

| Método | Endpoint | Descripción     |
| ------ | -------- | --------------- |
| -      | -        | Por implementar |

## 🧪 Probar la API

### Usando curl

```bash
curl http://localhost:3000/
```

### Usando un navegador

Abre tu navegador y visita: `http://localhost:3000/`

## 📦 Dependencias Principales

```go
github.com/gorilla/mux v1.8.1           // Router HTTP
gorm.io/gorm v1.31.1                    // ORM
gorm.io/driver/postgres v1.6.0          // Driver PostgreSQL para GORM
```

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

### 🔨 En Desarrollo

- Modelos de base de datos
- CRUD de tareas
- CRUD de usuarios
- Autenticación y autorización
- Validación de datos
- Manejo de errores

## 📝 Notas

- Asegúrate de que el contenedor de PostgreSQL esté corriendo antes de iniciar la aplicación
- La aplicación se conecta automáticamente a la base de datos al iniciar
- Si hay errores de conexión, verifica que las credenciales en `db/conection.go` coincidan con las de tu contenedor Docker

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
