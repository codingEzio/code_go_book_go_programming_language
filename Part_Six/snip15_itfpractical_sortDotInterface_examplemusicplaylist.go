package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track
type byYear []*Track
type byCustomSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {
	sort.Sort(byArtist(tracks)) // convert first, then sorting (in-place)
	printTracks(tracks)

	sort.Sort(byYear(tracks)) // it won't be reversed if you run it again (uh)
	printTracks(tracks)

	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)

	sort.Sort(byCustomSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}

		return false
	}})
	printTracks(tracks)
}

func length(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return duration
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tWriter := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tWriter, format, "\nTitle", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tWriter, format, "-----", "-----", "-----", "-----", "-----")

	for _, track := range tracks {
		fmt.Fprintf(tWriter, format,
			track.Title, track.Artist, track.Album, track.Year, track.Length)
	}

	_ = tWriter.Flush()
}

// For these nine methods, all we do is to satisfy the `sort.Interface`
// (that is, Len, Less, Swap), even the `Reverse` (Built-in) is doing the same thing.

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func (x byCustomSort) Len() int           { return len(x.t) }
func (x byCustomSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x byCustomSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
