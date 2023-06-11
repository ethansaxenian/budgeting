import sqlite3

from core.config import settings


def get_db():
    with sqlite3.connect(settings.DB_PATH, check_same_thread=False) as connection:
        connection.row_factory = sqlite3.Row
        yield connection.cursor()
