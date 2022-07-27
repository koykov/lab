package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	drv  = flag.String("driver", "", "SQL driver")
	dsn  = flag.String("dsn", "", "DSN string")
	mlt  = flag.Int("conn-max-lt", 0, "Connection max lifetime duration (seconds)")
	ping = flag.Uint("ping", 10, "Pings count")
)

func init() {
	flag.Parse()
	if len(*drv) == 0 {
		log.Fatal("driver is mandatory")
	}
	if len(*dsn) == 0 {
		log.Fatal("DSN is mandatory")
	}
}

func main() {
	var (
		dbi *sql.DB
		err error
	)
	if dbi, err = sql.Open(*drv, *dsn); err != nil {
		log.Fatal(err)
	}
	dbi.SetConnMaxLifetime(time.Duration(*mlt) * time.Second)

	for i := uint(0); i < *ping; i++ {
		if err = dbi.Ping(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 5)
		log.Printf("ping %d\n", i)
	}
	log.Println("done")
}
