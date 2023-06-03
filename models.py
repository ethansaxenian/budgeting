from dataclasses import dataclass
from enum import Enum
from datetime import date as _date


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


@dataclass(frozen=True)
class Transaction:
    type: TransactionType
    amount: float
    description: str | None
    category: Category = Category.OTHER
    date: _date = _date.today()

    @property
    def month_id(self):
        return f"{self.date.month}-{self.date.year}"

    @classmethod
    def from_db(cls, db_item: tuple[int, str, float, str, str, str, str]):
        _, date, amount, description, category, type, _ = db_item
        y, m, d = map(int, date.split("-"))
        return cls(
            TransactionType(type),
            amount,
            description,
            Category(category),
            _date(day=d, month=m, year=y),
        )

    def to_dict(self):
        return {
            "Date": self.date,
            "Amount": self.amount,
            "Description": self.description,
            "Category": self.category,
        }


@dataclass(frozen=True)
class Month:
    starting_balance: float
    name: int = MonthId(_date.today().month)
    year: int = _date.today().year

    @property
    def id(self):
        return f"{self.name}-{self.year}"

    @classmethod
    def from_db(cls, db_item: tuple[str, float]):
        id, starting_balance = db_item
        name, year = map(int, id.split("-"))
        return cls(starting_balance, name, year)

    def __lt__(self, other):
        if self.year == other.year:
            return self.name < other.name

        return self.year < other.year


@dataclass(frozen=True)
class Plan:
    amount: float
    type: TransactionType
    category: Category = Category.OTHER
    month: MonthId = MonthId(_date.today().month)
    year: int = _date.today().year

    @property
    def month_id(self):
        return f"{self.month}-{self.year}"

    @classmethod
    def from_db(cls, db_item: tuple[int, int, int, str, float, str, str]):
        _, month, year, category, amount, transaction_type, _ = db_item
        return cls(
            amount=amount,
            type=TransactionType(transaction_type),
            category=Category(category),
            month=MonthId(month),
            year=year,
        )
