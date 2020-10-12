import pytest
from fastapi import status

from products_api.db import database
from products_api.managers import ProductManager

pytestmark = pytest.mark.asyncio


async def test_list_products_without_stored_products(client, db_connection, product_data):
    response = await client.get("/products/")
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert response_data == []


@database.transaction(force_rollback=True)
async def test_list_products_with_stored_products(client, db_connection, product_data):
    del product_data["id"]
    await ProductManager.create(product_data)

    response = await client.get("/products/")
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert len(response_data) == 1
    assert response_data[0]["title"] == product_data["title"]
    assert response_data[0]["description"] == product_data["description"]
    assert response_data[0]["price_in_cents"] == product_data["price_in_cents"]
    assert response_data[0]["discount"] == {}
