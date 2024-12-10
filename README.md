# Public Vault Microservice

## Descripción
Este microservicio permite:
- Tokenizar números de tarjeta de crédito (convertir un número de tarjeta en un token único).
- Detokenizar (recuperar el número de tarjeta asociado a un token).

La aplicación utiliza encriptación AES-256 para proteger los datos sensibles y almacena la información en una base de datos SQLite.

## Estructura del Proyecto

```plaintext
public-vault-ms/
├── config/       # Configuración del sistema
├── controllers/  # Controladores de los endpoints
├── database/     # Inicialización de la base de datos
├── models/       # Modelos de datos
├── services/     # Lógica de negocio
├── utils/        # Utilidades para encriptación y generación de UUID
└── main.go       # Punto de entrada de la aplicación
```

## Requisitos Previos
- Go 1.19 o superior.
- SQLite3 instalado.
- Configuración opcional mediante variables de entorno:
  - `PORT`: Puerto donde se ejecutará el servicio (por defecto: 8080).
  - `ENCRYPTION_KEY`: Clave de encriptación AES-256 (por defecto: `public-vault-key-test-cba-2023`).

## Instalación y Ejecución

1. Clona el repositorio:
   ```bash
   git clone <url-repositorio>
   cd public-vault-ms
   ```

2. Instala las dependencias:
   ```bash
   go mod tidy
   ```

3. Inicia la aplicación:
   ```bash
   go run main.go
   ```

## Endpoints

### 1. Tokenización de Tarjeta
**POST** `/tokenize`

**Request Body:**
```json
{
  "card_number": "4111111111111111"
}
```

**Response:**
```json
{
  "token": "generated-token-uuid"
}
```

### 2. Detokenización de Tarjeta
**POST** `/detokenize`

**Request Body:**
```json
{
  "token": "generated-token-uuid"
}
```

**Response:**
```json
{
  "card_number": "4111111111111111"
}
```

## Base de Datos
La base de datos SQLite contiene una tabla `cards` con las siguientes columnas:
- `token` (clave primaria): Token generado (UUID).
- `encrypted_card`: Número de tarjeta encriptado.

## Seguridad
- Los números de tarjeta se encriptan utilizando AES-256.
- La clave de encriptación puede configurarse mediante la variable de entorno `ENCRYPTION_KEY`.

## Testing
Puedes probar los endpoints utilizando herramientas como `curl`, Postman o scripts personalizados.

## Licencia
Este proyecto está licenciado bajo los términos de la licencia MIT. Consulta el archivo `LICENSE` para más detalles.