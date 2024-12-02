package aoc

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"tikoflano/aoc/lib/constants"
	"tikoflano/aoc/lib/filemanager"
	"tikoflano/aoc/lib/utils"

	"github.com/PuerkitoBio/goquery"
)

type AdventOfCode struct {
	Year   int
	Days   [25]Day
	Client *http.Client
}

func NewAdventOfCode(year int) AdventOfCode {
	aoc := AdventOfCode{
		Year: year,
		Days: [25]Day{},
	}

	aoc.SetHttpClient()
	aoc.Init()

	return aoc
}

func (aoc *AdventOfCode) GetURL() *url.URL {
	urlString := fmt.Sprintf("%s/%d", constants.BaseURL, aoc.Year)
	parsedURL, err := url.Parse(urlString)
	utils.CheckError(err, "Failed to parse URL")

	return parsedURL
}

func (aoc *AdventOfCode) SetHttpClient() *http.Client {
	jar, err := cookiejar.New(nil)
	utils.CheckError(err, "Failed to create cookie jar")

	cookie := &http.Cookie{
		Name:  "session",
		Value: filemanager.GetSessionCookie(),
		Path:  "/",
	}

	jar.SetCookies(aoc.GetURL(), []*http.Cookie{cookie})

	aoc.Client = &http.Client{
		Jar: jar,
	}

	return &http.Client{
		Jar: jar,
	}
}

func (aoc *AdventOfCode) Init() {
	res, err := aoc.Client.Get(aoc.GetURL().String())
	utils.CheckError(err, fmt.Sprintf("Failed to GET input from %s", aoc.GetURL().String()))

	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	utils.CheckError(err, "Failed to create HTML document")

	// Find the review items
	doc.Find(".calendar a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("span.calendar-day").Text()
		dayNumber, err := strconv.Atoi(strings.Trim(title, " "))
		utils.CheckError(err, "Failed to parse day title")

		day := Day{Aoc: aoc, Number: dayNumber, Url: fmt.Sprintf("%s/%d/day/%d", constants.BaseURL, constants.Year, dayNumber)}

		problem1 := Problem{Number: 1, Day: &day}
		problem2 := Problem{Number: 2, Day: &day}

		day.Problems[0] = problem1
		day.Problems[1] = problem2

		if s.HasClass("calendar-complete") {
			day.Problems[0].Solved = true
		} else if s.HasClass("calendar-verycomplete") {
			day.Problems[0].Solved = true
			day.Problems[1].Solved = true
		}

		aoc.Days[dayNumber-1] = day
	})
}

func (aoc AdventOfCode) String() string {
	s := fmt.Sprintf("Advent of Code Year %d\n------------------------\n", aoc.Year)
	for _, day := range aoc.Days {
		s += fmt.Sprintf("%v\n", day)
	}

	return s
}

func (aoc *AdventOfCode) NextDay() *Day {
	for _, day := range aoc.Days {
		if !day.IsSolved() {
			return &day
		}
	}

	return nil
}
