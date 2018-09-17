package main

import (
        "flag"
        "fmt"
        "log"
        "os"
        "time"
)

func main() {
    filename := flag.String("file", "", "log file (default: stdout)")
    flag.Parse()
    outFile := os.Stdout

    if *filename != "" {
        file, err := os.OpenFile(*filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
        if err != nil {
            log.Fatal(err)
        }
        outFile = file
    }

    fmt.Fprintln(outFile, "Start logging...")
    t := time.Tick(5 * time.Second)
    for now := range t {
        fmt.Fprintf(outFile, "%v \n", now)
        time.Sleep(2 * time.Second)
    }

}
