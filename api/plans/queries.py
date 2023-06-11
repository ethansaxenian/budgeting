from sqlite3 import Cursor

from core.models import Table, TransactionType
from core.utils import build_query_with_optional_params
from plans.models import Plan


def add_plan(cursor: Cursor, plan: Plan):
    cursor.execute(
        f"INSERT INTO {Table.PLANS}"
        f"(month, year, category, amount, type, month_id) VALUES(?, ?, ?, ?, ?, ?)",
        (
            plan.month,
            plan.year,
            plan.category,
            plan.amount,
            plan.type,
            plan.month_id,
        ),
    )


def db_get_plans(
    cursor: Cursor,
    month_id: str | None = None,
    transaction_type: TransactionType | None = None,
):
    query_string, parameters = build_query_with_optional_params(
        Table.PLANS, month_id=month_id, type=transaction_type
    )

    query = cursor.execute(query_string, parameters)

    return query.fetchall()
