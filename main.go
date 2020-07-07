//+build !test

package main

import (
	"github.com/webmalc/it-stats-rankings-scrapper/cmd"
	"github.com/webmalc/it-stats-rankings-scrapper/common/config"
	"github.com/webmalc/it-stats-rankings-scrapper/common/db"
	"github.com/webmalc/it-stats-rankings-scrapper/common/logger"
	"github.com/webmalc/it-stats-rankings-scrapper/models"
	"github.com/webmalc/it-stats-rankings-scrapper/scrappers"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	models.Migrate(conn)
	cmdRouter := cmd.NewCommandRouter(log, scrappers.NewRunner())
	cmdRouter.Run()
}
