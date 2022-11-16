```go
type Todo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	IsCompleted bool               `bson:"is_completed" json:"is_completed"`
}
```

## Connect

```go
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
if err != nil {
	fmt.Printf("mongo.Connect error: %s", err)
	return
}
err = client.Ping(ctx, readpref.PrimaryPreferred())
if err != nil {
	fmt.Printf("client.Ping error: %s", err)
	return
}
defer func() {
	if err = client.Disconnect(ctx); err != nil {
		fmt.Printf("client.Disconnect error: %s", err)
		return
	}

	fmt.Println("mongodb disconnected")
}()
```

## Define database and collection

```go
db := client.Database("learn")
todoCol := db.Collection("todo")
```

## Find one

```go
var todo Todo
filter := bson.M{"is_completed": false}

err = todoCol.FindOne(ctx, filter).Decode(&todo)
if err != nil {
	if err == mongo.ErrNoDocuments {
	  fmt.Print("todo not found")
	} else {
    fmt.Printf("FindOne error: %s", err)
  }

  return
}
```

## Find with cur.All and cur.Next

```go
var todos []Todo

filter := bson.D{
	{Key: "$or",
		Value: bson.A{
			bson.D{{Key: "name", Value: "Todo 5"}},
			bson.D{{Key: "name", Value: "Todo 7"}},
			bson.D{{Key: "name", Value: "Todo 9"}},
		},
	},
}

cur, err := todoCol.Find(ctx, filter)
if err != nil {
	fmt.Printf("Find error: %s", err)
	return
}
defer func() {
	err := cur.Close(ctx)
	if err != nil {
		fmt.Printf("cur.Close error: %+v", err)
		return
	}
}()

// err = cur.All(ctx, &todos)
// if err != nil {
// 	fmt.Printf("cur.All error: %s", err)
// 	return
// }

for cur.Next(ctx) {
	var todo Todo

	err := cur.Decode(&todo)
	if err != nil {
		fmt.Printf("cur.Decode error: %s", err)
		return
	}

	todos = append(todos, todo)
}
```

## Insert one

```go
res, err := todoCol.InsertOne(ctx, &Todo{
	ID:          primitive.NewObjectID(),
	Name:        fmt.Sprintf("Todo %d", v),
	IsCompleted: false,
})
if err != nil {
	fmt.Printf("InsertOne error: %s", err)
	return
}

fmt.Println(res.InsertedID)
```

## Insert many

```go
_payload := []Todo{
	{
		ID:          primitive.NewObjectID(),
		Name:        "Test 8",
		IsCompleted: true,
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Test 9",
		IsCompleted: true,
	},
	{
		ID:          primitive.NewObjectID(),
		Name:        "Test 10",
		IsCompleted: true,
	},
}

var payload []interface{}

for _, todo := range _payload {
	payload = append(payload, todo)
}

res, err := todoCol.InsertMany(ctx, finalPayload)
if err != nil {
	fmt.Printf("InsertMany error: %s", err)
	return
}

fmt.Printf("inserted ids: %+v", res.InsertedIDs)
```

## Update one

```go
type TodoPayloadDTO struct {
	IsCompleted *bool  `bson:"is_completed,omitempty" json:"is_completed"`
	Name        string `bson:"name,omitempty" json:"name"`
}

filter := bson.M{"name": "Todo 1"}

payload := TodoPayloadDTO{
	Name: "Todo One",
}
update := bson.D{{Key: "$set", Value: payload}}

res, err := todoCol.UpdateOne(ctx, filter, update)
if err != nil {
	fmt.Printf("UpdateOne error: %s", err)
	return
}
```

## Update many

```go
type TodoPayloadDTO struct {
	IsCompleted *bool  `bson:"is_completed,omitempty" json:"is_completed"`
	Name        string `bson:"name,omitempty" json:"name"`
}

filter := bson.M{"is_completed": true}

payload := TodoPayloadDTO{
	IsCompleted: func(val bool) *bool {
		p := val
		return &p
	}(false),
}
update := bson.D{{Key: "$set", Value: payload}}

res, err := todoCol.UpdateMany(ctx, filter, update)
if err != nil {
	fmt.Printf("UpdateMany error: %s", err)
	return
}
```

## Delete one

```go
filter := bson.M{"name": "Todo One"}

res, err := todoCol.DeleteOne(ctx, filter)
if err != nil {
	fmt.Printf("DeleteOne error: %s", err)
	return
}

fmt.Println(res.DeletedCount)
```

## Delete many

```go
filter := bson.D{{
	Key: "$or",
	Value: bson.A{
		bson.D{{Key: "name", Value: "Todo 8"}},
		bson.D{{Key: "name", Value: "Todo 9"}},
		bson.D{{Key: "name", Value: "Todo 10"}},
	},
}}

res, err := todoCol.DeleteMany(ctx, filter)
if err != nil {
	fmt.Printf("delete many error: %s", err)
  return
}
if res.DeletedCount == 0 {
	fmt.Println("no document is deleted")
}
```
