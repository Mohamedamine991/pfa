package main



import (
    "log"
    "github.com/Mohamedamine991/pfa/handlers/authentication"
)

func main() {
    credentialFilePath := "~/pfa/creds.json" ;

   
    connInfo := azure.Connect(credentialFilePath);
    
    
    log.Println("Connected to Azure with Subscription ID:", connInfo.subscriptionID);
   
}
