package methods

import (
	"fmt"
	"io/fs"
	"log"
	"main/configs"
	"os"
	"path/filepath"
	"strings"
)

func ListFilesAndDirs() []string {
	dirList := []string{}
	currentDir := configs.CurrentDirPath

	err := filepath.WalkDir(currentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative_path, err := filepath.Rel(currentDir, path)
		if err != nil {
			return err
		}
		fmt.Println(relative_path)
		dirList = append(dirList, relative_path)
		return nil
	})
	if err != nil {
		log.Println("Error while reading the dir : ", err.Error())
		return []string{}
	}
	return dirList
}

func Ignore_Files_n_Dirs() {
	ignore_file_path := ".vcr.ignore"

	if _, err := os.Stat(ignore_file_path); os.IsNotExist(err) {
		fmt.Println("ignore file doesnot exits")
		return
	}
	content, err := os.ReadFile(".vcr.ignore")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ignore_list := strings.Split(string(content), "\n")
	for _, file := range ignore_list {
		fmt.Println(file)
	}
}
