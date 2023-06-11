from datetime import date as _date

from pydantic import BaseModel

from core.models import Category, TransactionType


class NewTransaction(BaseModel):
    type: TransactionType
    amount: float
    description: str | None
    category: Category = Category.OTHER
    date: _date = _date.today()


class Transaction(NewTransaction):
    id: int
    month_id: str
