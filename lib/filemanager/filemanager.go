package filemanager

import (
	"os"
	"path/filepath"
	"text/template"
	"tikoflano/aoc/lib/utils"
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetSessionCookie() string {
	dat, err := os.ReadFile(".session")
	utils.CheckError(err, "Failed to read .session file")

	return string(dat)
}
func CreateFile(path string) *os.File {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0750)

	file, err := os.Create(path)
	utils.CheckError(err, "Failed to create file")

	return file
}

func CreateFileFromTemplate(path string, tplFile string, tplData interface{}) {
	file := CreateFile(path)

	defer func() {
		err := file.Close()
		utils.CheckError(err, "Failed to close file")
	}()

	tpl, err := os.ReadFile(tplFile)
	utils.CheckError(err, "Failed to read template file")

	tmpl, err := template.New("tpl").Parse(string(tpl[:]))
	utils.CheckError(err, "Failed to parse template")

	err = tmpl.Execute(file, tplData)
	utils.CheckError(err, "Failed to write file")
}

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	utils.CheckError(err, "Error reading file")

	return string(data)
}
