from fastapi import APIRouter

from core.models import DBType
from months.models import Month
from months.queries import db_get_months

months_router = APIRouter()


@months_router.get("/", response_model=list[Month])
def get_months(db: DBType):
    return db_get_months(db)


# @router.get("/{id}", response_model=Transaction)
# def get_transaction(id: str, db: DBType):
#     pass
#
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
# @router.put("/{id}")
# def update_transaction(id: str, transaction: Transaction, db: DBType):
#     pass
