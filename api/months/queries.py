from psycopg2.extensions import cursor
from psycopg2.extras import RealDictRow

from core.models import Table
from months.models import NewMonth


def db_add_month(cursor: cursor, month: NewMonth) -> int:
    cursor.execute(
        f"INSERT INTO {Table.MONTHS}(month_id, starting_balance) VALUES(%s, %s)",
        (month.month_id, month.starting_balance),
    )
    return cursor.lastrowid


def db_get_months(cursor: cursor) -> list[RealDictRow]:
    cursor.execute(f"SELECT * from {Table.MONTHS}")

    return cursor.fetchall()


def db_get_month_by_id(
    cursor: cursor,
    month_id: str,
) -> RealDictRow:
    cursor.execute(
        f"SELECT * from {Table.MONTHS} WHERE id = %s",
        (month_id,),
    )

    return cursor.fetchone()


def db_update_month(
    cursor: cursor, month_id: str, updated_month: dict[str, int | str]
) -> int:
    updates = ", ".join(f"{key} = %s" for key in updated_month)

    query_string = f"UPDATE {Table.MONTHS} SET {updates} WHERE id = %s"

    parameters = [*updated_month.values(), month_id]

    cursor.execute(query_string, parameters)

    return cursor.rowcount
