package main

import (
        "os"
        "fmt"
        "bufio"
        "strings"
        "net"
)

var filename = "logs.txt"

var cnt int

func main() {
       args := os.Args[1:]
       if len(args) != 1 {
          fmt.Println("Please enter the IP address that you want to search")
          return
       }

       if val := net.ParseIP(args[0]); val == nil {
          fmt.Println("Enter IP address is not valid IPV4")
          return
       }

    
       lgf, _ := os.Open(filename) 

       scanner := bufio.NewScanner(lgf)
       scanner.Split(bufio.ScanLines)
       var text []string
       for scanner.Scan() {
           text = append(text, scanner.Text())
       }

       lgf.Close()


       for _, each_line := range text {
//           fmt.Printf("%[1]T %[1]s \n", each_line)
//           fmt.Println(strings.Contains(each_line, "10.10.34.22"))
//             fmt.Println(strings.Count(each_line, "10.10.34.22"))
//            cnt = cnt + strings.Count(each_line, "10.10.34.22")
             cnt = cnt + strings.Count(each_line, args[0])
       }

       fmt.Printf("The IP hit to our website %d\n" , cnt)
       
              

}
