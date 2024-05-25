import supertest from "supertest"

import { app } from "../src/libs/expres"
import { logger } from "../src/libs/winston"

describe('GET /api/nusantara/:regency_code/districts', () => {
    it('Should get data districts for a given regency code', async () => {
        const regencyCode = "36.03"
        const response = await supertest(app)
            .get(`/api/nusantara/${regencyCode}/districts`)
        logger.debug(response.body)
        expect(response.status).toBe(200)
        expect(response.body).toHaveProperty("data")
        expect(Array.isArray(response.body.data)).toBe(true)

        const districts = response.body.data;
        districts.forEach(district => {
            expect(district).toHaveProperty("code")
            expect(district).toHaveProperty("regency_code")
            expect(district).toHaveProperty("name")
            expect(typeof district.code).toBe("string")
            expect(typeof district.regency_code).toBe("string")
            expect(typeof district.name).toBe("string")
            expect(district.regency_code).toBe(regencyCode)

            expect(districts).toContainEqual({ code: "36.03.01", name: "Balaraja", regency_code: "36.03" })
            expect(districts).toContainEqual({ code: "36.03.27", name: "Sukamulya", regency_code: "36.03" })
        });
    });

    it('Should reject get and return 404 if regency_code not found', async () => {
        const invalidRegencyCode = "100"
        const response = await supertest(app)
            .get(`/api/nusantara/${invalidRegencyCode}/districts`)
        logger.debug(response.body);
        expect(response.status).toBe(404);
        expect(response.body).toHaveProperty('errors');
        expect(response.body.errors).toBe('regency_code not found');
    });
})
