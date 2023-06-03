from datetime import date, datetime


def is_float(num):
    try:
        float(num)
        return True
    except ValueError:
        return False


def convert_to_date(string):
    dt = datetime.strptime(string, "%Y-%m-%d")
    return date(year=dt.year, month=dt.month, day=dt.day)


def is_date(string):
    try:
        convert_to_date(string)
        return True
    except (ValueError, TypeError):
        return False


def transform_input(prompt, test, callback):
    while True:
        try:
            val = input(prompt)

            if test(val):
                return callback(val)
        except ValueError:
            print("Invalid input, try again.")
        else:
            print("Invalid input, try again.")
