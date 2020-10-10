import datetime
import uuid

from pydantic import BaseModel


class User(BaseModel):
    id: uuid.UUID
    first_name: str
    last_name: str
    date_of_birth: datetime.date
