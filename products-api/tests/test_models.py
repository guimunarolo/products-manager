import datetime

from products_api.models import Product, User


def test_user_parsing(user_data):
    user = User(**user_data)

    assert str(user.id) == user_data["id"]
    assert user.first_name == user_data["first_name"]
    assert user.last_name == user_data["last_name"]
    assert user.date_of_birth == datetime.date.fromisoformat(user_data["date_of_birth"])


def test_product_parsing(product_data):
    product = Product(**product_data)

    assert str(product.id) == product_data["id"]
    assert product.price_in_cents == product_data["price_in_cents"]
    assert product.title == product_data["title"]
    assert product.description == product_data["description"]
    assert product.discount == {}
