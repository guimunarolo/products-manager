from typing import List

from fastapi import APIRouter

from .managers import ProductManager
from .models import Product

router = APIRouter()


@router.get("/products/", response_model=List[Product])
async def list_products():
    return await ProductManager.get_all()
