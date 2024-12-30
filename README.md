# orienteering-results-scraper

Scrape orienteering results web pages to make a levy calculation for British Orienteering.
Produce totals for the number of entrants broken down by whether they are a member of
a British Orienteering affiliated club or not and whether they are a junior or not.

Example output:
```
Fetching web page: https://www.deeside-orienteering-club.org.uk/results/2024/240928BirkenheadPark/index.html
Title: Results - Birkenhead Park SL - 28 Sep 2024
Found 5 courses:
  Light Green      :  20 ;  MS 11, MJ  0 ;  NS  0, NJ  9
  Long Light Green :  40 ;  MS 29, MJ  2 ;  NS  3, NJ  6
  Orange           :  54 ;  MS  1, MJ  0 ;  NS  4, NJ 49
  Yellow           :  30 ;  MS  0, MJ  1 ;  NS  5, NJ 24
  White            :   7 ;  MS  0, MJ  1 ;  NS  0, NJ  6
Total entrants: 151
Members:     Senior  41 , Junior   4
Non-members: Senior  12 , Junior  94

Brian Burden +2, IND on Orange has bad age category: "": counted as senior
Helen Gardner, IND on Yellow has bad age category: "": counted as junior

Non-BO clubs: BIRK, BRID, CDLH, CHESS, FALL, HILBR, HILBRE, IND, STBP, TOWN, UPT
```

The current version handles some of the result web pages created by SiTiming software.

Usage information is available by running the program with no arguments.

Some result pages contain multiple views of the same courses (e.g. urban events).
This results in the same entrant appearing multiple times.
To cope with this scenario, the program has a command line option
to select the courses which are analysed using a regular expression.
The regular expression syntax is that of the Go programming language,
see https://pkg.go.dev/regexp and pages linked from there for details.
Example usage which selects courses "1 - All", "2 - All", etc.:
```
orienteering-results-scraper -only-courses '\d - All'  https://www.deeside-orienteering-club.org.uk/results/2024/241013ChesterUrban/index.html
```

# Running

The program is written in Go. To run it, you need to have Go installed on your system.
You can download Go from https://go.dev/dl/.
Then you can compile the program by running `go build` in the directory containing the source code.
This will create an executable file named `orienteering-results-scraper` in the same directory.
You can then run the program by executing this file from the command line.
Example usage:
```
./orienteering-results-scraper https://www.deeside-orienteering-club.org.uk/results/2024/241117RavenMeols/index.html
```

On Windows the program will be named `orienteering-results-scraper.exe`.
