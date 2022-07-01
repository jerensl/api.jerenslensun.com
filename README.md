# Jerens Personal Web Api

[![codecov](https://codecov.io/gh/jerensl/api.jerenslensun.com/branch/main/graph/badge.svg?token=RIDDKEIQW8)](https://codecov.io/gh/jerensl/api.jerenslensun.com) ![Continuous Integration](https://github.com/jerensl/api.jerenslensun.com/actions/workflows/ci.yml/badge.svg) ![Continuous Deployment](https://github.com/jerensl/api.jerenslensun.com/actions/workflows/cd.yml/badge.svg)

This Restful Api is notification api for my [personal web app](https://www.jerenslensun.com/)

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`SERVICE_ACCOUNT_FILE`: Firebase service account

`API_KEY`: Api key header token for authorization

`CORS_ALLOWED_ORIGINS`

`SQLITE_DB`: Sqlite Database file location

## Run Locally

Clone the project

```bash
  git clone https://github.com/jerensl/api.jerenslensun.com.git
```

Go to internal folder in the project directory

```bash
  cd api.jerenslensun.com/internal
```

Start the server

```bash
  go run main.go
```


## Features

- Status Notification
- Subscribe Notification
- Unsubscribe Notification
- Send Notification

## Documentation

[Documentation](https://api.jerenslensun.com/docs)

## C4 Diagram

![C4 Diagram](/tools/c4/out/view-notification.png)

## License

[MIT](https://choosealicense.com/licenses/mit/)

