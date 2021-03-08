package main
import "fmt"

func put(queue[] int, element int) []int {
  queue = append(queue, element); 
  fmt.Println("Put:", element);
  return queue
}


func main() {
  var queue[] int; 

  queue = put(queue, 1);
  queue = put(queue, 2);
  queue = put(queue, 3);

  fmt.Println("Queue:", queue);

  
}