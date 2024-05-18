import csv from "csvtojson"
import path from "path";

export class Utils {
    async readCSV(fileName: string) {
        const filePath = path.join(__dirname, '../../data', fileName);
        return await csv().fromFile(filePath);
    }
}