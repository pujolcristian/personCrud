# PersonAPI

## Descripción general

**PersonAPI** es un proyecto creado con Go que ofrece una API CRUD para gestionar una lista de personas. Este archivo lo
guiará para configurar y ejecutar el proyecto localmente por primera vez.

---

## **Requisitos previos**

Antes de comenzar, asegúrese de tener instalado lo siguiente:

- **Go** (1.18 o superior): [Instalar Go](https://golang.org/dl/)
- **MySQL** (8.0 o superior): [Descargar MySQL](https://dev.mysql.com/downloads/installer/)
- **Git**: [Descargar Git](https://git-scm.com/)

---

## **Instrucciones de configuración**

### **1. Clonar el repositorio**

```bash
git clone https://github.com/pujolcristian/PersonAPI.git
cd PersonAPI
```

### **2. Instalar dependencias**

```bash
go mod tidy
```

### **3. Configurar la base de datos**

- Instalar MySQL
- Crear una base de datos llamada `databaseperson`

Conéctese a su servidor MySQL y ejecute los siguientes comandos para crear la base de datos y la tabla:

```sql
CREATE
DATABASE databaseperson;

USE
databaseperson;

CREATE TABLE persons
(
    id        INT PRIMARY KEY AUTO_INCREMENT,
    name      VARCHAR(50)  NOT NULL,
    address   VARCHAR(100) NOT NULL,
    phone     BIGINT       NOT NULL,
    email     VARCHAR(100),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### **4. Actualizar la URL base**

**a.** Buscar la IP local de su máquina:

- Windows: Ejecute ipconfig en el terminal.
- Mac/Linux: Ejecute ifconfig en el terminal.

**b.** Configurar el archivo de conexión

- En el archivo internal/adapters/db/db_config.go, reemplace 127.0.0.1 con su dirección IP local:

```go
dsn := "root:password1234@tcp(<YOUR_IP>:3306)/person_crud?charset=utf8mb4&parseTime=True&loc=Local"
```

### **5. Ejecutar el proyecto**

```bash
go run main.go
```

---

## **Rutas de la API**

| Método     | Endpoint        | Descripción                     |
|------------|-----------------|---------------------------------|
| **GET**    | `/persons`      | Obtiene todas las personas      |
| **GET**    | `/persons/{id}` | Obtiene una persona por ID      |
| **POST**   | `/persons`      | Crea una nueva persona          |
| **PUT**    | `/persons/{id}` | Actualiza una persona existente |
| **DELETE** | `/persons/{id}` | Elimina una persona por ID      |

### Ejemplo: **Crear una nueva persona**

```curl
curl -X POST http://<YOUR_IP>:8080/persons \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "address": "123 Main St",
  "phone": 1234567890,
  "email": "johndoe@example.com"
}'
```