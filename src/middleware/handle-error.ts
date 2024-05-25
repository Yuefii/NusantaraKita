import { logger } from "../libs/winston";
import { ResponseError } from "../error/response-error";
import { Response, Request, NextFunction } from "express";

export const handleError = async (error: Error, req: Request, res: Response, next: NextFunction) => {
    if (error instanceof ResponseError) {
        logger.error("error invalid get by ID"),
            res.status(error.status).json({
                statusCode: error.status,
                errors: error.message
            });
    } else {
        logger.error(error)
        res.status(500).json({
            statusCode: 500,
            error: 'Internal Server Error'
        })
    }
}