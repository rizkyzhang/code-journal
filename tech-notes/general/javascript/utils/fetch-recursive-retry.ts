const fetchRecursive = async <T>(
  url: string,
  method: "get" | "post" | "put" | "del" = "get",
  body: any = {},
  retryLimit = Number.MAX_VALUE,
  retryCount = 0 // this param is only used internally to track retry count
): Promise<T> => {
  return fetch(url, {
    method,
    body: method != "get" ? JSON.stringify(body) : null,
  }).then((res) => {
    if (res.status >= 300 && retryCount === retryLimit) {
      throw new Error("retry limit reached");
    }

    if (res.status >= 300 && retryCount < retryLimit) {
      return fetchRecursive(url, method, body, retryLimit, retryCount + 1);
    }

    const data = res.json() as Promise<T>;
    return data;
  });
};

interface Product {
  name: string;
  price: number;
}

const getProduct = async () => {
  try {
    const product = await fetchRecursive<Product>(
      "https://httpbin.org/status/400",
      "get",
      null,
      10
    );
    return product;
  } catch (err) {
    if (err instanceof Error) {
      // handle error
    }
  }
};
