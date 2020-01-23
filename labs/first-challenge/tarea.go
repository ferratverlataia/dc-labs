package main
import (
  "bufio"
  "fmt"
  "os"
)

type void struct{};
var member void;

func main(){
  reader := bufio.NewReader(os.Stdin)

  inputstring, _ := reader.ReadString('\n');
  var maxsize=0;
  var currsize=0;
  set:= make(map[string]void);

  for i:=0;i<len(inputstring);i++ {

  set[string(inputstring[i])]=member;

  if (currsize==len(set)){
    if(maxsize<currsize){
      maxsize=currsize;
    }
    currsize=0;
    set= make(map[string]void);
  }else{
    currsize+=1;
    if(maxsize<len(set))&&(i==len(inputstring)-1)&&(currsize!=len(set)){
      maxsize=len(set);
    }
  }
 
  }
  fmt.Printf("%d",maxsize);   
}
