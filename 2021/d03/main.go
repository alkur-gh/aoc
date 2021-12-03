package main

import (
    "fmt"
    "day3/diagnostic"
)

func main() {
    path := "./tests/files/input.txt"
    report := diagnostic.ReadReportFromFile(path)
//    fmt.Println(report)
    fmt.Println(report.PowerConsumption())
    fmt.Println(report.LifeSupportRating())
}
