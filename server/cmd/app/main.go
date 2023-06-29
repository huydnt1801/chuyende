package main

import (
	"fmt"
	"log"
	"os"

	"github.com/huydnt1801/chuyende/api/server"
	"github.com/huydnt1801/chuyende/internal/config"
	"github.com/urfave/cli/v2"
	"github.com/xo/dburl"
)

func main() {
	os.Setenv("SECRET_KEY", "something-very-secret-that-you-cannot-know")
	os.Setenv("MYSQL_HOST", "127.0.0.1")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_USERNAME", "chuyende")
	os.Setenv("MYSQL_PASSWORD", "password")
	os.Setenv("MYSQL_DATABASE", "chuyende")
	app := &cli.App{
		Name:  "Chuyen de",
		Usage: "",
		Action: func(cCtx *cli.Context) error {
			cfg := config.MustParseConfig()
			db, err := dburl.Open(cfg.DatabaseURI())
			if err != nil {
				return fmt.Errorf("failed to open db: %w", err)
			}

			accSrv := server.NewAccountServer(db)
			tripSrv := server.NewTripServer(db)
			srv, err := server.NewServer(db, accSrv, tripSrv)
			if err != nil {
				return err
			}
			return srv.Serve()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
