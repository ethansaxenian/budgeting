from datetime import date

from fastapi.encoders import jsonable_encoder
from fastapi.testclient import TestClient

from core.dependencies import get_db
from core.models import Category, TransactionType
from main import app
from tests.utils import example_transaction, override_get_db
from transactions.models import NewTransaction

app.dependency_overrides[get_db] = override_get_db

client = TestClient(app)


def test_get_transactions():
    response = client.get("/api/transactions/")

    assert response.status_code == 200
    assert len(response.json()) == 1
    assert response.json()[0] == jsonable_encoder(example_transaction)


def test_get_transaction():
    response = client.get("/api/transactions/1")

    assert response.status_code == 200
    assert response.json() == jsonable_encoder(example_transaction)


def test_get_transaction_not_found():
    response = client.get("/api/transactions/9999")

    assert response.status_code == 404


def test_add_transaction():
    new_transaction = NewTransaction(
        type=TransactionType.EXPENSE,
        amount=100.0,
        description="Test food expense",
        category=Category.FOOD,
        date=date(day=3, month=6, year=2023),
    )
    response = client.post("/api/transactions/", json=jsonable_encoder(new_transaction.dict()))

    assert response.status_code == 200
    assert response.json() == 2


def test_remove_transaction():
    response = client.delete("/api/transactions/1")

    assert response.status_code == 200
    assert response.json() == 1


def test_update_transaction():
    new_transaction = NewTransaction(
        amount=50,
        description="Groceries",
        category=Category.FOOD,
        type=TransactionType.EXPENSE,
        date=date(day=11, month=6, year=2023),
    )

    response = client.post(
        "/api/transactions/1",
        json=jsonable_encoder(new_transaction.dict(exclude_unset=True)),
    )

    assert response.status_code == 200
    assert response.json() == 1
