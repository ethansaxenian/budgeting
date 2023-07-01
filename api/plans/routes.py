from fastapi import APIRouter, HTTPException, Query, status

from core.models import DBType, TransactionType
from plans.models import NewPlan, Plan, UpdatePlan
from plans.queries import (
    db_add_plan,
    db_delete_plan,
    db_get_plan_by_id,
    db_get_plans,
    db_update_plan,
)

plans_router = APIRouter(prefix="/plans", tags=["plans"])


@plans_router.get("/", response_model=list[Plan])
def get_plans(
    db: DBType,
    month_id: str | None = Query(default=None, description="The month id"),
    transaction_type: TransactionType
    | None = Query(default=None, description="The transaction type"),
):
    return db_get_plans(db, month_id, transaction_type)


@plans_router.get(
    "/{id}", response_model=Plan, responses={404: {"description": "Not found"}}
)
def get_plan(id: int, db: DBType):
    plan = db_get_plan_by_id(db, id)

    if plan is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Plan with id {id} not found",
        )

    return plan


@plans_router.post("/", response_model=int)
def add_plan(plan: NewPlan, db: DBType):
    return db_add_plan(db, plan)


@plans_router.delete("/{id}", response_model=int)
def remove_plan(id: int, db: DBType):
    return db_delete_plan(db, id)


@plans_router.patch("/{id}", response_model=int)
def update_plan(id: int, plan: UpdatePlan, db: DBType):
    return db_update_plan(db, id, plan.dict(exclude_unset=True))
