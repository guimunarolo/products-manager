
from fastapi import FastAPI

from .db import db_connect, db_disconnect
from .views import router as views_router


app = FastAPI()
app.include_router(views_router)


@app.on_event("startup")
async def startup():
    await db_connect()    


@app.on_event("shutdown")
async def shutdown():
    await db_disconnect()
