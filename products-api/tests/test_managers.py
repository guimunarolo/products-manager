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
            assert  str(exc.value) == f"{required_field} is a required fiel"

    @database.transaction(force_rollback=True)
    async def test_create_with_valid_data(self, db_connection, product_data):
        del product_data["id"]

        assert await ProductManager.create(product_data) is None

    @database.transaction(force_rollback=True)
    async def test_get_all(self, db_connection, product_data):
        del product_data["id"]

        assert len(await ProductManager.get_all()) == 0

        assert await ProductManager.create(product_data) is None
        
        assert len(await ProductManager.get_all()) == 1
