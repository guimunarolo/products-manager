from prettyconf import config


class Settings:
    DATABASE_URL = config("DATABASE_URL")
    CALCULATOR_URL = config("CALCULATOR_URL")


settings = Settings()
