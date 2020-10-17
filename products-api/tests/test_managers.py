from unittest import mock

import pytest

from products_api.db import database
from products_api.managers import ProductManager

pytestmark = pytest.mark.asyncio


class TestProductManager:
    @pytest.mark.parametrize("required_field", ProductManager.required_fields)
    async def test_create_without_required_field(self, db_connection, required_field, product_data):
        del product_data[required_field]

        with pytest.raises(AttributeError) as exc:
            await ProductManager.create(product_data)
            assert str(exc.value) == f"{required_field} is a required fiel"

    @database.transaction(force_rollback=True)
    async def test_create_with_valid_data(self, db_connection, product_data):
        del product_data["id"]

        assert await ProductManager.create(product_data) is None

    @database.transaction(force_rollback=True)
    async def test_get_all(self, db_connection, product_data):
        del product_data["id"]
        await ProductManager.create(product_data)

        result = await ProductManager.get_all()
        assert len(result) == 1
        assert result[0]["title"] == product_data["title"]
        assert result[0]["description"] == product_data["description"]
        assert result[0]["price_in_cents"] == product_data["price_in_cents"]

    @database.transaction(force_rollback=True)
    @mock.patch("products_api.managers.CalculatorClient.get_product_discount")
    async def test_get_all_with_discount(
        self, get_product_discount_mock, db_connection, product_data, user_id
    ):
        del product_data["id"]
        await ProductManager.create(product_data)
        discount = {"pct": 0.1, "value_in_cents": 100}
        get_product_discount_mock.return_value = discount

        result = await ProductManager.get_all_with_discount(user_id)
        assert len(result) == 1
        assert result[0]["discount"] == discount
        assert result[0]["title"] == product_data["title"]
        assert result[0]["description"] == product_data["description"]
        assert result[0]["price_in_cents"] == product_data["price_in_cents"]
