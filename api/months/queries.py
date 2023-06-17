from sqlite3 import Cursor, Row

from core.models import Table
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
) -> Row:
    query = cursor.execute(
        f"SELECT * from {Table.MONTHS} WHERE id = ?",
        (month_id,),
    )

    return query.fetchone()


def db_update_month(
    cursor: Cursor, month_id: str, updated_month: dict[str, int | str]
) -> int:
    updates = ", ".join(f"{key} = ?" for key in updated_month)

    query_string = f"UPDATE {Table.MONTHS} SET {updates} WHERE id = ?"

    parameters = [*updated_month.values(), month_id]

    cursor.execute(query_string, parameters)

    return cursor.rowcount
