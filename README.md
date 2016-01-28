# fsevents to docker-machine

Some file watchers running under docker-machine don’t pick up the file events under OSX. This attempts to hack an event on the container’s file by listening to fsevents on the host machine & “passing” to the guest machine’s container.
