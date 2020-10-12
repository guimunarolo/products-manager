import uuid

import pytest
from httpx import AsyncClient

from products_api.db import db_connect, db_disconnect
from products_api.main import app


@pytest.fixture
async def client():
    async with AsyncClient(app=app, base_url="http://localhost:8000") as client:
        yield client


@pytest.fixture
async def db_connection():
    await db_connect()
    yield None
    await db_disconnect()


@pytest.fixture
def user_data():
    return {
        "id": str(uuid.uuid4()),
        "first_name": "Jane",
        "last_name": "Doe",
        "date_of_birth": "1999-09-09",
    }


@pytest.fixture
def product_data():
    return {
        "id": str(uuid.uuid4()),
        "price_in_cents": 120000,
        "title": "Some Product",
        "description": "Some product description here.",
    }
