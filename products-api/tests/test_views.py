
from fastapi import status


def test_hello_world(client):
    response = client.get("/")
    response_data = response.json()

    assert response.status_code == status.HTTP_200_OK
    assert response_data == {"Hello": "World"}
