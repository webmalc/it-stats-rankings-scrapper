//+build !test

package main

import (
	"github.com/webmalc/it-stats-rankings-scrapper/admin"
	"github.com/webmalc/it-stats-rankings-scrapper/admin/bindatafs"
	"github.com/webmalc/it-stats-rankings-scrapper/cmd"
	"github.com/webmalc/it-stats-rankings-scrapper/common/config"
	"github.com/webmalc/it-stats-rankings-scrapper/common/db"
	"github.com/webmalc/it-stats-rankings-scrapper/common/logger"
	"github.com/webmalc/it-stats-rankings-scrapper/models"
	"github.com/webmalc/it-stats-rankings-scrapper/scrappers"
	"github.com/webmalc/it-stats-rankings-scrapper/services"
)

func main() {
	config.Setup()
	log := logger.NewLogger()
	conn := db.NewConnection()
	defer conn.Close()
	models.Migrate(conn)
	repo := models.NewLanguageRepository(conn, services.NewNameNormalizer())
	cmdRouter := cmd.NewCommandRouter(
		log,
		admin.NewAdmin(conn.DB),
		bindatafs.NewGenerator(),
		scrappers.NewRunner(repo, log),
	)
	cmdRouter.Run()
}
