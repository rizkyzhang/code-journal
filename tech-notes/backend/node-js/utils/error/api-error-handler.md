---
tags:
  - nodejs-util
---
```ts
export class ApiError extends Error {
  public readonly statusMessage: string;
  public readonly isOperational: boolean;

  constructor(
    message: string,
    public readonly statusCode: number
  ) {
    super(message);

    Object.setPrototypeOf(this, new.target.prototype);
    this.name = this.constructor.name;
    this.statusCode = statusCode;
    this.statusMessage = String(STATUS_CODES[this.statusCode]);
    this.isOperational = true;
    Error.captureStackTrace(this);
  }
}

export class ApiBadRequestError extends ApiError {
  constructor(message: string, statusCode = HttpStatusCode.BadRequest) {
    super(message, statusCode);
  }
}
export class ApiUnauthorizedError extends ApiError {
  constructor(
    message = "unauthorized",
    statusCode = HttpStatusCode.Unauthorized
  ) {
    super(message, statusCode);
  }
}
export class ApiForbiddenError extends ApiError {
  constructor(message = "forbidden", statusCode = HttpStatusCode.Forbidden) {
    super(message, statusCode);
  }
}
export class ApiNotFoundError extends ApiError {
  constructor(message = "not found", statusCode = HttpStatusCode.NotFound) {
    super(message, statusCode);
  }
}
export class ApiInternalServerError extends ApiError {
  constructor(
    message = "server error",
    statusCode = HttpStatusCode.InternalServerError
  ) {
    super(message, statusCode);
  }
}
```

```ts
export class ErrorHandler {
  public handleError(error: Error, res?: Response) {
    logger.error("", error);

    if (!this.isOperationalError(error)) {
      res?.status(HttpStatusCode.InternalServerError).json({
        status: STATUS_CODES[HttpStatusCode.InternalServerError],
        error: "server error",
      });

      process.exit(1);
    }

    if (error instanceof ApiError) {
      res?.status(error.statusCode).json({
        status: error.statusMessage,
        error: error.message,
      });

      return;
    }

    if (error instanceof ZodError) {
      const err = fromZodError(error);

      res?.status(HttpStatusCode.BadRequest).json({
        status: STATUS_CODES[HttpStatusCode.BadRequest],
        error: err.message,
      });

      return;
    }

    if (error instanceof PrismaClientKnownRequestError) {
      res?.status(HttpStatusCode.InternalServerError).json({
        status: STATUS_CODES[HttpStatusCode.InternalServerError],
        error: "server error",
      });
    }
  }

  public isOperationalError(error: Error) {
    if (
      error instanceof ApiError ||
      error instanceof ZodError ||
      error instanceof PrismaClientKnownRequestError
    ) {
      return true;
    }

    return false;
  }
}
```