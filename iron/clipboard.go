package main


import (
    "os"
    "fmt"
    "time"
    "github.com/atotto/clipboard"
)


// Write a string to the system clipboard. Automatically overwrite the
// clipboard after a ten second delay.
func writeToClipboard(value string) {

    if clipboard.Unsupported {
        exit("clipboard functionality is not supported on this system")
    }

    err := clipboard.WriteAll(value)
    if err != nil {
        exit(err)
    }

    fmt.Fprint(os.Stderr, "Clearing clipboard: ")

    ms_total := 10000
    intervals := 51
    ms_per_interval := ms_total / intervals
    strlen := intervals + 7

    for count := 0; count <= intervals; count++ {

        fmt.Fprintf(os.Stderr, "|")
        fmt.Fprintf(os.Stderr, charstr(count, '-'))
        fmt.Fprintf(os.Stderr, charstr(intervals - count, ' '))
        fmt.Fprintf(os.Stderr, "|")

        ms_remaining := ms_total - count * ms_per_interval
        fmt.Fprintf(os.Stderr, " %2.fs ", float64(ms_remaining)/1000)

        time.Sleep(time.Duration(ms_per_interval) * time.Millisecond)

        fmt.Fprintf(os.Stderr, charstr(strlen, '\b'))
        fmt.Fprintf(os.Stderr, charstr(strlen, ' '))
        fmt.Fprintf(os.Stderr, charstr(strlen, '\b'))
    }

    clipContent, err := clipboard.ReadAll()
    if err != nil {
        exit(err)
    }

    if clipContent == value {
        err = clipboard.WriteAll("[Clipboard overwritten by Ironclad]")
        if err != nil {
            exit(err)
        }
        fmt.Fprintln(os.Stderr, "[CLEARED]")
    } else {
        fmt.Fprintln(os.Stderr, "[ALTERED]")
    }
}