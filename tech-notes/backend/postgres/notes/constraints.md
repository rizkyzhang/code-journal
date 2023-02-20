## Primary key

- A primary key is a column or a group of columns used to identify a row uniquely in a table
- Primary key is a combination NOT NULL and UNIQUE constraint
- A table can have one and only one primary key
- When you add a primary key to a table, PostgreSQL creates a unique B-tree index on the column or a group of columns used to define the primary key
- If primary key name is not specified, it will be auto-generated in this format `table_name_pkey`

## Foreign key

- A foreign key is a column or a group of columns in a table that reference the primary key of another table
- The table that contains the foreign key is called the referencing table or child table. And the table referenced by the foreign key is called the referenced table or parent table
- A table can have multiple foreign keys depending on its relationships with other tables
- The delete and update actions determine the behaviors when the primary key in the parent table is deleted and updated. Since the primary key is rarely updated, the `ON UPDATE` action is not often used in practice. We’ll focus on the `ON DELETE` action
- If foreign key name is not specified, it will be auto-generated in this format `table_name_column_name_fkey`
- PostgreSQL supports the following actions:
  - `SET NULL`
  - `SET DEFAULT`
  - `RESTRICT`
  - `NO ACTION`
  - `CASCADE`
- The `ON DELETE SET NULL` automatically sets NULL to the foreign key columns in the referencing rows of the child table when the referenced rows in the parent table are deleted
- The `ON DELETE CASCADE` automatically deletes all the referencing rows in the child table when the referenced rows in the parent table are deleted. In practice, the `ON DELETE CASCADE` is the most commonly used option

## CHECK

- A `CHECK` constraint is a kind of constraint that allows you to specify if values in a column must meet a specific requirement
- The `CHECK` constraint uses a Boolean expression to evaluate the values before they are inserted or updated to the column
- If the values pass the check, PostgreSQL will insert or update these values to the column. Otherwise, PostgreSQL will reject the changes and issue a constraint violation error

## UNIQUE

- When a `UNIQUE` constraint is in place, every time you insert a new row, it checks if the value is already in the table. It rejects the change and issues an error if the value already exists. The same process is carried out for updating existing data
- When you add a `UNIQUE` constraint to a column or a group of columns, PostgreSQL will automatically create a unique index on the column or the group of columns
