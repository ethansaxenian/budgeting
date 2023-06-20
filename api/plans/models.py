from pydantic import BaseModel

from core.models import Category, TransactionType


class NewPlan(BaseModel):
    amount: float
    type: TransactionType
    category: Category = Category.OTHER
    month_id: str


class Plan(NewPlan):
    id: int
