package main


var Types []string=[]string{"p","e","r","f"} 

type TestServerResponse struct{
numType string	
windowPrevState  []int
windowCurrState []int
}



func (s *TestServerResponse)calculateAvg()int{
sum:=0;
currWindowSize:=len(s.windowCurrState)
prevWindowSize:=len(s.windowPrevState)
count:= currWindowSize +prevWindowSize
 for i:=0;i<currWindowSize;i++{
	sum+=s.windowCurrState[i]
 }
  for i:=0;i<prevWindowSize;i++{
	sum+=s.windowPrevState[i]
 }
 return sum/count
}



var PrimeResponse *TestServerResponse=&TestServerResponse{"p",make([]int,0),make([]int,0)}

 