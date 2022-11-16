## Bindvar

The ? query placeholders, called bindvars internally, are important; you should always use these to send values to the database, as they will prevent SQL injection attacks. database/sql does not attempt any validation on the query text; it is sent to the server as is, along with the encoded parameters. Unless drivers implement a special interface, the query is prepared on the server first before execution. Bindvars are therefore database specific:

MySQL uses the ? variant shown above
PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
SQLite accepts both ? and $1 syntax
Oracle uses a :name syntax
Other databases may vary. You can use the `sqlx.DB.Rebind(string)` string function with the ? bindvar syntax to get a query which is suitable for execution on your current database type.

A common misconception with bindvars is that they are used for interpolation. They are only for parameterization, and are not allowed to change the structure of an SQL statement. For instance, using bindvars to try and parameterize column or table names will not work:

// doesn't work
`db.Query("SELECT * FROM ?", "mytable")`

// also doesn't work
`db.Query("SELECT ?, ? FROM people", "name", "location")`
