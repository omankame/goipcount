//  https://gist.github.com/yehgdotnet/9779505f52d9138aae39dc5ae4de224b
// 1st defined the reg exp pattern, then open the file, scan the file , split into lines so ever lines is one string, then defined text text var
// then get the submatchshell with help of FindAllString function and loop it to extract IPS, now second stage is append all these IP into one 
// slice and make one fucntion call to count the occurence of IPS
// https://stackoverflow.com/questions/65900501/go-determine-number-of-word-occurences-on-a-string-slice
package main

import (
        "os"
        "fmt"
        "bufio"
//        "strings"
        "sort"
        "regexp"
)

var filename = "access.logs"

var cnt int


func main() {
       
            lgf, _ := os.Open(filename) 
            re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
//            re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

//            fmt.Printf("Pattern: %v\n", re.String()) // print pattern
            scanner := bufio.NewScanner(lgf)
            scanner.Split(bufio.ScanLines) 
            var text string
            var mapstr []string
            for scanner.Scan() {
               text = scanner.Text()         
//             fmt.Println(re.MatchString(text), text) // true

               submatchall := re.FindAllString(text, -1)         
               for _, element := range submatchall {
//                    fmt.Printf("%[1]T, %[1]s\n", element) 
                    mapstr = append(mapstr, element)
               }

             }
       lgf.Close()
//       fmt.Println(mapstr)  //printing slice with all IP values
 
       countOccurrences(mapstr)

}


func countOccurrences( mapstr []string) {
//     fmt.Println(mapstr)

     dict := make(map[string]int)

    for _, v := range mapstr {
         fmt.Println(v)
         dict[v]++ 
    }

    fmt.Println(dict)

    keys := make([]string, 0, len(dict))
    for k := range dict {
        keys = append(keys, k)
     }

     sort.Strings(keys)

     for _, k := range keys {
         fmt.Println(k, dict[k] )
     }

}
