package main


import "github.com/replit/database-go"

func main(){
  
  keys, _ := database.ListKeys("log_")


  for _, key := range  keys{
    
    database.Delete(key)
  }
}