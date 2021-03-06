package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func build() cli.Command {
	cmd := buildcmd{commoncmd: &commoncmd{}}
	flags := append([]cli.Flag{}, cmd.flags()...)
	return cli.Command{
		Name:   "build",
		Usage:  "build function version",
		Flags:  flags,
		Action: cmd.scan,
	}
}

type buildcmd struct {
	*commoncmd
}

func (b *buildcmd) scan(c *cli.Context) error {
	b.commoncmd.scan(b.walker)
	return nil
}

func (b *buildcmd) walker(path string, info os.FileInfo, err error) error {
	walker(path, info, err, b.build)
	return nil
}

// build will take the found valid function and build it
func (b *buildcmd) build(path string) error {
	fmt.Fprintln(b.verbwriter, "building", path)

	ff, err := b.buildfunc(path)
	if err != nil {
		return err
	}

	fmt.Printf("Function %v built successfully.\n", ff.FullName())
	return nil
}
