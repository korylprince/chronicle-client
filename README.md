[chronicle-client](https://github.com/korylprince/chronicle-client)

This client has been split from the [old repo](https://github.com/korylprince/chronicle) to be in its own repository.

The server may be found [here](https://github.com/korylprince/chronicle-server).

chronicle is for simple client/server data logging to a database.
The client will send the following data:
 
* uid of current user
* username of current user
* full name (GECOS) of current user
* serial number of computer
* [munki](http://munki.github.io/munki/) ClientIdentifier
* hostname of computer
* local IP address of computer

The server logs this to a database along with the time and IP of the client as seen by the server.

The client will only work on OS X, and has only been tested on 10.9. The server will probably run anywhere.

#Installing#

`go get github.com/korylprince/chronicle-client`

If you have any issues or questions, email the email address below, or open an issue at:
https://github.com/korylprince/chronicle-client/issues

#Usage#

Read the source. It's pretty simple and readable.

The following Enviroment Variables are configurable:

* Client:
    * `CHRONICLE_CHRONICLER` ( **required** )
        * Host:Port pair of server
    * `CHRONICLE_TLS`
        * http or https
    * `CHRONICLE_INTERVAL`
        * how often the client checks in specified in minutes. This value is skewed randomly to avoid flooding the server.

There are some sample build scripts in the `build` folder. Note that [pkggen](https://github.com/korylprince/pkggen) is required for `mkpkg.sh`.

#Copyright Information#

Copyright 2015 Kory Prince (korylprince at gmail dot com.)

This code is licensed under the same license go is licensed under with slight modification (my name in place of theirs.) If you'd like another license please email me.
