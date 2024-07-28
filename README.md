# BTC Alert System

To run this app,first git clone this repo. Then pull a postgres docker image inside your machine. Use the following command `docker run --name some-postgres -e POSTGRES_PASSWORD=pranav -d postgres`. Then inside the root of the directory run `make run`. This will start the app.Also run `soda migrate up` to do the database migrations.If soda is not available on your machine run `brew install gobuffalo/tap/pop` and you can run the previous command to do the database migrations. 

## EndPoints


**/register** - Registers an user to the database.

**/login** - Logins a user to the database. Without loggin in one cannot do anything.

**/alerts/create** - Creates an alert in the database. Make sure the request body has a Price in it.

**/alerts/delete** - Marks an alert as deleted in the status section. Make sure to give the AlertID in the request body.

**/alerts/list** - Lists all the alerts


