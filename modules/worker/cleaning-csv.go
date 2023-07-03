package worker

import (
	"context"
	"log"
	exportCsv "nashrul-be/crm/modules/export-csv"
	"nashrul-be/crm/repositories"
)

func CleanerCsv(csv exportCsv.UseCaseInterface) func() {
	return func() {
		allCsv, err := csv.GetAll(context.Background(), repositories.ExportCsvQuery{}, 0, 0)
		if err != nil {
			log.Println("error can't get data all data csv req while cleaning")
		}
		for _, csvReq := range allCsv {
			if err := csv.Delete(context.Background(), csvReq.ID); err != nil {
				log.Printf("failed delete request for csv request id %d\n", csvReq.ID)
			}
		}
	}
}
