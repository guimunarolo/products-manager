from typing import List, Optional

from fastapi import APIRouter, Header

from .managers import ProductManager
from .models import Product

router = APIRouter()


@router.get("/product", response_model=List[Product])
async def list_products(x_user_id: Optional[str] = Header(None)):
    if x_user_id:
        return await ProductManager.get_all_with_discount(user_id=x_user_id)

    return await ProductManager.get_all()
