package admin

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	clr "github.com/logrusorgru/aurora"
	"github.com/pkg/errors"
	"github.com/qor/admin"
	"github.com/webmalc/it-stats-rankings-scrapper/admin/bindatafs"
	"github.com/webmalc/it-stats-rankings-scrapper/models"
)

// Admin is the admin structure
type Admin struct {
	config *Config
	admin  *admin.Admin
	db     *gorm.DB
}

// Init initializes the admin
func (a *Admin) Init() {
	a.admin = admin.New(&admin.AdminConfig{DB: a.db})
	entry := models.Language{}
	lang := a.admin.AddResource(&entry)
	a.admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))

	lang.IndexAttrs(
		"ID",
		"Source",
		"Title",
		"Synonym",
		"Position",
		"CreatedAt",
		"UpdatedAt",
	)
	lang.NewAttrs(
		"Source",
		"Title",
		"Synonym",
		"Position",
	)
	lang.EditAttrs(
		"Source",
		"Title",
		"Synonym",
		"Position",
	)
	lang.Meta(&admin.Meta{
		Name: "Source",
		Config: &admin.SelectOneConfig{
			Collection: entry.GetAvailableSources(),
		},
	})
	lang.Filter(&admin.Filter{Name: "Source"})
	lang.Filter(&admin.Filter{Name: "CreatedAt"})

	a.admin.GetRouter().Get("/", func(c *admin.Context) {
		http.Redirect(
			c.Writer, c.Request, "/languages", http.StatusSeeOther,
		)
	})
}

// startupMessage prints the startup message
func (a *Admin) getStartupMessage() string {
	url := a.config.AdminURL
	if url[0] == ':' {
		url = "localhost" + url
	}
	address := "http"
	if a.config.AdminSSL {
		address = "https"
	}
	return clr.Sprintf(clr.Green("\nRunning the admin %s://%s%s\n"),
		clr.Red(address),
		clr.Red(url),
		clr.Red(a.config.AdminPath),
	)
}

// Run runs the admin
func (a *Admin) Run(args []string) {
	fmt.Println(a.getStartupMessage())
	mux := http.NewServeMux()
	a.admin.MountTo(a.config.AdminPath, mux)
	srv := &http.Server{
		Addr:    a.config.AdminURL,
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		panic(errors.Wrap(err, "admin"))
	}
}

// NewAdmin returns a new admin object
func NewAdmin(db *gorm.DB) *Admin {
	config := NewConfig()
	a := Admin{config: config, db: db}
	a.Init()
	return &a
}
