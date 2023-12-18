package main

import (
    "archive/zip"
    "fmt"
    "io"
    "io/fs"
    "os"
    "os/signal"
    "path/filepath"
    "syscall"
)

func main() {
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
    go func() {
        s := <-sigc
        fmt.Println("Caught signal: " + s.String())

        os.Exit(1)
    }()

    var files []string

    errFileWalk := filepath.Walk("./files", func(path string, info fs.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })

    if errFileWalk != nil {
        os.Exit(2)
    }

    archive, err := os.Create("archive.zip")
    if err != nil {
        os.Exit(3)
    }
    defer archive.Close()
    zipWriting := zip.NewWriter(archive)

    for _, file := range files {

        archivedFile, errOfOpeningFile := os.Open(file)
        if errOfOpeningFile != nil {
            os.Exit(4)
        }

        fileInfo, _ := os.Stat(file)

        if !fileInfo.IsDir() {
            w1, err := zipWriting.Create(file)
            if err != nil {
                os.Exit(5)
            }
            if _, err := io.Copy(w1, archivedFile); err != nil {
                os.Exit(6)
            }
        }

        archivedFile.Close()
	    os.Remove(file)
    }
    zipWriting.Close()
}