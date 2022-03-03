# MT940 Server

## Database Connection

To connect to a remote database e.g GCP SQL Postgres, you can setup the cloud sql proxy as follow

`cloud_sql_proxy.exe -instances='rutela:us-central1:riaviatech'=tcp:127.0.0.1:8484`
