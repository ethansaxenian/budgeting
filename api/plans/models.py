from datetime import date

from pydantic import BaseModel

from core.models import Category, MonthId, TransactionType


class NewPlan(BaseModel):
    amount: float
    type: TransactionType
    category: Category = Category.OTHER
    month: MonthId = MonthId(date.today().month)
    year: int = date.today().year


class Plan(NewPlan):
    id: int
    month_id: str
