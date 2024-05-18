import cors from 'cors'
import morgan from 'morgan'
import express, { Request, Response } from 'express'
import { router } from './api/routes';

const app = express()
app.use(cors());
app.use(express.json());
app.use(morgan('dev'));
app.use(router)

const port = process.env.PORT || 8080
app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
})

app.get("/", async (req: Request, res: Response) => {
    res.status(200).json({
        "author": "yuefii",
        "github_author": "https://github.com/yuefii",
        "version": "1.0.0"
    })
})