from pathlib import Path

from pydantic import BaseSettings


class Settings(BaseSettings):
    PROJECT_NAME: str = "Budgeting API"
    PROJECT_DESCRIPTION: str = "A budgeting api"
    API_PREFIX: str = "/api"
    ROOT_DIR: Path = Path(__file__).resolve().parent.parent
    AUTHOR_NAME: str = "Ethan Saxenian"
    AUTHOR_EMAIL: str = "ethansaxenian+github@proton.me"
    LICENSE: str = "MIT"
    DB_NAME: str = "postgres"
    DB_USER: str = "postgres"
    DB_PASSWORD: str = "postgres"
    DB_PORT: int = 5432
    DB_HOST: str = "db"

    class Config:
        env_file = ".env"
        case_sensitive = True


settings = Settings()
