# zoochecker
zoochecker uses zookeeper 4lw commands to monitor the health of a zookeeper ensemble.  it uses 'ruok' to check each node in the ensemble.  it will return the status of each node. It uses 'mntr' to check the numbers of leaders, followers and synced followers.


```bash
zoochecker localhost:21811 localhost:21812 localhost:21813
zoochecker version:  e35ffa96c0bbd6649dea6ff60aecefeacd1d8030
{"level":"info","version":"e35ffa96c0bbd6649dea6ff60aecefeacd1d8030","time":"2024-10-18T11:06:26-04:00","message":"Cluster: {Nodes:[{Host:localhost Port:21811 Timeout:5} {Host:localhost Port:21812 Timeout:5} {Host:localhost Port:21813 Timeout:5}]}"}
{"level":"info","version":"e35ffa96c0bbd6649dea6ff60aecefeacd1d8030","time":"2024-10-18T11:06:26-04:00","message":"Node localhost:21811 is OK"}
{"level":"info","version":"e35ffa96c0bbd6649dea6ff60aecefeacd1d8030","time":"2024-10-18T11:06:26-04:00","message":"Node localhost:21812 is OK"}
{"level":"info","version":"e35ffa96c0bbd6649dea6ff60aecefeacd1d8030","time":"2024-10-18T11:06:26-04:00","message":"Node localhost:21813 is OK"}
{"level":"info","version":"e35ffa96c0bbd6649dea6ff60aecefeacd1d8030","time":"2024-10-18T11:06:26-04:00","message":"Cluster is OK"}
```

## contributing

have docker and docker-compose installed in order to run a local zookeeper ensemble.  start the ensemble with `make compose-up` and stop it with `make compose-down`.  check the status of the ensemble with `echo ruok | nc localhost 21811` or `echo ruok | nc localhost 21812` or `echo ruok | nc localhost 21813`.

```bash
make compose-up
echo ruok | nc  localhost 21811
echo ruok | nc  localhost 21812
echo ruok | nc  localhost 21813
make compose-down
```


'make static' will run static checkers and unit tests.