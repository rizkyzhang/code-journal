Sometime after inserting a row, we might need to get the id usually for creating referenced table with foreign key of id. We can try to use `db.NamedExec` or `db.Exec` to do the operation since it will return a `sql.Result` interface which have `LastInsertId()` method inside it. The problem since we are using PostgreSQL with `pgx` driver, this method is not supported as stated in the documentation: **In MySQL, for instance, LastInsertId() will be available on inserts with an auto-increment key, but in PostgreSQL, this information can only be retrieved from a normal row cursor by using the RETURNING clause.**.

The workaround is to take advantage of `RETURNING` clause that returns the updated row and combine it with `db.GetContext` or `db.QueryRowx` to get the query result, I will use `db.GetContext` here.

For example, we have a `productPayload` struct that we are going to insert.

```go
type productPayload struct {
    UID         string                  `db:"uid" json:"uid"`
	Name        string                  `db:"name" json:"name"`
	Slug        string                  `db:"slug" json:"slug"`
	SKU         string                  `db:"sku" json:"sku"`
	Description string                  `db:"description" json:"description"`
	Image       string                  `db:"image" json:"image"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
```

And here is the full code:

```go
query, args, err := db.BindNamed(`
	INSERT INTO products (uid, name, slug, sku, description, image, created_at, updated_at)
	VALUES (:uid, :name, :slug, :sku, :description, :image, :created_at, :updated_at)
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

The usage of `db.BindNamed` is optional, I use it because I want to convert my `productPayload` struct into slice of positioned arguments, so I don't need to manually pass it to the `db.GetContext`.

I hope this article is useful for you, if it is please share it with your friends, thank you so much for reading to the end and see you in the next article!
