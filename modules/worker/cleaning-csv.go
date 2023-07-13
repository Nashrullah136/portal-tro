package worker

import (
	"nashrul-be/crm/utils/filesystem"
	"nashrul-be/crm/utils/logutils"
)

func CleanerCsv(folder filesystem.Folder) func() {
	return func() {
		logutils.Get().Println("Starting cleaning csv file...")
		files := folder.GetAllFiles()
		for _, file := range files {
			logutils.Get().Printf("Deleting file %s\n", file.Filename())
			if err := file.Remove(); err != nil {
				logutils.Get().Printf("Failed to delete file %s\n", file.Filename())
				continue
			}
			logutils.Get().Printf("Success delete file %s\n", file.Filename())
		}
	}
}
