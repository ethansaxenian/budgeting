import psycopg2
from psycopg2.extras import RealDictCursor

from core.config import settings


def get_db():
    with (
        psycopg2.connect(
            user=settings.DB_USER,
            password=settings.DB_PASSWORD,
            database=settings.DB_NAME,
            port=settings.DB_INNER_PORT,
            host=settings.DB_HOST,
        ) as connection,
        connection.cursor(cursor_factory=RealDictCursor) as cursor,
    ):
        yield cursor
