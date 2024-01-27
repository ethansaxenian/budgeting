from pydantic import BaseModel

from core.models import TransactionType


class NewPlan(BaseModel):
    type: TransactionType
    month_id: int
    food: float | None = 0
    gifts: float | None = 0
    medical: float | None = 0
    home: float | None = 0
    transportation: float | None = 0
    personal: float | None = 0
    savings: float | None = 0
    utilities: float | None = 0
    travel: float | None = 0
    other: float | None = 0
    paycheck: float | None = 0
    bonus: float | None = 0
    interest: float | None = 0


class UpdatePlan(NewPlan):
    type: TransactionType | None
    month_id: int | None


class Plan(NewPlan):
    id: int
