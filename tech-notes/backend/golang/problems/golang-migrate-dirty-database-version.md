## Problem

Golang migrate error `dirty database version` when some error encountered when migrate up or down, even if the problem is already fixed, migration still not working.

## Causes

For example your migrations directory look like this:

```
migrations
  20221009124227_create_product_table.up.sql
  20221009124227_create_product_table.down.sql
  20221009130110_add_price_to_product_table.up.sql
  20221009130110_add_price_to_product_table.down.sql
```

You try to migrate up and got this error:

`error: migration failed: syntax error at or near "IN" (column 18) in line 2: ALTER TABLE products ADD COLUMN price IN NOT NULL; (details: pq: syntax error at or near "IN")`

Easy, you have a typo, replace IN with INT and save the sql file.

Run migrate up again, but you get this error instead `error: Dirty database version 20221009130110. Fix and force version`.

The reason why this happen can be found on the golang-migrate docs chapter below:

In case you run a migration that contained an error, migrate will not let you run other migrations on the same database. You will see an error like Dirty database version 1. Fix and force version, even when you fix the erred migration. This means your database was marked as 'dirty'. You need to investigate the migration error - was your migration applied partially, or was it not applied at all? Once you know, you should force your database to a version reflecting it's real state. You can do so with force command:

`migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION`
Once you force the version and your migration was fixed, your database is 'clean' again and you can proceed with your migrations.

## Solution

The version shown in the error will be always the version you are going to migrate to,
but **it is not the version you are going to force**, **instead it is the current version**.

For example you try to migrate up to version `20221009130110`, then you should force version `20221009124227`, run `migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force 20221009124227`.
If you try to migrate down to version `20221009124227`, then you should force version `20221009130110`, run `migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force 20221009130110`.

After version forcing is done, you can run migrate up and down again.

## Reference

https://github.com/golang-migrate/migrate/blob/master/GETTING_STARTED.md#forcing-your-database-version
