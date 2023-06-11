from fastapi import APIRouter

from core.models import DBType, TransactionType
from months.queries import db_get_month_by_id
from plans.queries import db_get_plans
from transactions.queries import db_get_transactions

stats_router = APIRouter(prefix="/stats", tags=["stats"])


@stats_router.get("/{month_id}")
def get_stats(month_id: str, db: DBType):
    month = db_get_month_by_id(db, month_id)
    expenses = db_get_transactions(db, month["id"], TransactionType.EXPENSE)
    income = db_get_transactions(db, month["id"], TransactionType.INCOME)
    planned_expenses = db_get_plans(db, month["id"], TransactionType.EXPENSE)
    planned_income = db_get_plans(db, month["id"], TransactionType.INCOME)

    total_expenses = round(sum(e["amount"] for e in expenses), 2)
    total_income = round(sum(i["amount"] for i in income), 2)

    return {
        "Starting balance": month["starting_balance"],
        "Ending balance": month["starting_balance"] + total_income - total_expenses,
        "Amount saved": total_income - total_expenses,
        "Planned expenses": sum(e["amount"] for e in planned_expenses),
        "Actual expenses": total_expenses,
        "Planned income": sum(i["amount"] for i in planned_income),
        "Actual income": total_income,
    }
