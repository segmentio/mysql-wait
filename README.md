# mysql-wait

This program will wait until mysql is available, or until a timeout has occurred.

It is meant to be used in CI builds to ensure that mysql can respond to commands before running the tests that depend on it.

# usage

Some builds might be able to get away without supplying any flags to `mysql-wait`, as there are sensible defaults.

The available options are:

    host     = flag.String("host", "127.0.0.1", "")
    port     = flag.Int("port", 3306, "")
    user     = flag.String("user", "root", "")
    password = flag.String("password", "", "")
    dbname   = flag.String("dbname", "mysql", "")
    sleep    = flag.Duration("sleep", time.Second, "")
    timeout  = flag.Duration("timeout", 30*time.Second, "")

# example

    curl -SsL https://github.com/segmentio/mysql-wait/releases/download/v0.0.2/mysql-wait -o mysql-wait
    chmod +x mysql-wait
    ./mysql-wait -timeout 60s

