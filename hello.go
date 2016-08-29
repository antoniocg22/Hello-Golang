/* The code below is just to make a program to print "Hello, world" but 
this script has the characteristic that the .exe will keep the comand window 
opens by two diferent methods:
1. The first one will keep open the window open for ten seconds, then it 
will automaticly close. 
2. The other method will keep the window open until the user press 'enter'.

The code will be configured as the second method . If you want to close that 
the comand window closes automaticly do the next steps: 
1- In the 'import' script, uncomment "time"  
2- Also too, in the 'import', comment "bufio" n "os" 
3- Then, in the 'func main()' uncomment the variable 'sec' that is locate in 
the line 26. Also, uncomment the lines 34 and 35 wich have the code that make 
open the console.
4- Last, comment the lines 36 n 37. This two lines are code for the second method 

One last thing: For the first method, which closes automaticly the window, if 
you don't want it to close the window in shorter or longer time you just need 
to make the next adjusment: 
1- In the line  change the variable sec to the seconds you want the console 
opens  */ 

package main

import ("fmt" 
		//"time" //This is necesary for the first method (close automaticly)
		"bufio"  // "bufio" n "os" are for the second method (press 'enter')
		"os"
		)

func main(){
	//sec := 10 //This is the variable to set the seconds you want the console opens
	fmt.Printf("Hello, world.\n")
	/*duration := time.Duration(sec)*time.Second //This and the next line are first method
	time.Sleep(duration)*/
	fmt.Printf("Press 'Enter' to continue...") //This and the next line are for the second method
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
