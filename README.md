# URL Shorter
A simple and fast url shortener written in Go.

> This project is still early in development, **not** suited for production

## Features
- Incremented unique short urls
- No JS
- Easy setup
- REST API

## Configuration
You can configure the app using environment variables.

| Name               | Description                                 | Default |
| :----------------- | :------------------------------------------ | :------ |
| SECRET_KEY         | base64 encoded key                          |         |
| REQUIRE_LOGIN      | Whether to require login for short creation | false   |
| ALLOW_NEW_ACCOUNTS | Whether to allow new account registration   | true    |
| DB_SQLITE_PATH     | Path to SQLite database                     |         |

## License
This project is Copyright (c) 2022 Leo Spratt, licences shown below:

Code

    AGPL-3 or any later version. Full license found in `LICENSE.txt`
