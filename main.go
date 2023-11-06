package main

import (
  "fmt"
  "database/sql"
  "log"
  "github.com/go-sql-driver/mysql"
)

const (
  pass = "S-qGoL3-4az"
)

type (
  Paper struct {
    Id      int64
    Title   string
    Content string
    Status  int
  }
)

var db *sql.DB

func (paper Paper) String() string {
  return fmt.Sprintf("Paper {ID: %d, Status: %d, Title: \"%s\", Content:\n  \"%s\"}",
    paper.Id, paper.Status, paper.Title, paper.Content)
}

func connect() {
  cfg := mysql.Config{
    User:   "sqlgo",
    Passwd: pass,
    Net:    "tcp",
    Addr:   "127.0.0.1:3306",
    DBName: "sqlgo",
  }

  var err error
  db, err = sql.Open("mysql", cfg.FormatDSN())
  if err != nil {
    log.Fatal(err)
  }

  pingErr := db.Ping()
  if pingErr != nil {
    log.Fatal(pingErr)
  }
}

func test() {
  papers := make([]Paper, 0, 10)

  const minStatus = 1;

  rows, err := db.Query(`
    SELECT *
    FROM papers
    WHERE status >= ?
  `, minStatus)

  if err != nil {
    log.Fatal(err)
  }

  defer rows.Close()

  for rows.Next() {
    var paper Paper
    err := rows.Scan(&paper.Id, &paper.Title, &paper.Content, &paper.Status)
    if err != nil {
      log.Fatal(err)
    }
    papers = append(papers, paper)
  }

  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }

  fmt.Printf("Total %d paper(s) loaded.\n", len(papers))

  for _, paper := range papers {
    fmt.Println(paper)
  }
}

func main() {
  connect()
  fmt.Println("Connected.")
  test()
}

