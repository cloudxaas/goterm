package goterm

import (
        "fmt"
        "strings"
)


//if yn == 1, Yes as default otherwise No as default
func YesNo(yn uint8) uint8 {
        if yn == 0 {
                fmt.Printf("(y/N): ")
        }else{  
                fmt.Printf("(Y/n): ")
        }

        var s string

        for {   
                _, err := fmt.Scan(&s)
                if err == nil {
                        break
                        //fmt.Print(err)
                }
        }
        if strings.ToLower(strings.TrimSpace(s)) == "y" {
                return 1
        }
        return 0
}
