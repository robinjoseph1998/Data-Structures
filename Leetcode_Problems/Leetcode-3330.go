package main

import "fmt"


func possibleStringCount(word string) int {
    res:=1
    p:=word[0]
    for i:=1;i<len(word);i++{
        if word[i]==p{
            res++
        }else{
            p = word[i]
        }
    }
    return res
}


func main(){

word:="abbcccc"

fmt.Println( possibleStringCount(word))

}
