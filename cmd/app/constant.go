package main

var (
	flagName          = "dbtype"
	flagDBDesctiption = "Выберите тип базы данных: imdb или postgrace\n"

	dbValidTypes = map[string]struct{}{
		"imdb":      struct{}{},
		"postgrace": struct{}{},
	}
)
