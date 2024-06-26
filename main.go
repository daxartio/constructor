package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	packagePath := flag.String("p", ".", "Package full path.")
	structsStr := flag.String("s", "", "Match structs name. (comma separated)")
	write := flag.Bool("w", false, "Write generated code.")
	format := flag.String("f", "%s_constructor_gen.go", "A format of the filename.")
	noPrefix := flag.Bool("n", false, "No prefix.")
	help := flag.Bool("h", false, "Show help.")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	structsList := []string{}
	if *structsStr != "" {
		structsList = splitCommaSeparated(*structsStr)
	}

	structs, err := getPackageStructs(*packagePath)
	if err != nil {
		return
	}

	structsMatch := ``
	for _, value := range structsList {
		if value != `` {
			structsMatch += `|` + value + `|`
		}
	}

	g := &Generator{*noPrefix}
	for _, strct := range structs {
		if structsMatch != `` {
			name := strct.Name
			if !strings.Contains(structsMatch, `|`+name+`|`) {
				continue
			}
		}

		src, err := g.generateCode(strct)
		if err != nil {
			return
		}

		if *write {
			filename := fmt.Sprintf(*format, LowerFirst(CamelCase(strct.Name)))
			fullpath := path.Join(*packagePath, filename)

			err := os.WriteFile(fullpath, []byte(src), 0644)
			if err != nil {
				continue
			}
		} else {
			fmt.Print(src)
		}
	}
}

func splitCommaSeparated(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
