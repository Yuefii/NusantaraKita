import { Utils } from "./utils"
import { Request, Response } from "express";


export class AreaController {
    private utils: Utils

    constructor() {
        this.utils = new Utils()
    }
    async getProvinces(_req: Request, res: Response) {
        try {
            const result = await this.utils.readCSV("provinces.csv");
            res.json({ data: result });
        } catch (error) {
            console.error(error)
            res.status(500).json({
                statusCode: 500,
                error: 'Internal Server Error'
            })
        }
    }
    async getRegencies(req: Request, res: Response) {
        try {
            const { provinces_code } = req.params
            const result = await this.utils.readCSV("regencies.csv")
            const filtered = result.filter(item => item.province_code === provinces_code);
            if (filtered.length === 0) {
                return res.status(404).json({
                    statusCode: 404,
                    error: 'Not Found',
                    message: 'province_code not found'
                });
            }
            res.json({ data: filtered })
        } catch (error) {
            console.error(error)
            res.status(500).json({
                statusCode: 500,
                error: 'Internal Server Error'
            })
        }
    }
    async getDistricts(req: Request, res: Response) {
        try {
            const { regency_code } = req.params
            const result = await this.utils.readCSV("districts.csv")
            const filtered = result.filter(item => item.regency_code === regency_code);
            if (filtered.length === 0) {
                return res.status(404).json({
                    statusCode: 404,
                    error: 'Not Found',
                    message: 'regency_code not found'
                });
            }
            res.json({ data: filtered })
        } catch (error) {
            console.error(error)
            res.status(500).json({
                statusCode: 500,
                error: 'Internal Server Error'
            })
        }
    }
    async getVillages(req: Request, res: Response) {
        try {
            const { district_code } = req.params
            const result = await this.utils.readCSV("villages.csv")
            const filtered = result.filter(item => item.district_code === district_code);
            if (filtered.length === 0) {
                return res.status(404).json({
                    statusCode: 404,
                    error: 'Not Found',
                    message: 'district_code not found'
                });
            }
            res.json({ data: filtered })
        } catch (error) {
            console.error(error)
            res.status(500).json({
                statusCode: 500,
                error: 'Internal Server Error'
            })
        }
    }
}