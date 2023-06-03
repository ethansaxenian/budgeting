from db import db
from models import Month, Plan, Table, Transaction, TransactionType


def init_db(reset=False):
    with db() as cursor:
        if reset:
            for table_name in list(Table):
                cursor.execute(f"DROP TABLE IF EXISTS {table_name}")

        cursor.execute(
            f"CREATE TABLE IF NOT EXISTS {Table.TRANSACTIONS} (id INTEGER PRIMARY KEY, date DATE, amount FLOAT, description TEXT, category TEXT, type TEXT, month_id STRING)"
        )
        cursor.execute(
            f"CREATE TABLE IF NOT EXISTS {Table.MONTHS} (id STRING, starting_balance FLOAT)"
        )
        cursor.execute(
            f"CREATE TABLE IF NOT EXISTS {Table.PLANS} (id INTEGER PRIMARY KEY, month INTEGER, year INTEGER, category STRING, amount FLOAT, type TEXT, month_id STRING)"
        )


def add_transaction(transaction: Transaction):
    with db() as cursor:
        cursor.execute(
            f"INSERT INTO {Table.TRANSACTIONS}(date, amount, description, category, type, month_id) VALUES(?, ?, ?, ?, ?, ?)",
            (
                transaction.date,
                transaction.amount,
                transaction.description,
                transaction.category,
                transaction.type,
                transaction.month_id,
            ),
        )


def get_transactions(month_id: str, transaction_type: TransactionType):
    with db() as cursor:
        query = cursor.execute(
            f"SELECT * from {Table.TRANSACTIONS} WHERE month_id = ? AND type = ?",
            (month_id, transaction_type),
        )
        res = query.fetchall()

        return [Transaction.from_db(item) for item in res]


def add_month(month: Month):
    with db() as cursor:
        query = cursor.execute(
            f"SELECT COUNT(id) FROM {Table.MONTHS} WHERE id = ?",
            (month.id,),
        )
        month_exists, *_ = query.fetchall()[0]
        if month_exists:
            return

        cursor.execute(
            f"INSERT INTO {Table.MONTHS}(id, starting_balance) VALUES(?, ?)",
            (month.id, month.starting_balance),
        )


def get_months():
    with db() as cursor:
        query = cursor.execute(f"SELECT * from {Table.MONTHS}")
        res = query.fetchall()

        return [Month.from_db(item) for item in res]


def add_plan(plan: Plan):
    with db() as cursor:
        cursor.execute(
            f"INSERT INTO {Table.PLANS}(month, year, category, amount, type, month_id) VALUES(?, ?, ?, ?, ?, ?)",
            (
                plan.month,
                plan.year,
                plan.category,
                plan.amount,
                plan.type,
                plan.month_id,
            ),
        )


def get_plans(month_id: str, transaction_type: TransactionType):
    with db() as cursor:
        query = cursor.execute(
            f"SELECT * from {Table.PLANS} WHERE month_id = ? AND type = ?",
            (month_id, transaction_type),
        )
        res = query.fetchall()

        return [Plan.from_db(item) for item in res]
