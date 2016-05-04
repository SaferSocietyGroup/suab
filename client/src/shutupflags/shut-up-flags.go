package shutupflags

import (
	"flag"
	"fmt"
	"strings"
	"strconv"
)

type Flag struct {
	short string
	long string
	defaultValue string
	description string
}
var flags []Flag = make([]Flag, 0)

func AddFlag(short string, long string, defaultValue string, description string) *string {
	short = strings.TrimLeft(short, "-")
	long  = strings.TrimLeft(long , "-")

	var theVar string
	flag.StringVar(&theVar, short, defaultValue, description)
	flag.StringVar(&theVar, long, defaultValue, description)

	flags = append(flags, Flag{short, long, defaultValue, description})
	return &theVar
}


func Usage() string {
        maxLength := strconv.Itoa(calcMaxLength(flags) + 4); // The +4 is for the extra chars in flag := ... below

	usage := "USAGE: suab [flags]\n\n"
	for _, flag := range flags {
		flags := "-" + flag.short + " --" + flag.long
		usage += fmt.Sprintf("%-" +maxLength+ "s %s. Default %s\n", flags, flag.description, flag.defaultValue)
	}
	return usage
}

func calcMaxLength(a []Flag) int {
	max := 0
	for _, flag := range a {
		length := len(flag.short) + len(flag.long)
		if length > max {
			max = length
		}
	}

	return max
}
