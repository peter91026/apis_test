set DB_REPLICAS=0
set SOURCES_HOST=localhost
set SOURCES_PORT=5432
set SOURCES_USER=aa815
set SOURCES_PASSWORD=peter85220
set SOURCES_DATABASE=test
set SOURCES_SSLMODE=disable

set PUBLIC = ""
set PRIVATE = ""

@REM echo %SOURCES_DATABASE%
air.exe -c .air.toml

