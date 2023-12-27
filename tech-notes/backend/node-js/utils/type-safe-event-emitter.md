---
tags:
  - nodejs-util
---
```ts
/* eslint-disable @typescript-eslint/no-explicit-any */
import EventEmitter from "events";

export type LocalEventTypes = {
  "task.created": [
    {
      uid: string;
      title: string;
      content?: string;
    },
  ];
  "task.updated": [
    {
      uid: string;
      title: string;
      content?: string;
    },
  ];
};

export class TypedEventEmitter<TEvents extends Record<string, any>> {
  private emitter = new EventEmitter();

  emit<TEventName extends keyof TEvents & string>(
    eventName: TEventName,
    ...eventArg: TEvents[TEventName]
  ) {
    this.emitter.emit(eventName, ...(eventArg as []));
  }

  on<TEventName extends keyof TEvents & string>(
    eventName: TEventName,
    handler: (...eventArg: TEvents[TEventName]) => void
  ) {
    this.emitter.on(eventName, handler as any);
  }

  off<TEventName extends keyof TEvents & string>(
    eventName: TEventName,
    handler: (...eventArg: TEvents[TEventName]) => void
  ) {
    this.emitter.off(eventName, handler as any);
  }
}

// Usage
const eventBroker = new TypedEventEmitter<LocalEventTypes>(); 

eventBroker.on('task.created', (payloadJson) => { console.log(payloadJson.uid) }) 
eventBroker.emit('task.created', {
	uid: "123"
	title: "Test"
	content: "Test"
})
```

## Reference

https://blog.makerx.com.au/a-type-safe-event-emitter-in-node-js/