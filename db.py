import sqlite3
from contextlib import contextmanager


@contextmanager
def db():
    with sqlite3.connect("budget.sqlite") as connection:
        cursor = connection.cursor()
        yield cursor
