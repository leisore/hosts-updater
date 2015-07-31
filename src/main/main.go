package main

import (
	"fmt"
	"hostsupdater"
	"io"
	"os"
)

func main() {

	walkers := hostsupdater.GetWalkers()
	for i, w := range walkers {
		fmt.Printf("run walker %d:\n", i+1)
		fmt.Printf("  %-8s: %s\n", "name", w.Name())
		fmt.Printf("  %-8s: %s\n", "version", w.Version())
		fmt.Printf("  %-8s: %s\n", "desc", w.Desc())
		//fmt.Printf("  %-8s: %v\n", "author", w.Authors())
		fmt.Printf("  running......\n")

		reader, err := w.WalkedHosts()
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			if err2 := UpdateHosts(reader); err2 != nil {
				fmt.Printf("Error: %s\n", err.Error())
				os.Exit(-1)
			} else {
			}
		}
	}
	fmt.Printf("\nupdate hosts successfully!\nenter any key to exit>")
	bytes := []byte{0}
	os.Stdin.Read(bytes)
}

func UpdateHosts(r io.Reader) error {
	hosts := hostsupdater.GetHostsFilePath()
	fmt.Printf("  update hosts: %s\n", hosts)
	w, err := os.OpenFile(hosts, os.O_TRUNC|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	} else {
		defer w.Close()
		_, err := io.Copy(w, r)
		return err
	}
}
