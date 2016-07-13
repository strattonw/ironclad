package main


import (
    "fmt"
    "os"
    "path/filepath"
    "github.com/dmulholland/clio/go/clio"
)


// Help text for the 'export' command.
var exportHelptext = fmt.Sprintf(`
Usage: %s export [FLAGS] [OPTIONS] [ARGUMENTS]

  Export a list of entries in JSON format. Entries can be specified by ID or
  by title. If no entries are specified, all entries will be exported.

Arguments:
  [entries]                 List of entries to export by ID or title.

Options:
  -f, --file <str>          Database file. Defaults to the last used file.
  -t, --tag <str>           Filter entries using the specified tag.

Flags:
      --help                Print this command's help text and exit.
`, filepath.Base(os.Args[0]))


// Callback for the 'export' command.
func exportCallback(parser *clio.ArgParser) {

    // Load the database.
    password, _, db := loadDB(parser)

    // Default to exporting all active entries.
    list := db.Active()

    // Do we have query strings to filter on?
    if parser.HasArgs() {
        list = list.FilterByQuery(parser.GetArgs()...)
    }

    // Are we filtering by tag?
    if parser.GetStr("tag") != "" {
        list = list.FilterByTag(parser.GetStr("tag"))
    }

    // Create the JSON dump.
    dump, err := list.Export(db.Key(password))
    if err != nil {
        exit(err)
    }

    // Print the JSON to stdout.
    fmt.Println(dump)
}
