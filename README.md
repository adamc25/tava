This project is designed to work with a local PostGreSQL database.

The connection string is currently hardcoded in `src/server/internal/database/database.go`

I set it up to use no password. Even with no password, a role is required. With that in mind, this is the process I followed to get it working for me using a role called cragg.

`CREATE ROLE cragg WITH LOGIN;`

and to create the database

`CREATE DATABASE "tava-health"`

This command will need to be run on the query tool for the tava-health database where you can replace cragg with your specific username. This may need to be run depending on which connection string you setup.

`GRANT CREATE ON SCHEMA public TO cragg`

And finally, this is what my config file looks like in Windows for PostgreSQL (located at `Program Files/PostgreSQL/17/data/pg_hba.conf`)

```
# "local" is for Unix domain socket connections only
local   all             all                                     trust

# IPv4 local connections:
host    all             all             127.0.0.1/32            trust

# IPv6 local connections:
host    all             all             ::1/128                 trust
```

To seed the database when everything is configured (including installing Go if you don't have it setup), cd into `src/server/cmd` and run `go run seed.go`

To start the Go server cd into `src/server/cmd` and run `go run server.go`

To start the React server, make sure to run `npm install` and then `npm run start`.

You should now have a functioning server and React app.
