import supertest from "supertest";

import { app } from "../src/libs/expres"
import { logger } from "../src/libs/winston";

describe('GET /api/nusantara/:district_code/villages', () => {
    it('Should get data villages for a given district code', async () => {
        const districtCode = "36.03.27"
        const response = await supertest(app)
            .get(`/api/nusantara/${districtCode}/villages`)
        logger.debug(response.body)
        expect(response.status).toBe(200)
        expect(response.body).toHaveProperty("data")
        expect(Array.isArray(response.body.data)).toBe(true)

        const villages = response.body.data;
        villages.forEach(village => {
            expect(village).toHaveProperty("code")
            expect(village).toHaveProperty("district_code")
            expect(village).toHaveProperty("name")
            expect(typeof village.code).toBe("string")
            expect(typeof village.district_code).toBe("string")
            expect(typeof village.name).toBe("string")
            expect(village.district_code).toBe(districtCode)

            expect(villages).toContainEqual({ code: "36.03.27.2002", district_code: "36.03.27", name: "Sukamulya" })
            expect(villages).toContainEqual({ code: "36.03.27.2006", district_code: "36.03.27", name: "Merak" })
        });
    });

})