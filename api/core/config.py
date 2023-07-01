from pathlib import Path, PosixPath

from pydantic import BaseSettings

PROJECT_ROOT = Path = Path(__file__).resolve().parent.parent


class Settings(BaseSettings):
    PROJECT_NAME: str = "Budgeting API"
    PROJECT_DESCRIPTION: str = "A budgeting api"
    API_PREFIX: str = "/api"
    API_ROOT_DIR: PosixPath = PROJECT_ROOT / "api"
    AUTHOR_NAME: str = "Sample Name"
    AUTHOR_EMAIL: str = "sample_email@provider.tld"
    LICENSE: str = "MIT"
    DB_NAME: str = "postgres"
    DB_USER: str = "postgres"
    DB_PASSWORD: str = "postgres"
    DB_INNER_PORT: int = 5432
    DB_HOST: str = "db"

    class Config:
        env_file = PROJECT_ROOT / ".env"
        case_sensitive = True


settings = Settings()
