# inventory-tracking-crud

## Description

Simple CRUD for inventory tracking web application.

Inventory has the following fields:

- **Id**: inventory id
- **Name**: inventory name
- **Quantity**: quantity of inventory
- **Deleted**: flag showing if inventory is deleted
- **Comment**: deletion comment

### API

- GET `/api/inventory?show_deleted={true|false}`: Get all inventories, optional query parameter allows showing deleted inventories
- GET `/api/inventory/{id}`: Get inventory by id
- POST `/api/inventory`: Create new inventory
- PUT `/api/inventory/{id}`: Update existing inventory from provided json fields
- DELETE `/api/inventory/{id}?comment={deletion_comment}`: Delete inventory and optionally add deletion comment
- POST `/api/inventory/{id}/undelete`: Undelete (restore) deleted inventory

## Usage

To run locally, set `CONNECTION_STRING` environment variable as per your local database connection string.

Example:
```
postgresql://username:password@host:port/dbname[?paramspec]
```

```sh
> export CONNECTION_STRING="{connection_string}"
> go run .
```

The host for remote run: https://inventory-tracking-crud.alyonaantonova.repl.co/.

The corresponding Replit link: https://replit.com/@AlyonaAntonova/inventory-tracking-crud?v=1
