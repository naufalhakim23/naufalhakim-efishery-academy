# eFishery Warehouse Design

# 1. Summary

Warehouse management system to manage user orders, warehouse inventory, in and out inventory from factory to warehouse and to user.

# 2. Problem & Motivation

By creating such system we can track how much inventory that each warehouse have, when to replenish inventory, when warehouse need to send to customer(user), and for easier sending inventory to user from the nearest warehouse.

# 3. Detailed Design

## User Interface

## Entity Relationship Diagram

- User ERD

(id, username, password(hashed), phone number, address, role (user, head warehouse, officer warehouse), create at, modified at, )

Warehouse ERD

- Inventory in Mall ERD

(inventory_id, product_name, stock(can be seen for user&warehouse), description, price, created_at, modified_at, transactional.in_product, transactional.out_product)

- Transactional User and Warehouse.

(transactional_id, inventory.product, in_product(howmuch in kg)(can be nulled), out_product(howmuch in kg)(can be nulled), user.location, status, created_at, modified_at).

![Untitled](pictures/entity-relationship-diagram.png)

[efishery.users](https://www.notion.so/fc1e55398d014762b914bb721199c94e)
<br/>
![efishery.users](pictures/efishery-users.png)

[efishery.product](https://www.notion.so/6344661651d14967bd82c1c1e7199a9f)
<br/>
![efishery.product](pictures/efishery-products.png)

[efishery.address](https://www.notion.so/1e2441cf21094fe79779b2382300e9a3)
<br/>
![efishery.address](pictures/efishery-address.png)

[efishery.order](https://www.notion.so/e7c53e6f25c5458db6ac41836f5de3de)
<br/>
![efishery.order](pictures/efishery-orders.png)

## API Contract

### **Create User and authentication**

- Method: POST
- Data sent:

```json
{
  {
  "email": "test@gmail.com",
  "password": "test@123",
	"first_name": "Hakim",
	"last_name": "Naufal",
	"phone": "081285402208",
	}
}
```

- Endpoint: /api/user/registration
- Response Body:
  - 201 Created
- Body Payload:

```json
{
  "message": "User created succesfully",
  "user": {
    "auth_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc0YmNhZjUwLWFhNDAtNDc2Yy1hNTM5LTkyMTVmZTc4OGZkNiIsImlhdCI6MTY1NTYwNjcyMn0.gHs7YwXKx7yGSAakxx2_2pO700ShB5ApVjpttbiAy1A",
    "roles": "user",
    "created_at": "2022-06-19T08:18:28.288Z"
  }
}
```

### Login Auth

- Method: POST
- Data sent:

```json
{
  {
  "email": "test@gmail.com",
  "password": "test@123"
	}
}
```

- Endpoint: /api/user/login
- Response Body:
  - 201 Created
- Body Payload:

```json
{
  {
  "message": "User logged in successfully",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc0YmNhZjUwLWFhNDAtNDc2Yy1hNTM5LTkyMTVmZTc4OGZkNiIsImlhdCI6MTY1NTYwNjcyMn0.gHs7YwXKx7yGSAakxx2_2pO700ShB5ApVjpttbiAy1A"
	}
}
```

### Read User, and Warehouse profile for WebApp or NativeApp

- Method: GET
- Endpoint: /api/user/{id}
- Response Body:
  - 200 OK
- Body Payload:

```json
{
  "message": "User retrived succesfully",
  "user": {
    "id": "74bcaf50-aa40-476c-a539-9215fe788fd6",
    "first_name": "John",
    "last_name": "Doe",
    "email": "test@gmail.com",
    "role": "admin",
    "createdAt": "2022-06-19T08:18:28.288Z",
    "updatedAt": "2022-06-19T08:18:28.288Z"
  }
}
```

### Order from User and Warehouse

- Method: POST
- Data sent:

```json
{
  {
  "product": "eFeeder",
  "in_product": "",
	"out_product": "1",
	"location": "Jakarta",

	}
}
```

- Endpoint: /api/transaction/{id}
- Response Body:
  - 201 Created
- Body Payload:

```json
{
  "message": "Transaction succesful.",
  "details": {
    "product": "eFeeder",
    "status": "verified",
    "modified_at": "2022-10-19T12:18:28.288Z"
  }
}
```

# 4. Dependencies

Using Google login

# 5. Milestone

# 6. Deployment

# 7. Drawbacks/Risks/Possible Failure

# 8. Alternative

# 9. Unresolved/Future Possibilities
