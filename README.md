gomkafka
========

[![Build Status](https://travis-ci.org/jeffchao/gomkafka.svg?branch=master)](https://travis-ci.org/jeffchao/gomkafka)

Gomkafka is a rsyslog external plugin that collects rsyslog logs and ships them over to a Kafka cluster in near realtime.

## Usage

```shell
$ bin/zookeeper-server-start.sh config/zookeeper.properties
$ bin/kafka-server-start.sh config/server.properties
$ gomkafka "client id" localhost:9092 monitoring
(CTRL-C to exit)
```

## Building

You will first need godep.

```shell
go get github.com/tools/godep
```

Then you can build as normal.

```shell
godep restore
go install
```

Once you have the binary, you can put it anywhere as long as your `rsyslog.conf` correctly points to it and has correct permissions.

## Integration with Rsyslog

Add this to your `rsyslog.conf`

```shell
module(load="omprog")

if $rawmsg contains '[monitoring]' and ($syslogseverity == 6 or $syslogseverity == 5) then {
  action(type="omprog" binary="/path/to/gomkafka client_id localhost:9092 monitoring")
}
```

The `$rawmsg` is a default rsyslog property representing the raw message. The statement "if `property` contains `value` then..." conditionally executes gomkafka where contains must exactly match the `value`. It  cannot be a regex.

Additionally, you may filter on severity, per RFC 3164, using the `$syslogseverity` or `$syslogseverity-text` property.

Now in this example, in your app-tier code, you will have a logger that logs to syslog with log level `info` or `notice` with some text `[monitoring]`. If those conditions are met, then gomkafka will pick up the log event and ship it over to a Kafka cluster.

```
Numerical         Severity
  Code

    0       Emergency: system is unusable
    1       Alert: action must be taken immediately
    2       Critical: critical conditions
    3       Error: error conditions
    4       Warning: warning conditions
    5       Notice: normal but significant condition
    6       Informational: informational messages
    7       Debug: debug-level messages
```

## Testing

```shell
$ echo "foo" | gomkafka client_id localhost:9092 monitoring
```

## Author

Jeff Chao, @thejeffchao, http://thejeffchao.com
