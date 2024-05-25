import supertest from "supertest"

import { app } from "../src/libs/expres"
import { logger } from "../src/libs/winston"

describe('GET /api/nusantara/:provinces_code/regencies', () => {
    it('Should get data regencies for a given province code', async () => {
        const provinceCode = "36"
        const response = await supertest(app)
            .get(`/api/nusantara/${provinceCode}/regencies`)
        logger.debug(response.body)
        expect(response.status).toBe(200)
        expect(response.body).toHaveProperty("data")
        expect(Array.isArray(response.body.data)).toBe(true)

        const regencies = response.body.data
        regencies.forEach(regency => {
            expect(regency).toHaveProperty("code")
            expect(regency).toHaveProperty("province_code")
            expect(regency).toHaveProperty("name")
            expect(typeof regency.code).toBe("string")
            expect(typeof regency.province_code).toBe("string")
            expect(typeof regency.name).toBe("string")
            expect(regency.province_code).toBe(provinceCode)

            expect(regencies).toContainEqual({ code: "36.03", province_code: "36", name: "KABUPATEN TANGERANG" })
            expect(regencies).toContainEqual({ code: "36.71", province_code: "36", name: "KOTA TANGERANG" })
        });
    });

    it('Should reject get and return 404 if province_code not found', async () => {
        const invalidProvinceCode = "100"
        const response = await supertest(app)
            .get(`/api/nusantara/${invalidProvinceCode}/regencies`)
        logger.debug(response.body)
        expect(response.status).toBe(404)
        expect(response.body).toHaveProperty('errors')
        expect(response.body.errors).toBe('province_code not found')
    });
})
