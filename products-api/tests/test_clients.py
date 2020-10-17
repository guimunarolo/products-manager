import uuid
from unittest import mock

import pytest
from grpc import RpcError

from products_api.clients import CalculatorClient
from products_api.protos.calculator import calculator_pb2


class TestCalculatorClient:
    def test_get_product_discount_with_unavailable_server(self, user_id, product_id):
        client = CalculatorClient()
        client.stub = mock.Mock()
        client.stub.GetProductDiscount = mock.Mock(side_effect=RpcError)

        expected_request = calculator_pb2.ProductDiscountRequest(
            user_id=user_id,
            product_id=product_id,
        )

        assert client.get_product_discount(user_id, product_id) == {}
        client.stub.GetProductDiscount.assert_called_once_with(expected_request)

    def test_get_product_discount_with_successful_request(self, user_id, product_id):
        client = CalculatorClient()
        client.stub = mock.Mock()
        response = calculator_pb2.ProductDiscountResponse(pct=0.1, value_in_cents=100)
        client.stub.GetProductDiscount = mock.Mock(return_value=response)

        expected_request = calculator_pb2.ProductDiscountRequest(
            user_id=user_id,
            product_id=product_id,
        )
        expected_response = {"pct": response.pct, "value_in_cents": response.value_in_cents}

        assert client.get_product_discount(user_id, product_id) == expected_response
        client.stub.GetProductDiscount.assert_called_once_with(expected_request)
