package main

import (
	_ "github.com/lib/pq"
    "database/sql"
    "fmt"
    "log"
)

type Book struct {
    isbn  string
    title  string
    author string
    price  float32
}

func main2() {
    db, err := sql.Open("postgres", "postgres://postgres:s3cr3t@localhost/bookstore")
    if err != nil {
        log.Fatal(err)
    }

    rows, err := db.Query("SELECT * FROM myschema.books")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    bks := make([]*Book, 0)
    for rows.Next() {
        bk := new(Book)
        err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
        if err != nil {
            log.Fatal(err)
        }
        bks = append(bks, bk)
    }
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }

    for _, bk := range bks {
        fmt.Printf("%s, %s, %s, £%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
    }
}