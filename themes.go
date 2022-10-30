package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

//go:embed themes.json
var themesBts []byte

var (
	themesOnce sync.Once
	themesMap  = map[string]Theme{}
)

// Themes return the list of themes.
var Themes = func() map[string]Theme {
	themesOnce.Do(func() {
		var themes []windowsTerminalTheme
		if err := json.Unmarshal(themesBts, &themes); err != nil {
			fmt.Fprintf(os.Stderr, "could not load themes.json: %s\n", err)
			os.Exit(1)
		}
		for _, theme := range themes {
			themesMap[theme.Name] = Theme{
				Background:    theme.Background,
				Foreground:    theme.Foreground,
				Selection:     theme.Selection,
				Cursor:        theme.Cursor,
				CursorAccent:  theme.CursorAccent,
				Black:         theme.Black,
				BrightBlack:   theme.BrightBlack,
				Red:           theme.Red,
				BrightRed:     theme.BrightRed,
				Green:         theme.Green,
				BrightGreen:   theme.BrightGreen,
				Yellow:        theme.Yellow,
				BrightYellow:  theme.BrightYellow,
				Blue:          theme.Blue,
				BrightBlue:    theme.BrightBlue,
				Magenta:       theme.Purple,
				BrightMagenta: theme.BrightPurple,
				Cyan:          theme.Cyan,
				BrightCyan:    theme.BrightCyan,
				White:         theme.White,
				BrightWhite:   theme.BrightWhite,
			}
		}
	})
	return themesMap
}

type windowsTerminalTheme struct {
	Name         string `json:"name"`
	Background   string `json:"background"`
	Foreground   string `json:"foreground"`
	Selection    string `json:"selectionBackground"`
	Cursor       string `json:"cursorColor"`
	CursorAccent string `json:"cursorAccent"`
	Black        string `json:"black"`
	BrightBlack  string `json:"brightBlack"`
	Red          string `json:"red"`
	BrightRed    string `json:"brightRed"`
	Green        string `json:"green"`
	BrightGreen  string `json:"brightGreen"`
	Yellow       string `json:"yellow"`
	BrightYellow string `json:"brightYellow"`
	Blue         string `json:"blue"`
	BrightBlue   string `json:"brightBlue"`
	Purple       string `json:"purple"`
	BrightPurple string `json:"brightPurple"`
	Cyan         string `json:"cyan"`
	BrightCyan   string `json:"brightCyan"`
	White        string `json:"white"`
	BrightWhite  string `json:"brightWhite"`
}