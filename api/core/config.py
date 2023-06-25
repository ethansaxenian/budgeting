import os
from pathlib import Path, PosixPath

from pydantic import BaseSettings

PROJECT_ROOT = Path = Path(__file__).resolve().parent.parent

class Settings(BaseSettings):
    PROJECT_NAME: str = "Budgeting API"
    PROJECT_DESCRIPTION: str = "A budgeting api"
    API_PREFIX: str = "/api"
    API_ROOT_DIR: PosixPath = PROJECT_ROOT / "api"
    AUTHOR_NAME: str = "Ethan Saxenian"
    AUTHOR_EMAIL: str = "ethansaxenian+github@proton.me"
    LICENSE: str = "MIT"
    DB_NAME: str = "postgres"
    DB_USER: str
    DB_PASSWORD: str
    DB_PORT: int = os.getenv("DB_INNER_PORT")
    DB_HOST: str = "db"

    class Config:
        env_file = PROJECT_ROOT / ".env"
        case_sensitive = True


settings = Settings()
