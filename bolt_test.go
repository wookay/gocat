package gocat

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"regexp"
	"testing"
	"time"
)

type TestDB struct {
	*bolt.DB
}

func NewTestDB() *TestDB {
	path := "bolt.db"
	db, err := bolt.Open(path, 0666, nil)
	if err != nil {
		panic("cannot open db: " + err.Error())
	}
	return &TestDB{db}
}

var statsFlag = flag.Bool("stats", false, "show performance stats")

func truncDuration(d time.Duration) string {
	return regexp.MustCompile(`^(\d+)(\.\d+)`).ReplaceAllString(d.String(), "$1")
}

func (db *TestDB) Close() {
	// Log statistics.
	//if *statsFlag {
	db.PrintStats()
	//}

	// Check database consistency after every test.
	db.MustCheck()

	// Close database and remove file.
	defer os.Remove(db.Path())
	db.DB.Close()
}

func (db *TestDB) MustCheck() {
	db.Update(func(tx *bolt.Tx) error {
		// Collect all the errors.
		var errors []error
		for err := range tx.Check() {
			errors = append(errors, err)
			if len(errors) > 10 {
				break
			}
		}

		// If errors occurred, copy the DB and print the errors.
		if len(errors) > 0 {
			fmt.Printf("consistency check failed (%d errors)\n", len(errors))
			for _, err := range errors {
				fmt.Println(err)
			}
		}
		return nil
	})
}

// PrintStats prints the database stats
func (db *TestDB) PrintStats() {
	var stats = db.Stats()
	fmt.Printf("[db] %-20s %-20s %-20s\n",
		fmt.Sprintf("pg(%d/%d)", stats.TxStats.PageCount, stats.TxStats.PageAlloc),
		fmt.Sprintf("cur(%d)", stats.TxStats.CursorCount),
		fmt.Sprintf("node(%d/%d)", stats.TxStats.NodeCount, stats.TxStats.NodeDeref),
	)
	fmt.Printf("     %-20s %-20s %-20s\n",
		fmt.Sprintf("rebal(%d/%v)", stats.TxStats.Rebalance, truncDuration(stats.TxStats.RebalanceTime)),
		fmt.Sprintf("spill(%d/%v)", stats.TxStats.Spill, truncDuration(stats.TxStats.SpillTime)),
		fmt.Sprintf("w(%d/%v)", stats.TxStats.Write, truncDuration(stats.TxStats.WriteTime)),
	)
}

func TestBolt(t *testing.T) {
	// Open a data file.
	db := NewTestDB()
	path := db.Path()
	defer db.Close()
	log.Print(path)
}
