package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu/interfaces"
	"github.com/toniliesche/dockertool/modules/application/menu/models"
	"github.com/toniliesche/dockertool/modules/application/menu/pages/base"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker/images"
	"sort"
)

type SelectImage struct {
	base.MenuPage
}

func (p *SelectImage) GetHeadline() string {
	return "Image Management"
}

func (p *SelectImage) Run() (interfaces.PageInterface, int, error) {
	return p.CreateAndRunMenuTask(p.createEntries())
}

func (p *SelectImage) createEntries() (models.EntryList, models.EntryList, error) {
	menuEntries := models.EntryList{}
	specialEntries := models.EntryList{}

	imageList, err := images.FetchImageList()
	if nil != err {
		return nil, nil, err
	}

	mapping := map[string]int{}
	keys := make([]string, 0, len(imageList))
	for key, image := range imageList {
		keys = append(keys, image.Repository+":"+image.Tag)
		mapping[image.Repository+":"+image.Tag] = key
	}

	sort.Strings(keys)

	for _, key := range keys {
		image := imageList[mapping[key]]
		menuEntries = append(
			menuEntries,
			&models.Entry{Label: image.Repository + ":" + image.Tag, Page: &SelectImageAction{ID: image.ID, Repository: image.Repository, Tag: image.Tag}},
		)
	}

	return menuEntries, specialEntries, nil
}
