Sometimes when performing a MongoDB query with a long `$in` list, you might want to get return documents in the same order as the elements of the `$in` array are in.

- Actual (by default mongodb sort by insertion order)

```
$in =  ["a", "c", "b"]
result = ["a", "b", "c"]
```

- Expected

```
$in =  ["a", "c", "b"]
result = ["a", "c", "b"]
```

```go
func (b *baseProductRepository) consumeEventCursor(ctx context.Context, cur *mongo.Cursor) ([]*domain.Product, error) {
  var res []*domain.Product
	for cur.Next(ctx) {
		var p domain.Product
		if err := cur.Decode(&p); err != nil {
			return res, err
		}
		res = append(res, &p)
	}

	return res, nil
}

func (b *baseProductRepository) ListByIDS(ctx context.Context, ids []primitive.ObjectID) ([]*domain.Product, error) {
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: bson.D{
			{Key: "$in", Value: ids},
		}}}}},
		{{Key: "$addFields", Value: bson.D{{Key: "__order", Value: bson.D{{Key: "$indexOfArray", Value: bson.A{ids, "$_id"}}}}}}},
		{{Key: "$sort", Value: bson.D{{Key: "__order", Value: 1}}}},
	}

	cur, err := b.col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	defer func() { _ = cur.Close(ctx) }()
	products, err := b.consumeEventCursor(ctx, cur)
	if err != nil {
		return nil, err
	}

	return products, nil
}
```

Reference: http://www.kamsky.org/stupid-tricks-with-mongodb/using-34-aggregation-to-return-documents-in-same-order-as-in-expression
