# nfs-web

This is a small little project aming to provide a REST interface to the /etc/exports file.
It allows to mangage linux nfs shares via a http.
So far it is just a simple proof of concept that supports

- Adding a new share by POST request
- Deleting a share by DELETE request
- Getting all existing shares by GET request
- Getting the current `/etc/exports` configuration via GET

It does not (yet) provide:

- authentication/authorization
- a web ui
- any nfs functionality what so ever (purely manages the nfs config)
- a (working) dockerfile. (still needs the actual nfs server added)
