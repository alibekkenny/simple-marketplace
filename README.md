Simple Marketplace API

A simple marketplace backend written in Golang 1.24 build with **mux router**, and **middleware chain**
It consists of several services that communicate via gRPC. An API Gateway serves as the entry point for clients and routes requests to the appropriate service. Services talk to each other only via gRPC. The API Gateway acts as a bridge between external HTTP requests and internal gRPC calls. No service directly accesses another service’s database.
The project has a /shared folder that acts as a central place for resources used across multiple services. All gRPC services define their contracts (service definitions, request/response messages, etc.) inside .proto files in /shared/proto folder. Each service imports these proto files to generate client and server stub.

--
## Services Overview
### API-Gateway
- The single entry point for clients
- Routes HTTP requests to appropriate gRPC calls to internal services.
- Has no direct access to database

### User Service (Postgres)
- Register, Login: returns JWT (contains user_id, role).

### Product Service (Postgres)
- **Categories**
  - Create, Update, Delete (admin only)
  - List categories
- **Products**
  - CRUD products (supplier only for create/update/delete)
  - List products by category
- **Product Offers**
  - Supplier create/update/delete **offers** for products
  - List offers by product
  - List offers by supplier

### Order Service (Redis, Postgres)
- **Cart** (Redis)
  - Add/Update/Remove cart items
  - Get cart
  - Clear cart
- **Orders** (Postgres)
  - Checkout (converts cart to order, pulling actual prices from Product Offer)
  - Get order by ID
  - List orders for authenticated user

### Auth && roles
- JWT parsing at API Gateway; context keys have user_id and roles.
- Middlewares:
  - Logger (access logs for each request)
  - Recovery (recovery from panic errors)
  - AuthMiddleware(JWT_SECRET): validates token, sets context
  - RoleMiddleware("admin" | "supplier" | "buyer"): controls role

## Config (examples)
- API Gateway: `ADDR`, `JWT_SECRET`, `USER_SERVICE_ADDR`, `PRODUCT_SERVICE_ADDR`, `ORDER_SERVICE_ADDR`
- Product Service: `DSN`, `SERVICE_ADDR`
- Order Service: `DSN`, `REDIS_URL`, `SERVICE_ADDR`
- User Service: `DSN`, `SERVICE_ADDR`, `JWT_SECRET`

## Docker/Docker-Compose
- Each service has its own **Dockerfile**.
- `docker-compose.yml` defines:
  - `user-service` + `users-db`
  - `product-service` + `products-db`
  - `order-service` + `orders-db` + `orders-redis`
  - `*_migrate` one-off migration containers

## To get started
1. Clone this repository
```bash
git clone https://github.com/your-username/marketplace.git
cd marketplace
```
2. Start all services with docker-compose
```bash
docker-compose up --build
```
3. Access the API-Gateway at
```bash
http://localhost:8000
```
