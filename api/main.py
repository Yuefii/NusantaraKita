import uvicorn
from fastapi import FastAPI

from routers.V2 import router as version_2

app = FastAPI()

app.include_router(version_2)

if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8080, reload=True)
