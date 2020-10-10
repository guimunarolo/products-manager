import datetime
import uuid
from typing import Dict

from pydantic import BaseModel


class User(BaseModel):
    id: uuid.UUID
    first_name: str
    last_name: str
    date_of_birth: datetime.date


class Product(BaseModel):
    id: uuid.UUID
    price_in_cents: int
    title: str
    description: str
    discount: Dict = {}
