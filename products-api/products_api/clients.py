import grpc

from .configs import settings
from .protos.calculator import calculator_pb2_grpc as pb2_grpc
from .protos.calculator import calculator_pb2 as pb2


class CalculatorClient:
    """
    Client for gRPC functionality
    """

    def __init__(self):
        self.channel = grpc.insecure_channel(settings.CALCULATOR_URL)
        self.stub = pb2_grpc.CalculatorStub(self.channel)

    def get_product_discount(self, user_id, product_id):
        request = pb2.ProductDiscountRequest(user_id=user_id, product_id=product_id)

        try:
            response = self.stub.GetProductDiscount(request)
        except grpc.RpcError:
            return {}
        else:
            return {
                "pct": response.pct,
                "value_in_cents": response.value_in_cents,
            }
