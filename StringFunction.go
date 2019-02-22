package main

import s "strings"
import "fmt"
var p = fmt.Println
func main(){
	p("Contains:\t\t",s.Contains("test", "es"))
	p("Count\t\t",s.Count("test","t"))
	p("HasPrefix\t\t",s.HasPrefix("test", "te"))
	p("HasSuffix\t\t",s.HasSuffix("test", "st"))
	p("Index\t\t", s.Index("test","e"))
	p("Join\t\t",s.Join([]string{"a","b"}, "====="))
	p("Repeat\t\t",s.Repeat("a",5))
	p("Replace\t\t", s.Replace("foo","o","0",-1))
	p("Replace\t\t", s.Replace("foo","o","0",1))
	p("Split\t\t",s.Split("a-a-a-a","-"))
	p("ToLower\t\t",s.ToLower("TEST"))
	p("ToUpper\t\t",s.ToUpper("test"))
	p()
	
}