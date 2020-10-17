from unittest import mock\

import pytest
from fastapi import status

from products_api.db import database
from products_api.managers import ProductManager

pytestmark = pytest.mark.asyncio


async def test_list_products_without_stored_products(client, db_connection, product_data):
    response = await client.get("/product")
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert response_data == []


@database.transaction(force_rollback=True)
async def test_list_products_with_stored_products(client, db_connection, product_data):
    del product_data["id"]
    await ProductManager.create(product_data)

    response = await client.get("/product")
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert len(response_data) == 1
    assert response_data[0]["title"] == product_data["title"]
    assert response_data[0]["description"] == product_data["description"]
    assert response_data[0]["price_in_cents"] == product_data["price_in_cents"]
    assert response_data[0]["discount"] == {}


@database.transaction(force_rollback=True)
@mock.patch("products_api.managers.CalculatorClient.get_product_discount")
async def test_list_products_with_given_used_id(
    get_product_discount_mock, client, db_connection, product_data, user_id
):
    del product_data["id"]
    await ProductManager.create(product_data)
    
    discount = {"pct": 0.1, "value_in_cents": 100}
    get_product_discount_mock.return_value = discount
    response = await client.get("/product", headers={"X-USER-ID": user_id})
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert len(response_data) == 1
    assert response_data[0]["title"] == product_data["title"]
    assert response_data[0]["description"] == product_data["description"]
    assert response_data[0]["price_in_cents"] == product_data["price_in_cents"]
    assert response_data[0]["discount"] == discount
    get_product_discount_mock.assert_called_once_with(
        user_id=user_id, product_id=response_data[0]["id"]
    )
