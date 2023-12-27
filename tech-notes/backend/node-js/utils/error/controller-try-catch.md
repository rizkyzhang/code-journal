---
tags:
  - nodejs-util
---
```ts
import { NextFunction, Request, Response } from "express";

export const controllerTryCatch =
  (
    controller: (
      req: Request,
      res: Response,
      next: NextFunction
    ) => Promise<void>
  ) =>
  async (req: Request, res: Response, next: NextFunction) => {
    try {
      await controller(req, res, next);
    } catch (error) {
      const errorHandler = new ErrorHandler();

      if (error instanceof Error) {
        return errorHandler.handleError(error, res);
      }
    }
  };
```