from functions import add_month, add_plan, add_transaction, get_months, get_transactions
from models import Category, Month, MonthId, Plan, Transaction, TransactionType
from datetime import date as _date
import polars as pl

from utils import convert_to_date, is_date, transform_input

pl.Config.set_tbl_rows(1000)
pl.Config.set_tbl_hide_column_data_types()
pl.Config.set_tbl_hide_dataframe_shape()


def prompt_action():
    while True:
        action = input(
            "What do you want to do?\n"
            "a: add an item\n"
            "s: show a month\n"
            "q: exit\n"
            ">>> "
        )
        if action in "asq":
            return action

        print("Invalid action\n")


def prompt_add():
    while True:
        action = input(
            "What do you want to add?\n"
            "m: add an month\n"
            "e: add an expense \n"
            "i: add a source of income\n"
            "pe: add a planned expense\n"
            "pi: add a planned source of income\n"
            "b: go back\n"
            ">>> "
        )
        if action in ("m", "e", "i", "pe", "pi", "b"):
            return action

        print("Invalid item\n")


def prompt_add_month():
    starting_balance = transform_input(
        "Enter your starting balance:\n", test=float, callback=float
    )

    today = _date.today()

    month = transform_input(
        f"Enter the month (leave blank for {MonthId(today.month)}):\n",
        test=lambda val: val == "" or MonthId(int(val)),
        callback=lambda val: MonthId(int(val or today.month)),
    )

    year = transform_input(
        f"Enter the year (leave blank for {today.year}):\n",
        test=lambda val: val == "" or int(val),
        callback=lambda val: int(val or today.year),
    )

    new_month = Month(starting_balance=starting_balance, name=month, year=year)

    print(f"Adding {new_month}")

    add_month(new_month)


def prompt_add_transaction(transaction_type: TransactionType):
    today = _date.today()

    date = transform_input(
        f"Enter the date (leave blank for {today}):\n",
        test=lambda val: val == "" or is_date(val),
        callback=lambda val: convert_to_date(val) if val else today,
    )

    amount = transform_input("Enter the amount:\n", test=float, callback=float)

    description = input("Enter a description:\n")

    category_list = "\n".join(f"{i}: {cat}" for i, cat in enumerate(Category))
    category = transform_input(
        f"Enter a category (leave blank for {Category.OTHER.value}):\n{category_list}\n",
        test=lambda val: val == "" or int(val) in range(len(Category)),
        callback=lambda val: list(Category)[int(val)] if val else Category.OTHER,
    )

    new_transaction = Transaction(
        type=transaction_type,
        amount=amount,
        description=description,
        category=category,
        date=date,
    )

    print(f"Adding {new_transaction}")
    add_transaction(new_transaction)


def prompt_add_plan(transaction_type: TransactionType):
    category_list = "\n".join(f"{i}: {cat}" for i, cat in enumerate(Category))
    category = transform_input(
        f"Enter a category (leave blank for {Category.OTHER.value}):\n{category_list}\n",
        test=lambda val: val == "" or int(val) in range(len(Category)),
        callback=lambda val: list(Category)[int(val)] if val else Category.OTHER,
    )

    amount = transform_input(
        "Enter the amount:\n", test=lambda val: val == "0" or float(val), callback=float
    )

    today = _date.today()

    month = transform_input(
        f"Enter the month (leave blank for {MonthId(today.month)}):\n",
        test=lambda val: val == "" or MonthId(int(val)),
        callback=lambda val: MonthId(int(val or today.month)),
    )

    year = transform_input(
        f"Enter the year (leave blank for {today.year}):\n",
        test=lambda val: val == "" or int(val),
        callback=lambda val: int(val or today.year),
    )

    new_plan = Plan(
        month=month, year=year, category=category, amount=amount, type=transaction_type
    )

    print(f"Adding {new_plan}")
    add_plan(new_plan)


def add():
    match prompt_add():
        case "m":
            prompt_add_month()
        case "e":
            prompt_add_transaction(TransactionType.EXPENSE)
        case "i":
            prompt_add_transaction(TransactionType.INCOME)
        case "pe":
            prompt_add_plan(TransactionType.EXPENSE)
        case "pi":
            prompt_add_plan(TransactionType.INCOME)
        case "b":
            return


def prompt_show(months):
    month_list = "\n".join(
        f"{i}: {month.id}"
        for i, month in enumerate(
            sorted(
                months,
            )
        )
    )

    while True:
        try:
            action = int(
                input(f"What month do you want to show?\n{month_list}\n" ">>> ")
            )
            if action in range(len(months)):
                return months[action]
        except ValueError:
            continue


def show():
    month = prompt_show(get_months())
    expenses = get_transactions(month.id, TransactionType.EXPENSE)
    income = get_transactions(month.id, TransactionType.INCOME)

    total_expenses = round(sum(e.amount for e in expenses), 2)
    total_income = round(sum(i.amount for i in income), 2)

    print(f"Starting balance: {month.starting_balance}")
    print(f"Ending balance: {month.starting_balance + total_income - total_expenses}")
    print(f"Amount saved: {total_income-total_expenses}\n")
    print(f"Planned expenses: {''}")
    print(f"Actual expenses: {total_expenses}\n")

    print(f"Planned income: {''}")
    print(f"Actual income: {total_income}\n")

    expenses_df = pl.DataFrame([item.to_dict() for item in expenses])
    income_df = pl.DataFrame([item.to_dict() for item in income])

    if not expenses_df.is_empty():
        print("EXPENSES:")
        print(expenses_df.sort("Date"), end="\n\n")

    if not income_df.is_empty():
        print("INCOME:")
        print(income_df.sort("Date"))


if __name__ == "__main__":
    while True:
        match prompt_action():
            case "a":
                add()
            case "s":
                show()
            case "q":
                break
