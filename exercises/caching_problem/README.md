Create a library for caching user information to increase database throughput and availability.
Every request for user data should return the user information while simultaneously
taking care of database throughput.

Why:
Let's say we have a bunch of user information (identified by id) inside our main database.
Simultaneously there's thousands of request per second for user information.
We can consider that the database is a bottleneck. To avoid that we need a cache mechanism.

Notes:
You can use any language, library, database and framework - anything you want and think is best for this case.

Example:
If within 1000 requests there are 100 unique user ids then there should be only a maximum 100 requests into 
the database but all 100 requests should get a response with user data.
