package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/buger/goterm"
	"github.com/latiif/ara/goarabic"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ttywidth := int(goterm.Width())

	for scanner.Scan() {

		fmt.Println(
			goarabic.MakeRTL(ttywidth,
				goarabic.ReversePreservingNonArabic(
					goarabic.ToGlyph(
						goarabic.Smooth(
							goarabic.RemoveTashkeel(scanner.Text()))))))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
