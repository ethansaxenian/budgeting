from enum import Enum
from typing import Annotated

from fastapi import Depends
from psycopg2.extensions import cursor

from core.dependencies import get_db

DBType = Annotated[cursor, Depends(get_db)]


class Table(str, Enum):
    TRANSACTIONS = "transactions"
    MONTHS = "months"
    PLANS = "plans"

    def __str__(self):
        return self.value


class TransactionType(str, Enum):
    INCOME = "income"
    EXPENSE = "expense"


class Category(str, Enum):
    FOOD = "food"
    GIFTS = "gifts"
    MEDICAL = "medical"
    HOME = "home"
    TRANSPORTATION = "transportation"
    PERSONAL = "personal"
    SAVINGS = "savings"
    UTILITIES = "utilities"
    TRAVEL = "travel"
    OTHER = "other"
    PAYCHECK = "paycheck"
    BONUS = "bonus"
    INTEREST = "interest"

    def __str__(self):
        return self.value
