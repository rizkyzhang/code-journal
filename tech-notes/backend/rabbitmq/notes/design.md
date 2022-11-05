Different types of object events are so natural to match different types of messages to be published, but it is not really necessary sometimes. What if we abstract all the 3 types of events as a “write” event, whose sub-types are “created”, “updated” and “deleted”?

| object | event | sub-type |
| ------ | ----- | -------- |
| user   | write | created  |
| user   | write | updated  |
| user   | write | deleted  |

Solution 1

The simplest solution to support this is we could only design a “user.write” queue, and publish all user write event messages to this queue directly via the global default exchange. When publishing to a queue directly, the biggest limitation is it assumes that only one app subscribes to this type of messages. Multiple instances of one app subscribing to this queue is also fine.

| queue      | app  |
| ---------- | ---- |
| user.write | app1 |

Solution 2

The simplest solution could not work when there is a second app (having different processing logic) want to subscribe to any messages published to the queue. When there are multiple apps subscribing, we at least need one “fanout” type exchange with bindings to multiple queues. So that messages are published to the exchange, and the exchange duplicates the messages to each of the queues. Each queue represents the processing job of each different app.

| queue           | subscriber |
| --------------- | ---------- |
| user.write.app1 | app1       |
| user.write.app2 | app2       |

| exchange   | type   | binding_queue   |
| ---------- | ------ | --------------- |
| user.write | fanout | user.write.app1 |
| user.write | fanout | user.write.app2 |

This second solution works fine if each subscriber does care about and want to handle all the sub-types of “user.write” events or at least to expose all these sub-type events to each subscribers is not a problem. For instance, if the subscriber app is for simply keeping the transaction log; or although the subscriber handles only user.created, it is ok to let it know about when user.updated or user.deleted happens. It becomes less elegant when some subscribers are from external of your organization, and you only want to notify them about some specific sub-type events. For instance, if app2 only wants to handle user.created and it should not have the knowledge of user.updated or user.deleted at all.

Solution 3

To solve the issue above, we have to extract “user.created” concept from “user.write”. The “topic” type of exchange could help. When publishing the messages, let’s use user.created/user.updated/user.deleted as routing keys, so that we could set the binding key of “user.write.app1” queue be “user.\*” and the binding key of “user.created.app2” queue be “user.created”.

| queue             | subscriber |
| ----------------- | ---------- |
| user.write.app1   | app1       |
| user.created.app2 | app2       |

| exchange   | type  | binding_queue     | binding_key  |
| ---------- | ----- | ----------------- | ------------ |
| user.write | topic | user.write.app1   | user.\*      |
| user.write | topic | user.created.app2 | user.created |

Solution 4

The “topic” exchange type is more flexible in case potentially there will be more event sub-types. But if you clearly know the exact number of events, you could also use the “direct” exchange type instead for better performance.

| queue             | subscriber |
| ----------------- | ---------- |
| user.write.app1   | app1       |
| user.created.app2 | app2       |

| exchange   | type   | binding_queue     | binding_key  |
| ---------- | ------ | ----------------- | ------------ |
| user.write | direct | user.write.app1   | user.created |
| user.write | direct | user.write.app1   | user.updated |
| user.write | direct | user.write.app1   | user.deleted |
| user.write | direct | user.created.app2 | user.created |

Come back to the “one exchange, or many?” question. So far, all the solutions use only one exchange. Works fine, nothing wrong. Then, when might we need multiple exchanges? There is a slight performance drop if a "topic" exchange has too many bindings. If performance difference of too many bindings on “topic exchange” really becomes an issue, of course you could use more “direct” exchanges to reduce number of “topic” exchange bindings for better performance. But, here I want to focus more on function limitations of “one exchange” solutions.

Solution 5

One case we might naturally consider multiple exchanges is for different groups or dimensions of events. For instance, besides the created, updated and deleted events mentioned above, if we have another group of events: login and logout - a group of events describing “user behaviors” rather than “data write”. Coz different group of events might need completely different routing strategies and routing key & queue naming conventions, it is so that natural to have a separate user.behavior exchange.

| queue              | subscriber |
| ------------------ | ---------- |
| user.write.app1    | app1       |
| user.created.app2  | app2       |
| user.behavior.app3 | app3       |

| exchange      | type  | binding_queue      | binding_key  |
| ------------- | ----- | ------------------ | ------------ |
| user.write    | topic | user.write.app1    | user.\*      |
| user.write    | topic | user.created.app2  | user.created |
| user.behavior | topic | user.behavior.app3 | user.\*      |

Other Solutions

There are other cases when we might need multiple exchanges for one object type. For instance, if you want to set different permissions on exchanges (e.g. only selected events of one object type are allowed to be published to one exchange from external apps, while the other exchange accepts any the events from internal apps). For another instance, if you want to use different exchanges suffixed with a version number to support different versions of routing strategies of same group of events. For another another instance, you might want to define some “internal exchanges” for exchange-to-exchange bindings, which could manage routing rules in a layered way.

In summary, still, “the final solution depends on your system needs”, but with all the solution examples above, and with the background considerations, I hope it could at least get one thinking in the right directions.

Reference: https://stackoverflow.com/a/35591738
