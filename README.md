# Vrecipes

A recipe management website.

## Fork and updating from the main repo

This projekt is a fork of `https://github.com/ViddeM/vrecipes` and during development many changes were needed at both the main repo and this. So the change was made to the main repo to then be imported using the fork feature.

Here is how you update this repository with new desired commits from the main repo.

1. Below the `Code`, i.e. download button, _sync fork_ is present and can be pressed.
2. If git can auto-merge the commits then a green button exists which can update TasteIT with the changes automatically.
3. If there are conflicts Github only displays the `discard XX commits` and the merge has to be done automatically.
4. Run `git merge --no-ff upstream/master` and solve the merge conflicts. IntelliJ has an excellent merge-tool.
5. Commits and push.
6. Profit!

## Development setup

To setup the development of the project there are some things that are necessary to be setup.

### Frontend

For the frontend the following steps are necessary:

1. Install the node dependencies in the `frontend/` folder (e.g. whilst inside of the `frontend/` folder run `yarn` or equivalent command).
1. Be aware that there is a `.env.development` file in the `frontend/` folder, however it should work out of the box on any `linux` based system.
1. In the **project root** folder, run `docker compose up`.

### Backend

The backend is not included into the `docker compose` to simplify developing the backend without having to restart everything else.  
The steps to setup the backend is as follows (all of these assume that you are inside of the `backend/` folder):

1. Copy the `.env.example` file to `.env`, an explaination of all the fields in this file can be found [below](#environment-variables).
2. Setup the Oauth2 login for Gamma.
3. Run the main method in `backend/cmd/tasteit/main.go`.

### Makefile

In the root folder there is also a Makefile with the following commands:

- `mock` (also the default): inserts mock values into the database.
- `clear-db`: `clean` is an alias for this. Clears the database (completely!) will require migrations to be re-run (i.e. restarting the backend).
- `clean`: Alias for `clear-db`.
- `new-migration mig_name_arg=*insert-migration-name*`: creates a new migration with the specified name.
- `run-migrations`: runs all migrations.
- `reset`: Perform `clean`, `run-migrations`, `mock` in that order.

### Migrations

To update the schema for the database you need access to the migrations.

1. `go get "github.com/golang-migrate/migrate/v4"`
2. Then run `make new-migration mig_name_arg=*insert-migration-name*`

## Environment variables

The environment variables that can / have to be specified for the project.
Note that for the moment ALL variables must exist / be non-empty to start the project.

### Frontend

- `NEXT_PUBLIC_BASE_URL`: Url to the backend seen from the server-side nextjs docker container, default is `http://host.docker.internal:5000/api` which (together with the `docker-compose`) specifies the host machine on port `5000` where the backend should exist in development. In production this should be set to the domain the website is hosted on.

### Backend

- `db_user` (string): Username for the database.
- `db_password` (string): Password for the database.
- `db_name` (string): Name of the database.
- `db_host` (string): Host address for the database.
- `reset_db` (boolean): Whether to clear the database on startup or not.
- `image_folder` (string): Path to where uploaded images should be stored.
- `secret` (string): Secret used to encrypt session cookies with.
- `GIN_MODE` (string): The mode for the `Gin` framework, see [github](https://github.com/gin-gonic/gin).
- `PORT` (integer): The port to host the backend on.

#### Gamma

Superadmin login: username=`admin` password=`password`.

For development, the client and secret is a bit funky. The default values wouldn't work so I set up a new client in gamma and retrieved its secret and clientId.

- `GAMMA_AUTHORIZATION_URI`: The uri to Gamma endpoint to login. (should be **gamma backend url**`/api/oauth/authorize`)
- `GAMMA_REDIRECT_URI`: The uri to redirect to with the response from Gamma (should be **frontend base url**`/api/auth/account/callback`)
- `GAMMA_TOKEN_URI`: (should be **gamma backend url**`/api/oath/token`)
- `GAMMA_ME_URI`: The uri to fetch the me user object from gamma. (should be **gamma backend url**`/api/users/me`)
- `GAMMA_SECRET`: The secret that gamma has stored for tasteIT, retrieved when setting up the oauth2 service. (example BefJFlmvJJjWTGPmpXWQXIpD6jzbSYAiwJRLqyUfUwrepd3gn4MGYAvROsZKQBVTiMapoRiDkRY)
- `GAMMA_CLIENT_ID`: The client id that gamma has stored for tasteIT, retrieved when setting up the oauth2 service. (example meY2VSxhZIxtCCeKgJ7jH3Odli3UmmcvrwiELhtSsvzZ1bMt33E2QpD8ctHR7CJMBTBKRismZSX)
- `GAMMA_LOGOUT_URL`: The url to log the user out. (should be **gamma backend url**`/api/logout`)
