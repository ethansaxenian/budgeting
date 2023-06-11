from fastapi import HTTPException, status

from core.models import Table


def build_month_id(month: int, year: int) -> str:
    return f"{month}-{year}"


def build_query_with_optional_params(table: Table, **kwargs) -> tuple[str, list[str]]:
    if table not in list(Table):
        raise HTTPException(status_code=status.HTTP_500_INTERNAL_SERVER_ERROR)

    query_string = f"SELECT * FROM {table} WHERE 1=1"  # noqa: S608

    parameters = []

    for var, value in kwargs.items():
        if value is not None:
            query_string += f" AND {var} = ?"
            parameters.append(value)

    return query_string, parameters
