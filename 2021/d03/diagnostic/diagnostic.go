package diagnostic

import (
    "os"
    "bufio"
    "strconv"
)

type DiagnosticReport struct {
    recordCount int
    oneBitCount map[int]int
    maxLength int
    records []string
}

type DiagnosticRecord struct {
    line string
}

func NewDiagnosticReport() DiagnosticReport {
    return DiagnosticReport{0, make(map[int]int), 0, make([]string, 0)}
}

func (report *DiagnosticReport) Apply(record DiagnosticRecord) {
    report.recordCount++
    report.records = append(report.records, record.line)
    n := len(record.line)
    if n > report.maxLength {
        report.maxLength = n
    }
    for i, b := range record.line {
        if b == '1' {
            report.oneBitCount[n - 1 - i]++
        }
    }
}

func (report *DiagnosticReport) LifeSupportRating() int {
    oxygen := oxygenGeneratorRating(report.records, report.maxLength)
    co2 := co2ScrubberRating(report.records, report.maxLength)
    return oxygen * co2
}

func (report *DiagnosticReport) PowerConsumption() int {
    commonBits := report.mostCommonBits()
    gamma := 0
    for key, value := range commonBits {
        gamma += (value << key)
    }
    epsilon := gamma ^ ((1 << report.maxLength) - 1)
    return gamma * epsilon
}

func ParseDiagnosticRecord(line string) DiagnosticRecord {
    return DiagnosticRecord{line}
}

func ReadReportFromFile(path string) *DiagnosticReport {
    f, err := os.Open(path)
    if (err != nil) {
        panic(err)
    }
    reader := bufio.NewScanner(f)
    report := NewDiagnosticReport()
    for reader.Scan() {
        line := reader.Text()
        record := ParseDiagnosticRecord(line)
        report.Apply(record)
    }
    return &report
}

func (report *DiagnosticReport) mostCommonBits() map[int]int {
    threshold := (report.recordCount + 1) / 2
    res := make(map[int]int)
    for key, value := range report.oneBitCount {
        if value >= threshold {
            res[key] = 1
        } else {
            res[key] = 0
        }
    }
    return res
}

func mostCommonBits(records map[string]bool) map[int]byte {
    oneCount := make(map[int]int)
    for record, _ := range records {
        n := len(record)
        for i, b := range record {
            if b == '1' {
                oneCount[n - 1 - i]++
            }
        }
    }

    threshold := (len(records) + 1) / 2
    res := make(map[int]byte)
    for key, value := range oneCount {
        if value >= threshold {
            res[key] = 1
        } else {
            res[key] = 0
        }
    }

    return res
}


func oxygenGeneratorRating(records []string, nBits int) int {
    set := make(map[string]bool)
    for _, record := range records {
        set[record] = true
    }

    i := 0
    for len(set) > 1 {
        commonBits := mostCommonBits(set)
        for _, record := range records {
            if record[i] - '0' != commonBits[nBits - 1 - i] {
                delete(set, record)
            }
        }
        i++
    }

    for record, _ := range set {
        res, err := strconv.ParseUint(record, 2, len(record))
        check(err)
        return int(res)
    }

    panic("unexpected")
}

func co2ScrubberRating(records []string, nBits int) int {
    set := make(map[string]bool)
    for _, record := range records {
        set[record] = true
    }

    i := 0
    for len(set) > 1 {
        commonBits := mostCommonBits(set)
        for _, record := range records {
            if record[i] - '0' == commonBits[nBits - 1 - i] {
                delete(set, record)
            }
        }
        i++
    }

    for record, _ := range set {
        res, err := strconv.ParseUint(record, 2, len(record))
        check(err)
        return int(res)
    }

    panic("unexpected")
}

func check(e error) {
    if (e != nil) {
        panic(e)
    }
}
