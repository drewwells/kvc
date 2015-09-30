package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/codegangsta/cli"
)

var cmds []cli.Command

func main() {

	app := cli.NewApp()
	app.Name = "Bolt DB Viewer"
	app.Usage = "Inspect a Bolt DB database"
	app.Commands = cmds
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db",
			Usage: "path to bolt DB file",
		},
	}
	app.Run(os.Args)
}

func init() {
	cmds = append(cmds, []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "list all buckets",
			Action: list,
		},
		cli.Command{
			Name:   "get",
			Usage:  "get all data in a bucket",
			Action: get,
		},
	}...)
}

func resolveDB(ctx *cli.Context) *bolt.DB {
	path := ctx.GlobalString("db")
	if len(path) == 0 {
		log.Fatal("Must pass --db")
	}
	db, err := bolt.Open(path, 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal("error opening bolt file", err)
	}
	return db
}

var ellipsis = []byte("...")
var maxKeySize = 31

func get(ctx *cli.Context) {
	args := ctx.Args()
	if len(args) != 1 {
		log.Fatal("you must select a bucket format: get bucket")
	}
	sBkt := args[0]
	var err error
	db := resolveDB(ctx)

	err = db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(sBkt))
		if bkt == nil {
			return fmt.Errorf("No bucket called: %s", sBkt)
		}
		fmt.Println("|  ------------ Key  -----------  | ---------------- Value ---------------------")
		indent := "                                    "
		return bkt.ForEach(func(k []byte, v []byte) error {
			skey := string(k)
			if len(skey) > maxKeySize {
				skey = skey[:28] + "..."
			}
			// attempt json unmarshalling of value
			var m map[string]interface{}
			err := json.Unmarshal(v, &m)
			if err == nil {
				// Indent formatting looked bad
				v, _ = json.MarshalIndent(m, indent, "  ")
			}
			fmt.Printf(" %32s : %s\n", skey, string(v))
			return nil
		})
	})
	if err != nil {
		log.Fatal(err)
	}
}

func list(ctx *cli.Context) {
	args := ctx.Args()
	_ = args
	var err error
	db := resolveDB(ctx)

	err = db.View(func(tx *bolt.Tx) error {
		err := tx.ForEach(func(name []byte, bkt *bolt.Bucket) error {
			_ = bkt
			fmt.Println(string(name))
			return nil
		})
		return err
	})

	if err != nil {
		log.Fatal(err)
	}
}
