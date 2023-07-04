// Reference: https://www.aleksandrhovhannisyan.com/blog/javascript-promise-tricks/

const getData = (data, successRate = 0.98, maxLatencyMs = 1000) =>
  new Promise((resolve, reject) => {
    const successRoll = Math.random();
    // interval: [0, maxLatencyMs]
    const latency = Math.floor(Math.random() * (maxLatencyMs + 1));

    if (successRoll <= successRate) {
      setTimeout(() => resolve(data), latency);
    } else {
      setTimeout(() => reject(new Error("API failed to return data")), latency);
    }
  });
