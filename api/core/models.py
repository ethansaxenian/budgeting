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
    FOOD = "Food"
    GIFTS = "Gifts"
    MEDICAL = "Medical"
    HOME = "Home"
    TRANSPORTATION = "Transportation"
    PERSONAL = "Personal"
    SAVINGS = "Savings"
    UTILITIES = "Utilities"
    TRAVEL = "Travel"
    OTHER = "Other"
    PAYCHECK = "Paycheck"
    BONUS = "Bonus"
    INTEREST = "Interest"

    def __str__(self):
        return self.value


class MonthId(int, Enum):
    JANUARY = 1
    FEBRUARY = 2
    MARCH = 3
    APRIL = 4
    MAY = 5
    JUNE = 6
    JULY = 7
    AUGUST = 8
    SEPTEMBER = 9
    OCTOBER = 10
    NOVEMBER = 11
    DECEMBER = 12

    def __str__(self):
        return str(self.value)
