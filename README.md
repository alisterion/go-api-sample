API
---------------------------

#### Configure env:

Copy .env file:

    cp .env.sample .env

and configure .env file

#### Create database:

Change system user to postgres:

    sudo su postgres

Read variables from .env file:

    (source .env && eval "echo \"$(cat tools/sql/create_db.sql)\"") | psql

#### Run server:

    go build main.go
    ./main
 