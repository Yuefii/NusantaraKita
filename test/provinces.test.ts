import supertest from "supertest";

import { app } from "../src/libs/expres"
import { logger } from "../src/libs/winston";

describe('GET /api/nusantara/provinces', () => {
    it('Should get data provinces', async () => {
        const response = await supertest(app)
            .get("/api/nusantara/provinces")
        logger.debug(response.body)
        expect(response.status).toBe(200)
        expect(response.body).toHaveProperty('data');
        expect(Array.isArray(response.body.data)).toBe(true);

        const provinces = response.body.data;
        provinces.forEach(province => {
            expect(province).toHaveProperty('code');
            expect(province).toHaveProperty('name');
            expect(typeof province.code).toBe('string');
            expect(typeof province.name).toBe('string');
        });

        expect(provinces).toContainEqual({ code: "11", name: "ACEH" });
        expect(provinces).toContainEqual({ code: "12", name: "SUMATERA UTARA" });
    });
})
