from sqlite3 import Cursor

from core.models import QueryResult, Table
from months.models import Month


def add_month(cursor: Cursor, month: Month):
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


def db_get_months(cursor: Cursor):
    query = cursor.execute(f"SELECT * from {Table.MONTHS}")

    return query.fetchall()


def db_get_month_by_id(
    cursor: Cursor,
    month_id: str,
) -> QueryResult:
    query = cursor.execute(
        f"SELECT * from {Table.MONTHS} WHERE id = ?",
        (month_id,),
    )

    return query.fetchone()
