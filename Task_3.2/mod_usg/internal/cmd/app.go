package cmd

import "github.com/Bazhenator/StringFuncs/pkg/strSearch"

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Search(filePath string, toFindSubStr string) (bool, error) {
	return strSearch.Contains(filePath, toFindSubStr)
}
