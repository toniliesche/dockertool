package containers

import (
	"fmt"
	"strings"
)

type FilterOptions struct {
	RepoFilter  string
	NameFilter  string
	StateFilter bool
	RunningOnly bool
	StoppedOnly bool
}

type filterInterface interface {
	runFilter([]*Container) []*Container
}

type repoFilter struct {
	filter string
}

func (r *repoFilter) runFilter(containers []*Container) []*Container {
	filteredContainers := make([]*Container, 0)

	for _, container := range containers {
		if strings.HasPrefix(container.Image, r.filter) {
			filteredContainers = append(filteredContainers, container)
		}
	}

	return filteredContainers
}

type nameFilter struct {
	filter string
}

func (n *nameFilter) runFilter(containers []*Container) []*Container {
	filteredContainers := make([]*Container, 0)

	for _, container := range containers {
		if strings.Contains(container.Name, n.filter) {
			filteredContainers = append(filteredContainers, container)
		}
	}

	return filteredContainers
}

type stateFilter struct {
	running bool
}

func (s stateFilter) runFilter(containers []*Container) []*Container {
	filteredContainers := make([]*Container, 0)

	for _, container := range containers {
		if s.running == container.IsRunning() {
			filteredContainers = append(filteredContainers, container)
		}
	}

	return filteredContainers
}

func runFilters(options *FilterOptions, containers []*Container) ([]*Container, error) {
	filters, err := getFilterFunctions(options)
	if err != nil {
		return nil, err
	}

	for _, filter := range filters {
		containers = filter.runFilter(containers)
	}

	return containers, nil
}

func getFilterFunctions(options *FilterOptions) ([]filterInterface, error) {
	filters := make([]filterInterface, 0)

	if "" != options.RepoFilter {
		filters = append(filters, &repoFilter{filter: options.RepoFilter})
	}

	if "" != options.NameFilter {
		filters = append(filters, &nameFilter{filter: options.NameFilter})
	}

	if true == options.StateFilter {
		if (options.RunningOnly && options.StoppedOnly) || (!options.RunningOnly && !options.StoppedOnly) {
			return nil, fmt.Errorf("exactly one option out of [RunningOnly, StoppedOnly] MUST be true for StateFilter to work")
		}

		filters = append(filters, &stateFilter{running: options.RunningOnly})
	}

	return filters, nil
}
