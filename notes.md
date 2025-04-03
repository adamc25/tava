I setup an actual db since that seemed like an obvious step to take.

I was choosing between MySQL and PostgreSQL. I chose the latter becasue the timestamp in the provided data is not considered a valid timestamp by MySQL but it is by PostgreSQL.

I could have chose to use an ORM but tried to keep it simple so stuck with raw sql.

My database is setup to use no password. It still needs a role. The `cragg` part of the connection string in `database.go#GetConnectionString` represents the username. This bit might need a little bit of configuration on your end.

For me, it was simply executing the following in pgadmin. Replace `cragg` with your specific username.:

`CREATE ROLE cragg WITH LOGIN;`

and to create the database

`		
CREATE DATABASE "tava-health"
`

This command will need to be run on the query tool for the `tava-health` database where you can replace `cragg` with your specific username. This may need to be run depending on which connection string you setup.

`GRANT CREATE ON SCHEMA public TO cragg`

To seed the database when everything is configured (including installing Go if you don't have it setup), cd into `src/server/cmd` and run `go run seed.go`

To start the Go server cd into `src/server/cmd` and run `go run server.go`

To start the `React server`, make sure to run `npm install` and then `npm run start`.

With that in mind, I think it would be easier to say what I didn't do instead of listing what I did do. I probably split my time about 50/50 between the React frontend and the Go backend. I have already spent quite a bit of time on it so didn't include a test suite or a themable UI. I could talk to you all day long about how to test in Angular, mocking depenencies, etc. As far as I know, none of that applies in React-land. I'm about as equally knowledable in testing in Go as I am in testing React.

Some topics of interest that you might want to bring up:

```
1) Displaying multiple errors vs a single error (when multiple errors occurred)
2) What validation should we do on UI forms and how do we handle setting / unsetting associated state?
3) Supporting more than just equals and wildcard matching for search terms (ie, spawniwng a datepicker to select a date range)
4) Using error codes instead of error messages. This keeps devops team happy.
5) The backend is setup to offer pagination but there isn't anything setup in the React code.
6) What other kinds of dates might be interesting to add to the employment_date table (resigned_on, laid_off_on, etc)
7) I structured the sql in such a way that searches are possible on any column from any table. The possible columns to search on is controlled by what exists in the search table. This means we can setup any indexes, db constraints, etc before adding a new searchabe / sortable field to the database.
8) EditEmployee.tsx could probably be split up a bit more
9) ORM vs Raw SQL approach
10) The table schemas are viewable at internal/database/*.schema.go
11) The available public api is viewable at cmd/server.go (CORS is setup on every http handler)
12) The React code has the server url hardcoded and the connection string for the database is hardcoded. In a real world scenario, these would be driven by something else (environment variables, configuration files, etc)
13) The server code and frontend code is in the same Git repo. This may or may not be desirable based on the makeup of the team, individual responsibilities, etc.
14) net/http is the only "router" setup right now for handling incoming http requests. It doesn't natively support params in the url such as /employee/:id. This is probably a valid use case (maybe not for that specific endpoint but likely for others) so should almost certainly be supported in a production Go backend. This is something to look into to see what options there are.
15) There isn't much in the way of a standardized folder structure for Go so that seems like a team decision. I structured it based on how I structured a personal Go project. I looked into what was considered to be a standard folder layout at the time and cmd / internal seemed to be the most common setup.
16) I structured the css such that there is one css file per component with any common styles going in styles.css. In reality, this would be better accomplished by something similar but with scss files instead. I like SCSS (or SASS or LESS for that matter) since it allows you to structure your style rules according to the structure of the html for that component. It doesn't get better than that!
17) For editing an employee, I added end date (not editable), employment status (active vs inactive, could potentially have other values in the future), and avatar url. If you display a valid http address for an image, it will now display in the list view. I found a default image of an empty head which is what gets displayed in the scenario that there is no avatar url setup / it's a faulty address.
18) Currently, there is no limit on how long a quote someone can add. It would probably make sense to restrict this in some way which would result in changing the type of that column VARCHAR to reflect that.
19) My Go code has a mix of short and long variable names. Go enthusiasts REALLY REALLY love single, short, nonsensical variable names. I'm personally not a fan but would rather not upset the Go crowd so tried to follow that paradigm when it was possible without obfuscating what the variable was representing.
20) I tried to follow standard pratice around working with resources. DELETE for deleting data. PUT for updating existing records. GET for fetching resources except in cases where we need to attach a rather complex looking object (POST in those cases). There's serious considerations around URI string length which has caught me out in the past. Sometimes, you have to violate norms.
21) I don't believe I added any additional dependencies to the React project. The only additional dependency I included was a database driver for working with Postgres in the Go codebase. The rest of the imports are local packages in the project or packages from the standard library in Go.
```

Hopefully, that wall of text gives everyone 1 or more talking points!