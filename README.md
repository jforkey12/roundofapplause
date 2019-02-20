# roundofapplause

How to run
==========
Download and install MongoDB:
https://www.mongodb.com

You'll need ports 558 and 27017 available on localhost.

create a directory and start the mongoDB service by executing the following on a command prompt:

>md /data/db
>"C:\Program Files\MongoDB\Server\4.0\bin\mongo.exe"

Create the database user for the connection via mongoDB shell:
db.createUser({user:"admin1",pwd: "qwerty123456!",roles: [{ role: "readWrite", db:"applause"}]})

start one of the services:
app_small.exe or app_scale.exe

3rd party dependencies:

github.com/globalsign/mgo

-The MongoDB (community supported branch) driver for Go.

github.com/gorilla/mux

-For quickly spinning up routes for rest endpoints

Querying
======== 
Here is an example query string:

http://localhost:558/applause/v1/users?country=US&devices=iPhone 4,iPhone 4S

Use "?" to start the query string.

Use "&" to indicate different filter parameters.

Use "," to specify multiple devices or countries.

NOTE: "ALL" follows the same implicit rules of typical RESTAPIs, that is, specifying no filter is equivalent to all. 

Example:
http://localhost:558/applause/v1/users
-- would be the equivalent of ALL countries and ALL devices


Testing
========
I tested on the small dataset (app_small.exe) and also created a scale app for test (app_scale.exe). The scale app will initalize the database with 1M bugs and 100K users with random data and each having 2 random devices.  The app_small.exe will initialize the database with the CSV data.

You can use any REST client or curl to test with.  It's important to note, the response itself is much faster than the client shows, the client struggles to handle the size of the response payloadon the scale app.

Caveats
========
I didn't write unit tests. 

I didn't comment the code. 

 It's relatively simple code and reads pretty easy. 

Normally, my code is commented appropriately. 

I didn't use constants, there are many places where it makes sense to use constants in this code.

I drop a lot of errors on the floor throughout the code, simply due to time constraints.


Building
========
You'll need go installed with GOBIN and GOROOT set correctly to build.  Just run "go build" from ../app directory to generate the binary.  It will build on windows or linux.