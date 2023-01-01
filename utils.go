package main

import "github.com/konveyor/tackle2-hub/api"

// addTags ensure tags created and associated with application.
// Ensure tag exists and associated with the application.
func addTags(application *api.Application, names ...string) error {
	addon.Activity("Adding tags: %v", names)
	appTags := appTags(application)
	// Fetch tags and tag types.
	tpMap, err := tpMap()
	if err != nil {
		return err
	}
	tagMap, err := tagMap()
	if err != nil {
		return err
	}
	// Ensure type exists.
	wanted := api.TagType{
		Name:  "DIRECTORY",
		Color: "#2b9af3",
		Rank:  3,
	}
	tp, found := tpMap[wanted.Name]
	if !found {
		tp = wanted
		if err := addon.TagType.Create(&tp); err != nil {
			return err
		}
		tpMap[tp.Name] = tp
	} else {
		if wanted.Rank != tp.Rank || wanted.Color != tp.Color {
			return &SoftError{Reason: "Tag (TYPE) conflict detected."}
		}
	}
	// Add tags.
	for _, name := range names {
		if _, found := appTags[name]; found {
			continue
		}
		wanted := api.Tag{
			Name:    name,
			TagType: api.Ref{ID: tp.ID},
		}
		tg, found := tagMap[wanted.Name]
		if !found {
			tg = wanted
			if err := addon.Tag.Create(&tg); err != nil {
				return err
			}
			tagMap[wanted.Name] = tg
		} else {
			if wanted.TagType.ID != tg.TagType.ID {
				return &SoftError{Reason: "Tag conflict detected."}
			}
		}
		addon.Activity("[TAG] Associated: %s.", tg.Name)
		application.Tags = append(
			application.Tags,
			api.Ref{ID: tg.ID},
		)
	}
	// Update application.
	return addon.Application.Update(application)
}

// tagMap builds a map of tags by name.
func tagMap() (map[string]api.Tag, error) {
	list, err := addon.Tag.List()
	if err != nil {
		return nil, err
	}
	m := map[string]api.Tag{}
	for _, tag := range list {
		m[tag.Name] = tag
	}
	return m, nil
}

// tpMap builds a map of tag types by name.
func tpMap() (map[string]api.TagType, error) {
	list, err := addon.TagType.List()
	if err != nil {
		return nil, err
	}
	m := map[string]api.TagType{}
	for _, t := range list {
		m[t.Name] = t
	}
	return m, nil
}

// appTags builds map of associated tags.
func appTags(application *api.Application) map[string]uint {
	m := map[string]uint{}
	for _, ref := range application.Tags {
		m[ref.Name] = ref.ID
	}
	return m
}
