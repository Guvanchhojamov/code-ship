package main
import "fmt"


func main(){
 fmt.Println(strStr("mississippi", "issip"))
 reverseString([]byte{'h','e','l','l','o'})
}

func strStr(haystack string, needle string) int {
  var l,r int 
  for l<len(haystack) {
	if haystack[l]==needle[r] {
		if r == len(needle)-1 {
			return l-r   
		}
		fmt.Println(l,r)
		r++
	}else if r>0{	  
		l-=r  // left-i yene onki yerine goyyas we yzyny barlap gityas. 
		fmt.Println(l)
		r=0
	} 
	l++
  }
  return -1  
}


func reverseString(s []byte)  {
	l,r:=0,len(s)-1
	for l<r{
	  s[r],s[l]=s[l],s[r]
	  l++
	  r--
	}   
  }