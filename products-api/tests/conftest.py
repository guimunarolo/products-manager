import uuid

import pytest
from fastapi.testclient import TestClient

from products_api.main import app


@pytest.fixture
def client():
    with TestClient(app) as client:
        yield client


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
