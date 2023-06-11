from datetime import date

from pydantic import BaseModel

from core.models import MonthId


class NewMonth(BaseModel):
    starting_balance: float
    name: int = MonthId(date.today().month)
    year: int = date.today().year

    def __lt__(self, other):
        if self.year == other.year:
            return self.name < other.name

        return self.year < other.year


class Month(NewMonth):
    id: str