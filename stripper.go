package main

import (
    "fmt"
    "net/url"
    "regexp"
    "os"
)

func main() {
    utm := regexp.MustCompile("^utm_")

    for i := 1; i < len(os.Args); i++ {
        getURL, _ := url.QueryUnescape(os.Args[i])
        u, _ := url.Parse(getURL)
        m, _ := url.ParseQuery(u.RawQuery)

        for k := range m {
            if utm.MatchString(k) {
              m.Del(k)
            }
        }
        q := m.Encode()

        url := ""
        if u.Scheme != "" { url += u.Scheme + "://" } else { url += "http://" }
        url += u.Host + u.Path
        if q != "" { url += "?" + q }
        if u.Fragment != "" { url += "#" + u.Fragment }
        fmt.Printf(`{"url":"%s"}` + "\n", url)
    }

}
