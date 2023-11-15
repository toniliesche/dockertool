package images

import (
	"github.com/toniliesche/dockertool/modules/application/menu"
	"github.com/toniliesche/dockertool/modules/infrastructure/docker"
	"sort"
)

type Index struct {
	menu.Base
	menu.Menu
}

func (p *Index) GetHeadline() string {
	return "Image Management"
}

func (p *Index) Run() (menu.PageInterface, int, error) {
	images, err := docker.FetchImages()
	if err != nil {
		return p.HandleError(err, true)
	}

	menuEntries := menu.EntryList{}

	mapping := map[string]int{}
	keys := make([]string, 0, len(images))
	for key, image := range images {
		keys = append(keys, image.Repository+":"+image.Tag)
		mapping[image.Repository+":"+image.Tag] = key
	}

	sort.Strings(keys)

	for _, key := range keys {
		image := images[mapping[key]]
		menuEntries = append(
			menuEntries,
			&menu.Entry{Label: image.Repository + ":" + image.Tag, Page: &SelectAction{ID: image.ID, Repository: image.Repository, Tag: image.Tag}},
		)
	}

	result := p.RunMenu(menuEntries, menu.EntryList{})
	return p.EvaluateResult(result)
}
