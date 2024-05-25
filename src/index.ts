import { app } from "./libs/expres";

const port = process.env.PORT || 8080
app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
})

app.get("/", async (req, res) => {
    res.status(200).json({
        "author": "yuefii",
        "github_author": "https://github.com/yuefii",
        "github_repository": "https://github.com/Yuefii/api-nusantara-kita",
        "documentation": "https://yuefii.my.id/nusantara-kita",
        "version": "1.0.1"
    })
})