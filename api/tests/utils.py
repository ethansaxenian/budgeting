import sqlite3
from datetime import date
from sqlite3 import Cursor

from core.config import settings
from core.models import Category, MonthId, Table, TransactionType
from plans.models import Plan
from transactions.models import Transaction

example_transaction = Transaction(
    id=1,
    amount=20.25,
    date=date(day=11, month=6, year=2023),
    description="test expense",
    type=TransactionType.EXPENSE,
    category=Category.OTHER,
    month_id="6-2023",
)

example_plan = Plan(
    id=1,
    amount=100,
    type=TransactionType.EXPENSE,
    category=Category.OTHER,
    month_id="6-2023",
    month=MonthId.JUNE,
    year=2023,
)


def init_db(cursor: Cursor):
    for table in list(Table):
        cursor.execute(f"DROP TABLE IF EXISTS {table}")

    cursor.execute(
        f"CREATE TABLE IF NOT EXISTS {Table.TRANSACTIONS} (id INTEGER PRIMARY KEY, date DATE, amount FLOAT, description TEXT, category TEXT, type TEXT, month_id STRING)"
    )
    cursor.execute(f"CREATE TABLE IF NOT EXISTS {Table.MONTHS} (id STRING, starting_balance FLOAT)")
    cursor.execute(
        f"CREATE TABLE IF NOT EXISTS {Table.PLANS} (id INTEGER PRIMARY KEY, month INTEGER, year INTEGER, category STRING, amount FLOAT, type TEXT, month_id STRING)"
    )


def seed_db(cursor: Cursor):
    cursor.execute(
        f"INSERT INTO {Table.TRANSACTIONS}(id, date, amount, description, category, type, month_id) VALUES(%s, %s, %s, %s, %s, %s, %s)",
        (
            example_transaction.id,
            example_transaction.date,
            example_transaction.amount,
            example_transaction.description,
            example_transaction.category,
            example_transaction.type,
            example_transaction.month_id,
        ),
    )

    cursor.execute(
        f"INSERT INTO {Table.PLANS}(id, month, year, category, amount, type, month_id) VALUES(%s, %s, %s, %s, %s, %s, %s)",
        (
            example_plan.id,
            example_plan.month,
            example_plan.year,
            example_plan.category,
            example_plan.amount,
            example_plan.type,
            example_plan.month_id,
        ),
    )


def override_get_db():
    with sqlite3.connect(settings.TEST_DB_PATH) as connection:
        connection.row_factory = sqlite3.Row
        cursor = connection.cursor()
        init_db(cursor)
        seed_db(cursor)
        yield connection.cursor()
