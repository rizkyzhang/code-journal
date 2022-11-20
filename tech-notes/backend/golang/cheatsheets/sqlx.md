## Connect

### Connect & ping

```go
db, err := sqlx.Connect("pgx", "postgres://root:root@localhost:5432/test")
if err != nil {
	fmt.Printf("sqlx.Connect: %s", err)
	return
}
defer func() {
	db.Close()
	fmt.Println("db closed")
}()
```

## Query

### Query row with extension to scan result into struct or map

```go
var name string

err := db.QueryRowxContext(ctx, "SELECT name FROM product LIMIT 1").Scan(&name)
if err != nil {
	fmt.Printf("db.QueryRowxContext: %s", err)
	return
}
```

### Query rows with named queries provided by struct

```go
var filter = struct{
  Price uint32 `db:"price"`
  Limit uint8 `db:"limit"`
}{
	Price: 125000,
	Limit: 25,
}

var products []Product

rows, err := db.NamedQueryContext(ctx, `
SELECT * FROM products
WHERE price > :price
LIMIT :limit
`, &filter)
if err != nil {
	fmt.Printf("db.NamedQueryContext: %s", err)
	return
}
if rows.Err() != nil {
	fmt.Printf("db.NamedQueryContext rows.Err: %s", err)
	return
}
for rows.Next() {
	var product Product
	err := rows.StructScan(&product)
  if err != nil {
	  fmt.Printf("rows.StructScan: %s", err)
	  return
  }

  products = append(products, product)
}
```

### Query row auto scan result into slice or struct or map

```go
var product Product

err = db.GetContext(ctx, &products, "SELECT * FROM products LIMIT 1")
	if err != nil {
		fmt.Printf("db.GetContext: %s", err)
		return
	}
```

### Query row get last inserted row id

```go
query, args, err := db.BindNamed(`
	INSERT INTO products (uid, name, slug, sku, description, images, created_at, updated_at)
	VALUES (:uid, :name, :slug, :sku, :description, :images, :created_at, :updated_at)
	RETURNING id;
	`, productPayload)
if err != nil {
	return err
}

var productID int64
err = db.GetContext(ctx, &productID, query, args...)
if err != nil {
	return err
}
```

### Query rows without iterating `rows.Next` and auto scan result into struct or map

```go
var products []Product

err = db.SelectContext(ctx, &products, "SELECT * FROM products WHERE price = $1", 25500)
	if err != nil {
		fmt.Printf("db.SelectContext: %s", err)
		return
	}
```

## Exec

### Named exec

```go
err := db.NamedExecContext(ctx, `
	INSERT INTO products (id, name, slug, description, images, created_at, updated_at)
	VALUES (:id, :name, :slug, :description, :images, :created_at, :updated_at)
	`, payload)
if err != nil {
	fmt.Printf("db.NamedExec: %s ", err)
	return
}
```
