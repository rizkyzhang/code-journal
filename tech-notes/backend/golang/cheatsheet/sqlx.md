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

### Convert from sql to sqlx

```go
var mySqlDb *sql.DB

mySqlxDb := sqlx.NewDb(mySqlDb, "pgx") // returns *sqlx.DB
```

### Convert to sql

```go
var mySqlxDb *sqlx.DB

var mySqlDb := mySqlxDb.DB
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

## Scan into nested field

The easiest way is to embed the objects, but I do it in a temporary object instead of making the original messier.

You then add a function to convert your temp (flat) object into your real (nested) one.

```go
type Customer struct {
  Id      int     `json:"id" db:"id"`
	Name    string  `json:"name" db:"name"`
	Address Address `json:"address"`
}

type Address struct {
  Street string `json:"street" db:"street"`
  City   string `json:"city" db:"city"`
}

type tempCustomer struct {
  Customer
  Address
}

func (c *tempCustomer) ToCustomer() Customer {
  customer := c.Customer
  customer.Address = c.Address
  return customer
}
```

Now you can scan into tempCustomer and simply call tempCustomer.ToCustomer before you return. This keeps your JSON clean and doesn't require a custom scan function.

Reference: https://stackoverflow.com/a/72322421

## In query

```go
var carts []model.Cart

query, args, _ := sqlx.In("SELECT * FROM carts WHERE id IN (?)", IDS)
query = b.db.Rebind(query)

_ = b.db.SelectContext(ctx, &carts, query, args...)
```
