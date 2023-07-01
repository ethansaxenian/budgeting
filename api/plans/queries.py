from psycopg2.extensions import cursor
from psycopg2.extras import RealDictRow

from core.models import Table, TransactionType
from core.utils import build_query_with_optional_params
from plans.models import NewPlan


def db_add_plan(cursor: cursor, plan: NewPlan) -> int:
    cursor.execute(
        f"INSERT INTO {Table.PLANS}"
        f"(type, month_id, food, gifts, medical, home, transportation, "
        f"personal, savings, utilities, travel, other, paycheck, bonus, interest) "
        f"VALUES(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
        (
            plan.type,
            plan.month_id,
            plan.food,
            plan.gifts,
            plan.medical,
            plan.home,
            plan.transportation,
            plan.personal,
            plan.savings,
            plan.utilities,
            plan.travel,
            plan.other,
            plan.paycheck,
            plan.bonus,
            plan.interest,
        ),
    )
    return cursor.lastrowid


def db_get_plans(
    cursor: cursor,
    month_id: str | None = None,
    transaction_type: TransactionType | None = None,
) -> list[RealDictRow]:
    query_string, parameters = build_query_with_optional_params(
        Table.PLANS, month_id=month_id, type=transaction_type
    )

    cursor.execute(query_string, parameters)

    return cursor.fetchall()


def db_get_plan_by_id(
    cursor: cursor,
    plan_id: int,
) -> RealDictRow:
    cursor.execute(
        f"SELECT * from {Table.PLANS} WHERE id = %s",
        (plan_id,),
    )

    return cursor.fetchone()


def db_delete_plan(cursor: cursor, plan_id: int) -> int:
    cursor.execute(
        f"DELETE from {Table.PLANS} WHERE id = %s",
        (plan_id,),
    )
    return cursor.rowcount


def db_update_plan(
    cursor: cursor, plan_id: int, updated_plan: dict[str, int | str]
) -> int:
    updates = ", ".join(f"{key} = %s" for key in updated_plan)

    query_string = f"UPDATE {Table.PLANS} SET {updates} WHERE id = %s"

    parameters = [*updated_plan.values(), plan_id]

    cursor.execute(query_string, parameters)

    return cursor.rowcount
