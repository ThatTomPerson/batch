package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var d = flag.Duration("d", time.Duration(0), "time to add")

func main() {
	flag.Parse()

	args := flag.Args()

	for _, name := range args {
		o, err := os.Open(name)
		if err != nil {
			logrus.Warnf("could not open %s: %v", name, err)
			continue
		}
		defer o.Close()

		s, err := o.Stat()
		if err != nil {
			logrus.Warnf("could not stat %s: %v", name, err)
			continue
		}

		newName := fmt.Sprintf("%s-%s", s.ModTime().Add(*d).Format(time.RFC3339), name)
		f, err := os.Create(newName)
		if err != nil {
			logrus.Warnf("could not create %s: %v", newName, err)
			continue
		}
		defer f.Close()

		io.Copy(o, f)
		logrus.Infof("Copied %s to %s", name, newName)
	}
}
