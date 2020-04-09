# mysql-wait

This program will wait until mysql is available, or until a timeout has occurred.

It is meant to be used in CI builds to ensure that mysql can respond to commands before running the tests that depend on it.

# example

    curl -SsL https://github.com/segmentio/mysql-wait/releases/download/v0.0.2/mysql-wait -o mysql-wait
    chmod +x mysql-wait
    ./mysql-wait -timeout 60s -user root

