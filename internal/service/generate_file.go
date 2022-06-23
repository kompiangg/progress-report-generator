package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kompiangg/report-generator/pkg/errors"
)

func createTitle(params *GenerateFileParams) string {
	year, month, day := time.Now().Date()
	hour, minute, second := time.Now().Clock()
	title := params.RepositoryOwner + "-" + params.RepositoryName + " " + strings.Join([]string{
		strconv.Itoa(year),
		strconv.Itoa(int(month)),
		strconv.Itoa(day),
	}, "-") + " " + strings.Join([]string{
		strconv.Itoa(hour),
		strconv.Itoa(minute),
		strconv.Itoa(second),
	}, ":")

	return title
}

func (s *service) GenerateFile(params *GenerateFileParams) error {
	fmt.Printf("Generating report file\n")

	title := createTitle(params)
	folderName := "report"

	_, err := os.Stat(folderName)

	if os.IsNotExist(err) {
		if err = os.Mkdir(folderName, 0744); err != nil {
			log.Println(errors.ErrCreatingNewFolder.Error())
			return errors.ErrCreatingNewFolder
		}
	}

	file, err := os.Create(folderName + "/" + title + ".md")
	if err != nil {
		log.Println(errors.ErrCreatingNewFile.Error())
		return errors.ErrCreatingNewFile
	}

	defer file.Close()

	file.WriteString("# " + params.RepositoryName + "\n")
	file.WriteString("## " + params.RepositoryOwner + "\n")

	for assignedName, value := range *params.ReportData {
		if _, err := file.WriteString(fmt.Sprintf("* %s\n", assignedName)); err != nil {
			log.Println(err)
			log.Println("ERROR: error while writing", assignedName, "data")
			continue
		}

		for _, issue := range value {
			if _, err := file.WriteString(fmt.Sprintf("  * %s\n", issue)); err != nil {
				log.Println(err)
				log.Println("ERROR: error while writing issue", issue, "owned by", assignedName)
				continue
			}
		}

		file.WriteString("\n")
	}

	if err = file.Sync(); err != nil {
		log.Println(errors.ErrWhileSync.Error())
		return errors.ErrWhileSync
	}

	fmt.Printf("Report file generated\n")
	fmt.Printf("Check the 'report' folder to see your report\n")
	fmt.Printf("Filename: %s\n\n", title)
	return nil
}
