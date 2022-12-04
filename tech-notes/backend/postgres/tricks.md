## Get total rows

```sql
ANALYZE;
SELECT reltuples::bigint AS estimate
FROM   pg_class
WHERE  oid = to_regclass('public.table_name');
```
