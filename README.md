# EWF - Enterprise Web Framework

A Spring Boot 3.2.x Product Catalog REST API.

## Tech Stack

- **Java 17** + **Spring Boot 3.2.5**
- **Spring Data JPA** with H2 in-memory database
- **Bean Validation** (`@Valid`, `@NotBlank`, `@DecimalMin`, etc.)
- **Spring Boot Actuator** (health & info endpoints)

## Running the Application

```bash
mvn spring-boot:run
```

The API starts on `http://localhost:8080`.

## API Endpoints

| Method | URL                    | Description         |
|--------|------------------------|---------------------|
| GET    | /api/products          | List all products   |
| GET    | /api/products/{id}     | Get product by ID   |
| POST   | /api/products          | Create product      |
| PUT    | /api/products/{id}     | Update product      |
| DELETE | /api/products/{id}     | Delete product      |

### Example: Create a product

```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{"name":"Widget","description":"A nice widget","price":4.99,"stock":100}'
```

## H2 Console

Available at `http://localhost:8080/h2-console` (JDBC URL: `jdbc:h2:mem:testdb`).

## Actuator

- `GET /actuator/health`
- `GET /actuator/info`

## Running Tests

```bash
mvn test
```
