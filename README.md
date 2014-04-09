gomkafka
========

Per rsyslog external plugin requirements, we should read from stdin, use a single thread, and ensure no output is sent to stdout.

### Usage

```shell
$ bin/zookeeper-server-start.sh config/zookeeper.properties
$ bin/kafka-server-start.sh config/server.properties
$ gomkafka "client id" localhost:9092 monitoring
(CTRL-C to exit)
```

### Integration with Rsyslog

Add this to your `rsyslog.conf`

```shell
module(load="omprog")

if $rawmsg contains "arbitrary trigger" then
  action(type="omprog"
         binary="/path/to/gomkafka" --param1=\"client id\" --param2=localhost:9092 --param3=monitoring"
```

### Testing

```shell
$ echo "foo" | gomkafka "client id" localhost:9092 monitoring
```
