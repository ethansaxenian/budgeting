from fastapi.encoders import jsonable_encoder
from fastapi.testclient import TestClient

from core.dependencies import get_db
from core.models import Category, MonthId, TransactionType
from main import app
from plans.models import NewPlan
from tests.utils import example_plan, override_get_db

app.dependency_overrides[get_db] = override_get_db

client = TestClient(app)


def test_get_plans():
    response = client.get("/api/plans/")

    assert response.status_code == 200
    assert len(response.json()) == 1
    assert response.json()[0] == jsonable_encoder(example_plan)


def test_get_plan():
    response = client.get("/api/plans/1")

    assert response.status_code == 200
    assert response.json() == jsonable_encoder(example_plan)


def test_get_plan_not_found():
    response = client.get("/api/plans/9999")

    assert response.status_code == 404


def test_add_plan():
    new_plan = NewPlan(
        type=TransactionType.EXPENSE,
        amount=100.0,
        category=Category.FOOD,
        month=MonthId.JUNE,
        year=2023,
    )
    response = client.post("/api/plans/", json=jsonable_encoder(new_plan.dict()))

    assert response.status_code == 200
    assert response.json() == 2


def test_remove_plan():
    response = client.delete("/api/plans/1")

    assert response.status_code == 200
    assert response.json() == 1


def test_update_plan():
    new_plan = NewPlan(
        type=TransactionType.EXPENSE,
        amount=200.0,
        category=Category.FOOD,
        month=MonthId.JUNE,
        year=2023,
    )

    response = client.post(
        "/api/plans/1",
        json=jsonable_encoder(new_plan.dict(exclude_unset=True)),
    )

    assert response.status_code == 200
    assert response.json() == 1
