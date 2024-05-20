import express from "express"
import { AreaController } from "./controllers";

const area = new AreaController()
export const router = express.Router();
router.get('/api/nusantara/provinces', area.getProvinces.bind(area));
router.get('/api/nusantara/:provinces_code/regencies', area.getRegencies.bind(area));
router.get('/api/nusantara/:regency_code/districts', area.getDistricts.bind(area));
router.get('/api/nusantara/:district_code/villages', area.getVillages.bind(area));