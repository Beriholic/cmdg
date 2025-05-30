package ui

import "github.com/charmbracelet/huh"

func RenderStringsSelect(models []string) (string, error) {
	if len(models) == 0 {
		return "", nil
	}

	var selectedModel string

	huhOptions := make([]huh.Option[string], len(models))
	for i, model := range models {
		huhOptions[i] = huh.NewOption[string](model, model)
	}

	selectField := huh.NewSelect[string]().
		Title("Selecting a Gemini Model.").
		Options(huhOptions...).
		Value(&selectedModel)

	group := huh.NewGroup(selectField)
	form := huh.NewForm(group)

	err := form.Run()
	if err != nil {
		return "", nil
	}
	return selectedModel, nil
}
