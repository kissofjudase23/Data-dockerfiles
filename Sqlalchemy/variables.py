
import os
from pprint import pprint as pp


def get_os_env(key):
    if key not in os.environ:
        raise ValueError()

    return os.environ[key]


class DbVars(object):
    HOST = get_os_env("DB_HOST")
    PORT = int(get_os_env("DB_PORT"))
    MAX_CONNECTIONS = int(get_os_env("DB_MAX_CONNECTIONS"))
    DATABASE = get_os_env("DB_DATABASE")
    USER = get_os_env("DB_USER")
    PASSWORD = get_os_env("DB_PASSWORD")
    ENGINE = 'InnoDB'
    DRIVER = 'pymysql'
    CHARSET = 'utf8mb4'
    COLLATE = 'utf8mb4_unicode_ci'


if __name__ == '__main__':
    output = {"DbVars": DbVars}

    pp(output)
