from psycopg2.extensions import cursor
from psycopg2.extras import RealDictRow

from core.models import Table, TransactionType
from core.utils import (
    build_query_with_optional_params,
)
from transactions.models import NewTransaction


def db_add_transaction(cursor: cursor, transaction: NewTransaction) -> int:
    cursor.execute(
        f"INSERT INTO {Table.TRANSACTIONS}(date, amount, description, category, type, month_id) VALUES(%s, %s, %s, %s, %s, %s)",
        (
            transaction.date,
            transaction.amount,
            transaction.description,
            transaction.category,
            transaction.type,
            transaction.month_id,
        ),
    )
    return cursor.lastrowid


def db_get_transactions(
    cursor: cursor,
    month_id: str | None = None,
    transaction_type: TransactionType | None = None,
) -> list[RealDictRow]:
    query_string, parameters = build_query_with_optional_params(
        Table.TRANSACTIONS, month_id=month_id, type=transaction_type
    )

    cursor.execute(query_string, parameters)

    return cursor.fetchall()


def db_get_transaction_by_id(
    cursor: cursor,
    transaction_id: int,
) -> RealDictRow:
    cursor.execute(
        f"SELECT * from {Table.TRANSACTIONS} WHERE id = %s",
        (transaction_id,),
    )

    return cursor.fetchone()


def db_delete_transaction(cursor: cursor, transaction_id: int) -> int:
    cursor.execute(
        f"DELETE from {Table.TRANSACTIONS} WHERE id = %s",
        (transaction_id,),
    )
    return cursor.rowcount


def db_update_transaction(
    cursor: cursor, transaction_id: int, updated_transaction: dict[str, str | int]
) -> int:
    updates = ", ".join(f"{key} = %s" for key in updated_transaction)

    query_string = f"UPDATE {Table.TRANSACTIONS} SET {updates} WHERE id = %s"

    parameters = [*updated_transaction.values(), transaction_id]

    cursor.execute(query_string, parameters)

    return cursor.rowcount
