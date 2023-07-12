package worker

import (
	"log"
	"nashrul-be/crm/utils/filesystem"
)

func CleanerCsv(folder filesystem.Folder) func() {
	return func() {
		log.Println("Starting cleaning csv file...")
		files := folder.GetAllFiles()
		for _, file := range files {
			log.Printf("Deleting file %s\n", file.Filename())
			if err := file.Remove(); err != nil {
				log.Printf("Failed to delete file %s\n", file.Filename())
				continue
			}
			log.Printf("Success delete file %s\n", file.Filename())
		}
	}
}
