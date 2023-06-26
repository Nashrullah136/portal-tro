package worker

import (
	"context"
	"github.com/go-co-op/gocron"
	"log"
	exportCsv "nashrul-be/crm/modules/export-csv"
	"nashrul-be/crm/repositories"
	"time"
)

func CleanerCsv(csv exportCsv.UseCaseInterface) {
	wib, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(wib)
	s.Every(1).Day().At("00:00").Do(func() {
		allCsv, err := csv.GetAll(context.Background(), repositories.ExportCsvQuery{}, 0, 0)
		if err != nil {
			log.Println("error can't get data all data csv req while cleaning")
		}
		for _, csvReq := range allCsv {
			if err := csv.Delete(context.Background(), csvReq.ID); err != nil {
				log.Printf("failed delete request for csv request id %d\n", csvReq.ID)
			}
		}
	})
	s.StartAsync()
}
