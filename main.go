package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	wd, _ := os.Getwd()
	dir := wd


	for !FolderContainsFonts(dir) {

		if dir == wd {
			fmt.Println("Sorry, the current working directory does not contain any fonts.")
		} else {
			fmt.Println("Sorry, the directory selected does not contain any fonts.")
		}
 		

		dir = ""

		fmt.Print("Enter the path of where your fonts are located: ")
		fmt.Scanln(&dir)
	}
	
	fmt.Println("Fonts have been found.")
	fmt.Printf("dir: %s", dir)
	fmt.Println()
	fmt.Println()

	fmt.Print("Rename fonts for Android Resource? y/n: ")
	var option string

	fmt.Scanln(&option)

	if len(option) == 0 {
		os.Exit(-1)
	}

	if option[0] != 'y' && option[0] != 'Y' {
		os.Exit(-1)
	}

	RenameFontsForAndroid(dir)

	fmt.Println("Renamed fonts for Android resource.")
}


func FolderContainsFonts(dir string) bool {

	common_font_extensions := []string{ ".otf", ".ttf", ".woff", ".woff2" }

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		extension := path.Ext(filename)

		for _, x := range common_font_extensions {
			if x == extension {
				return true
			}
		}
	}

	return false

}

func RenameFontsForAndroid(dir string) {

	common_font_extensions := []string{ ".otf", ".ttf", ".woff", ".woff2" }

	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		_path := entry.Name()

		filename := path.Base(_path)
		extension := path.Ext(_path)

		old_path := dir + "\\" + filename

		for _, x := range common_font_extensions {
			if x == extension {
				filename = strings.ToLower(filename)

				filename = strings.ReplaceAll(filename, "-", "_")
				filename = strings.ReplaceAll(filename, " ", "_")

				new_path := dir + "\\" + filename

				if err := os.Rename(old_path, new_path); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
