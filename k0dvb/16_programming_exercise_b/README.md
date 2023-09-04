# Exercise B
[Video](https://www.youtube.com/watch?v=mFB6_sOiggI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=18)

### Problem:
Exercise 4.12 from GOPL: fetching from the web

[Solution](https://github.com/matt4biz/go-class-exer-4.12)

- The popular web comic "xkcd" has a JSON interface.
- A request to https://xkcd.com/571/info.0.json produces a detailed description of comic 571.
- Download each URL once, and build an offline index.
- Write a tool `xkcd` that using this index, prints the URL, date, and title of each comic whose title or transcript matches a list of search terms provided on the command line.

### Data example:
```json
{"month": "4", "num": 571, "link": "", "year": "2009", "news": "", "safe_title": "Can't Sleep", "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}", "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.", "img": "https://imgs.xkcd.com/comics/cant_sleep.png", "title": "Can't Sleep", "day": "20"}
```

### Approach:
Need two programs:
1. go and download data once to disk.
2. search through local data for cartoon.

#### Sample output:
```
$ go run xkcd-load.go xkcd.json
skipping 404: got 404
skipping 2403: got 404
skipping 2404: got 404
read 2401 comics
```
```
$ go run xkcd-find.go xkcd.json some bed sleep
read 2318 comics
https://xkcd.com/571/ 4/20/2009 "Can't Sleep"
found 1 comics
```