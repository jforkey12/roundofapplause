# roundofapplause

========
WINDOWS
========
Download and install MongoDB:
https://www.mongodb.com

create a directory and start the mongoDB service by executing the following on a command prompt:

>md /data/db
>"C:\Program Files\MongoDB\Server\4.0\bin\mongo.exe"

Create the database user for the connection via mongoDB shell:
db.createUser({user:"admin1",pwd: "qwerty123456!",roles: [{ role: "readWrite", db:"applause"}]})

Start the mongodb service on 27017 and create a user admin1 / qwerty123456!
You'll need ports 555, 556, and 27017 available on localhost.


To get 
The dataset was unnecessarily relational.  When you boil it down, the only things that matter here are users and bugs.
Bugs are found on a particular device.
Users have a set of devices.

3rd party dependencies:
github.com/globalsign/mgo
-The MongoDB (community supported branch) driver for Go.
github.com/go-resty/resty
-To quickly spin up rest endpoints and wrap responses
github.com/gorilla/mux
-For quickly spinning up routes for said endpoints



You used a non-relation database to solve a coding changing that was obviously relational?
Yes.

Why?
The dataset was incredibly small and the requirements doc doesn't say the solution needs to scale.  All good solutions scale, and this solution scales too.  There are two blocks of code in .  Uncomment them and comment out the csv parsing code chunks to scale this to 300,000 users and 500,000 bugs.

Caveats
I didn't write unit tests.  It's a coding excerise.
I didn't comment the code.  It's relatively simple code and reads pretty easy.  Normally, my code is commented appropriately. 
I didn't use constants, there are many places where it makes sense to use constants in this code.
I drop a lot of errors on the floor throughout the code, simply due to time constraints.