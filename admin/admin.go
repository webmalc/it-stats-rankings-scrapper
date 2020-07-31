package admin

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/qor/admin"
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

// Run runs the admin
func (a *Admin) Run(args []string) {
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
