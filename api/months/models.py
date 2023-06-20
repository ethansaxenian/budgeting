from pydantic import BaseModel


class NewMonth(BaseModel):
    month_id: str
    starting_balance: float


class Month(NewMonth):
    id: int
