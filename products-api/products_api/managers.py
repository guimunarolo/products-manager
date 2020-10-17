from .clients import CalculatorClient
from .db import database


class ProductManager:
    table = "products"
    required_fields = ("price_in_cents", "title", "description")

    @classmethod
    def _validate_required_fields(cls, data):
        for field in cls.required_fields:
            if field not in data:
                raise AttributeError(f"{field} is a required field")

    @classmethod
    async def get_all(cls):
        query = f"SELECT * FROM {cls.table}"
        return await database.fetch_all(query)

    @classmethod
    async def create(cls, data):
        cls._validate_required_fields(data)
        query = f"INSERT INTO {cls.table}(price_in_cents, title, description) VALUES (:price_in_cents, :title, :description)"
        return await database.execute(query=query, values=data)

    @classmethod
    async def get_all_with_discount(cls, user_id):
        calculator_client = CalculatorClient()
        products = []
        for product in await cls.get_all():
            product = dict(product)
            product["discount"] = calculator_client.get_product_discount(
                user_id=user_id, product_id=str(product["id"]),
            )
            products.append(product)

        return products
