import databases

from .configs import settings


database = databases.Database(settings.DATABASE_URL)


async def db_connect():
    await database.connect()


async def db_disconnect():
    await database.disconnect()
