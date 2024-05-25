import { Utils } from "./utils"
import { logger } from "../libs/winston";
import { ResponseError } from "../error/response-error";
import { NextFunction, Request, Response } from "express";


export class AreaController {
    private utils: Utils

    constructor() {
        this.utils = new Utils()
    }
    async getProvinces(_req: Request, res: Response, next: NextFunction) {
        try {
            const result = await this.utils.readCSV("provinces.csv");
            logger.info("successfully to get provinces.")
            res.status(200).json({ data: result });
        } catch (error) {
            next(error)
        }
    }
    async getRegencies(req: Request, res: Response, next: NextFunction) {
        try {
            const { provinces_code } = req.params
            const result = await this.utils.readCSV("regencies.csv")
            const filtered = result.filter(item => item.province_code === provinces_code);
            if (filtered.length === 0) {
                throw new ResponseError(404, "province_code not found")
            }
            logger.info("successfully to get regencies.")
            res.status(200).json({ data: filtered })
        } catch (error) {
            next(error)
        }
    }
    async getDistricts(req: Request, res: Response, next: NextFunction) {
        try {
            const { regency_code } = req.params
            const result = await this.utils.readCSV("districts.csv")
            const filtered = result.filter(item => item.regency_code === regency_code);
            if (filtered.length === 0) {
                throw new ResponseError(404, "regency_code not found")
            }
            logger.info("successfully to get districts.")
            res.status(200).json({ data: filtered })
        } catch (error) {
            next(error)
        }
    }
    async getVillages(req: Request, res: Response, next: NextFunction) {
        try {
            const { district_code } = req.params
            const result = await this.utils.readCSV("villages.csv")
            const filtered = result.filter(item => item.district_code === district_code);
            if (filtered.length === 0) {
                throw new ResponseError(404, "district_code not found")
            }
            logger.info("successfully to get villages.")
            res.status(200).json({ data: filtered })
        } catch (error) {
            next(error)
        }
    }
}