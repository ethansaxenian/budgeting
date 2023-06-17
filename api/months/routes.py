from fastapi import APIRouter

from core.models import DBType
from months.models import Month, NewMonth
from months.queries import db_get_month_by_id, db_get_months, db_update_month

months_router = APIRouter(prefix="/months", tags=["months"])


@months_router.get("/", response_model=list[Month])
def get_months(db: DBType):
    return db_get_months(db)


@months_router.get("/{id}", response_model=Month)
def get_month(id: str, db: DBType):
    return db_get_month_by_id(db, id)


#
# @router.post("/")
# def add_transaction(transaction: Transaction, db: DBType):
#     pass
#
#
# @router.delete("/")
# def remove_transaction(id: str, db: DBType):
#     pass
#
#
@months_router.put("/{id}", response_model=int)
def update_month(id: str, month: NewMonth, db: DBType):
    return db_update_month(db, id, month.dict())
