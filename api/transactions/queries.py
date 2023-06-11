from sqlite3 import Cursor

from core.models import QueryResult, Table, TransactionType
from core.utils import (
    build_month_id,
    build_query_with_optional_params,
)
from transactions.models import NewTransaction


def db_add_transaction(cursor: Cursor, transaction: NewTransaction) -> int:
    cursor.execute(
        f"INSERT INTO {Table.TRANSACTIONS}(date, amount, description, category, type, month_id) VALUES(?, ?, ?, ?, ?, ?)",
        (
            transaction.date,
            transaction.amount,
            transaction.description,
            transaction.category,
            transaction.type,
            build_month_id(transaction.date.month, transaction.date.year),
        ),
    )
    return cursor.lastrowid


def db_get_transactions(
    cursor: Cursor,
    month_id: str | None = None,
    transaction_type: TransactionType | None = None,
) -> list[QueryResult]:
    query_string, parameters = build_query_with_optional_params(
        Table.TRANSACTIONS, month_id=month_id, type=transaction_type
    )

    query = cursor.execute(query_string, parameters)

    return query.fetchall()


def db_get_transaction_by_id(
    cursor: Cursor,
    transaction_id: int,
) -> QueryResult:
    query = cursor.execute(
        f"SELECT * from {Table.TRANSACTIONS} WHERE id = ?",
        (transaction_id,),
    )

    return query.fetchone()


def db_delete_transaction(cursor: Cursor, transaction_id: int) -> int:
    cursor.execute(
        f"DELETE from {Table.TRANSACTIONS} WHERE id = ?",
        (transaction_id,),
    )
    return cursor.rowcount
