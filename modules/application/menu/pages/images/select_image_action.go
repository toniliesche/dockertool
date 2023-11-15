package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/images"
)

type SelectImageAction struct {
	base.Page
	base.MenuPage
	ID         string
	Repository string
	Tag        string
}

func (p *SelectImageAction) GetHeadline() string {
	return p.Repository + ":" + p.Tag
}

func (p *SelectImageAction) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectImageAction) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	_, err := images.FetchImage(p.ID)
	if nil != err {
		return nil, nil, err
	}

	menuEntries = append(menuEntries, &models.Entry{Label: "Show Info", Page: &InspectImage{Image: p.ID}})

	return menuEntries, specialEntries, nil
}
