package main 
import ("fmt")

var (
	n int = 027
)

func main (){ 
fmt.Printf("%b\n",n)
fmt.Printf("%d\n",n) 
fmt.Printf("%+d\n",n) 
fmt.Printf("%o\n",n) 
fmt.Printf("%x\n",n) 
fmt.Printf("%X\n",n)
fmt.Printf("%#x\n",n)
fmt.Printf("%4d\n",n)
fmt.Printf("%-4d\n",n)
fmt.Printf("%04d\n",n)

}
