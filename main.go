package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		host     = flag.String("host", "127.0.0.1", "")
		port     = flag.Int("port", 3306, "")
		user     = flag.String("user", "root", "")
		password = flag.String("password", "", "")
		dbname   = flag.String("dbname", "mysql", "")
		sleep    = flag.Duration("sleep", time.Second, "")
		timeout  = flag.Duration("timeout", 30*time.Second, "")
	)
	flag.Parse()
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			err := func() error {
				var dsn string
				if *password == "" {
					dsn = fmt.Sprintf("%s@tcp(%s:%d)/%s", *user, *host, *port, *dbname)
				} else {
					dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", *user, *password, *host, *port, *dbname)
				}
				db, err := sql.Open("mysql", dsn)
				if err != nil {
					return err
				}
				err = db.Ping()
				return err
			}()
			if err != nil {
				fmt.Println(err)
				time.Sleep(*sleep)
				continue
			}
			return
		}
	}()

	select {
	case <-time.After(*timeout):
		fmt.Println("mysql-wait: timed out trying to connect")
		os.Exit(1)
	case <-done:
		fmt.Println("mysql-wait: connected")
	}
}
