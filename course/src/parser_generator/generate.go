package parser_generator

//go:generate antlr4 -visitor -Dlanguage=Go -o ../parser -package parser tinyc.g4
