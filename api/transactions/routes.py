from fastapi import APIRouter, HTTPException, Query, status

from core.models import DBType, TransactionType
from transactions.models import NewTransaction, Transaction
from transactions.queries import (
    db_add_transaction,
    db_delete_transaction,
    db_get_transaction_by_id,
    db_get_transactions,
    db_update_transaction,
)

transactions_router = APIRouter(
    prefix="/transactions",
    tags=["transactions"],
)


@transactions_router.get("/", response_model=list[Transaction])
def get_transactions(
    db: DBType,
    month_id: str | None = Query(default=None, description="The month id"),
    transaction_type: TransactionType
    | None = Query(default=None, description="The transaction type"),
):
    return db_get_transactions(db, month_id, transaction_type)


@transactions_router.get(
    "/{id}", response_model=Transaction, responses={404: {"description": "Not found"}}
)
def get_transaction(id: int, db: DBType):
    transaction = db_get_transaction_by_id(db, id)

    if transaction is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail=f"Transaction with id {id} not found",
        )

    return transaction


@transactions_router.post("/", response_model=int)
def add_transaction(transaction: NewTransaction, db: DBType):
    return db_add_transaction(db, transaction)


@transactions_router.delete("/{id}", response_model=int)
def remove_transaction(id: int, db: DBType):
    return db_delete_transaction(db, id)


@transactions_router.put("/{id}", response_model=int)
def update_transaction(id: int, transaction: NewTransaction, db: DBType):
    return db_update_transaction(db, id, transaction.dict())
