package main

import (
        filebeat "github.com/elastic/beats/filebeat/beat"
        "github.com/elastic/beats/libbeat/beat"

        "fmt"
        "git.dhpark.lab/dhpark/tools/ps"
        "os"
        "strconv"
)

var Version = "1.2.3-pdh"
var Name = "filebeat"

// The basic model of execution:
// - prospector: finds files in paths/globs to harvest, starts harvesters
// - harvester: reads a file, sends events to the spooler
// - spooler: buffers events until ready to flush to the publisher
// - publisher: writes to the network, notifies registrar
// - registrar: records positions of files read
// Finally, prospector uses the registrar information, on restart, to
// determine where in each file to restart a harvester.

func main() {

        // create .pid file
        filename := os.Args[0] + ".pid"
        if ps.IsAliveByPidFile(filename) {
                fmt.Println("FileBeat Already Running!")
                return
        }
        pid_file, err := os.Create(filename)

        defer pid_file.Close()
        if err == nil {
                pid := os.Getpid()
                pid_str := strconv.Itoa(pid)
                pid_file.Write([]byte(pid_str))
        }

        beat.Run(Name, Version, filebeat.New())

        if _, err := os.Stat(filename); err == nil {
                os.Remove(filename)
        }

}
