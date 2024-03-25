package main



import (
    "log"
    "handlers/authentication/azure" 
)

func main() {
    credentialFilePath := "~/pfa/creds.json" ;

   
    connInfo := azure.Connect(credentialFilePath);
    
    
    log.Println("Connected to Azure with Subscription ID:", connInfo.subscriptionID);
   
}
