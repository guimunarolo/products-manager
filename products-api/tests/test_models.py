import datetime

from products_api.models import User


def test_user_parsing(user_data):
    user = User(**user_data)

    assert user.id == user_data["id"]
    assert user.first_name == user_data["first_name"]
    assert user.last_name == user_data["last_name"]
    assert user.date_of_birth == datetime.date.fromisoformat(user_data["date_of_birth"])
