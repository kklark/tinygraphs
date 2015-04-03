package spaceinvaders

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/taironas/tinygraphs/colors"
)

var (
	colorTheme = []color.RGBA{
		{255, 245, 249, 255},
		{232, 70, 134, 255},
		{232, 70, 186, 255},
		{232, 70, 81, 255},
	}
	key            string
	keys           []string
	mapExpectedSVG map[string]string
)

func init() {
	h := md5.New()
	io.WriteString(h, "hello")
	key = fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello2")
	key2 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello3")
	key3 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello4")
	key4 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello5")
	key5 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello6")
	key6 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello7")
	key7 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello8")
	key8 := fmt.Sprintf("%x", h.Sum(nil)[:])
	io.WriteString(h, "hello9")
	key9 := fmt.Sprintf("%x", h.Sum(nil)[:])

	keys = []string{key, key2, key3, key4, key5, key6, key7, key8, key9}

	// svg:
	mapExpectedSVG = map[string]string{"tinygraphs": `<?xml version="1.0"?>
<!-- Generated by SVGo -->
<svg width="220" height="220"
     xmlns="http://www.w3.org/2000/svg" 
     xmlns:xlink="http://www.w3.org/1999/xlink">
<rect x="0" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="100" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="120" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="140" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="0" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="100" y="20" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="140" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="20" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="100" y="40" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="140" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="40" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="60" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="80" y="60" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="60" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="60" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="60" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="160" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="60" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="80" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="80" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="60" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="80" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="160" y="80" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="180" y="80" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="80" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="100" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="60" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="100" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="120" y="100" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="100" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="180" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="100" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="120" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="40" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="60" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="80" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="160" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="180" y="120" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="200" y="120" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="140" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="40" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="140" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="140" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="140" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="140" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="200" y="140" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="100" y="160" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="140" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="160" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="40" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="60" y="180" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="80" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="100" y="180" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="120" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="140" y="180" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="160" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="180" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="200" y="180" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="0" y="200" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="20" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="40" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="60" y="200" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="80" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="100" y="200" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="120" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="140" y="200" width="20" height="20" style="fill:rgb(226,255,222)"/>
<rect x="160" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="180" y="200" width="20" height="20" style="fill:rgb(93,214,75)"/>
<rect x="200" y="200" width="20" height="20" style="fill:rgb(226,255,222)"/>
</svg>
`}

}

func TestSpaceInvaders(t *testing.T) {
	t.Parallel()
	for _, k := range keys {
		rec := httptest.NewRecorder()
		SpaceInvaders(rec, k, colorTheme, 100)
		if rec.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
		}
	}
}

func TestSVGGeneratedBySpaceInvaders(t *testing.T) {
	t.Parallel()
	mapOfColors := colors.MapOfColorThemes()
	for k, expected := range mapExpectedSVG {
		rec := httptest.NewRecorder()

		h := md5.New()
		io.WriteString(h, k)
		md5Key := fmt.Sprintf("%x", h.Sum(nil)[:])

		SpaceInvaders(rec, md5Key, mapOfColors["frogideas"][:2], 220)
		if rec.Code != http.StatusOK {
			t.Errorf("returned %v. Expected %v.", rec.Code, http.StatusOK)
		}
		if rec.Body.String() != expected {
			fExpected := k + "_expected"
			fGot := k + "_got"
			writeToFile(fExpected, expected)
			writeToFile(fGot, rec.Body.String())
			t.Errorf("wrong body response for key: %v\nexpected file: %s\ngot file: %s", k, fExpected, fGot)
		}
	}
}

func writeToFile(name, content string) {
	if f, err := os.Create(name); err != nil {
		log.Fatalln(err)
	} else {
		writer := bufio.NewWriter(f)
		if _, err := writer.Write([]byte(content)); err != nil {
			log.Fatalln(err)
		}
		writer.Flush()
	}
}
