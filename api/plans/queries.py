from sqlite3 import Cursor, Row

from core.models import Table, TransactionType
from core.utils import build_month_id, build_query_with_optional_params
from plans.models import NewPlan


def db_add_plan(cursor: Cursor, plan: NewPlan) -> int:
    cursor.execute(
        f"INSERT INTO {Table.PLANS}"
        f"(month, year, category, amount, type, month_id) VALUES(?, ?, ?, ?, ?, ?)",
        (
            plan.month,
            plan.year,
            plan.category,
            plan.amount,
            plan.type,
            build_month_id(plan.month, plan.year),
        ),
    )
    return cursor.lastrowid


def db_get_plans(
    cursor: Cursor,
    month_id: str | None = None,
    transaction_type: TransactionType | None = None,
) -> list[Row]:
    query_string, parameters = build_query_with_optional_params(
        Table.PLANS, month_id=month_id, type=transaction_type
    )

    query = cursor.execute(query_string, parameters)

    return query.fetchall()


def db_get_plan_by_id(
    cursor: Cursor,
    plan_id: int,
) -> Row:
    query = cursor.execute(
        f"SELECT * from {Table.PLANS} WHERE id = ?",
        (plan_id,),
    )

    return query.fetchone()


def db_delete_plan(cursor: Cursor, plan_id: int) -> int:
    cursor.execute(
        f"DELETE from {Table.PLANS} WHERE id = ?",
        (plan_id,),
    )
    return cursor.rowcount


def db_update_plan(
    cursor: Cursor, plan_id: int, updated_plan: dict[str, int | str]
) -> int:
    updates = ", ".join(f"{key} = ?" for key in updated_plan)

    query_string = f"UPDATE {Table.PLANS} SET {updates} WHERE id = ?"

    parameters = [*updated_plan.values(), plan_id]

    cursor.execute(query_string, parameters)

    return cursor.rowcount
