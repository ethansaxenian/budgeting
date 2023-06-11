from fastapi import APIRouter

from months.routes import months_router
from plans.routes import plans_router
from stats.routes import stats_router
from transactions.routes import transactions_router

api_router = APIRouter()
api_router.include_router(transactions_router)
api_router.include_router(plans_router)
api_router.include_router(months_router)
api_router.include_router(stats_router)
