from fastapi import APIRouter

from core.models import DBType
from months.models import Month, NewMonth
from months.queries import (
    db_add_month,
    db_get_month_by_id,
    db_get_months,
    db_update_month,
)

months_router = APIRouter(prefix="/months", tags=["months"])


@months_router.get("/", response_model=list[Month])
def get_months(db: DBType):
    return db_get_months(db)


@months_router.get("/{id}", response_model=Month)
def get_month(id: str, db: DBType):
    return db_get_month_by_id(db, id)


@months_router.post("/")
def add_month(month: NewMonth, db: DBType):
    return db_add_month(db, month)


@months_router.put("/{id}", response_model=int)
def update_month(id: str, month: NewMonth, db: DBType):
    return db_update_month(db, id, month.dict())
