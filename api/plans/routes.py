from fastapi import APIRouter, Query

from core.models import DBType, TransactionType
from plans.models import Plan
from plans.queries import db_get_plans

plans_router = APIRouter()


@plans_router.get("/", response_model=list[Plan])
def get_plans(
    db: DBType,
    month_id: str | None = Query(default=None),
    transaction_type: TransactionType | None = Query(default=None),
):
    return db_get_plans(db, month_id, transaction_type)


#
# @router.get("/{id}", response_model=Plan)
# def get_plan(id: str, db: DBType):
#     pass
#
#
# @router.post("/")
# def add_plan(plan: Plan, db: DBType):
#     pass
#
#
# @router.delete("/")
# def remove_plan(id: str, db: DBType):
#     pass
#
#
# @router.put("/{id}")
# def update_plan(id: str, plan: Plan, db: DBType):
#     pass
