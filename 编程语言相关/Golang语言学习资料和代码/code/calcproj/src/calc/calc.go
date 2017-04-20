package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func(){
	fmt.Println("usage:calc command [arguments]...")
	fmt.Println("\nThe commands are:\n\tAddition of two values.\n\tsqrt\tSquare root of a non-negative value")
}

func main(){
	args := os.Args[1:]
	if args == nil || len(args)<2{
		Usage()
		return
	}
	switch args[0] {
	      case "add":
		if len(args)!=3{
			fmt.Println("Usage:calc add <integer1><integer2>")
			return
		}
		v1,err1 := strconv.Atoi(args[1])
		v2,err2 := strconv.Atoi(args[2])
		if(err1 != nil || err2 !=nil){
			fmt.Println("Usage:calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1,v2)
		fmt.Println("result:",ret)
	      case "sqrt":
		if len(args) !=2{
			fmt.Println("usage:calc sqrt <integer>")
			return
		}
		v,err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("usage:calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("result:",ret)
	      default:
		Usage()
	}
}


