// Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
//      * Cloud-Barista: https://github.com/cloud-barista
//
// AZURE COMPUTE Hander (AZURE-SDK-FOR-GO:COMPUTE Version 28.0.0, Thanks AZURE.)
//
// by powerkim@powerkim.co.kr, 2019.04.
package azurehandler

import (
        "context"
        "fmt"
        "io/ioutil"
        "log"
        "os"
        "time"
        "encoding/json"
        "strconv"

        //"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
        //"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"
        "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"

        "github.com/Azure/go-autorest/autorest"
        "github.com/Azure/go-autorest/autorest/to"
        "github.com/Azure/go-autorest/autorest/azure/auth"
        "github.com/Azure/go-autorest/autorest/azure"
        
)
// Information for connection
type ConnectionInfo struct {
	context context.Context
    subscriptionID string
	authorizer autorest.Authorizer
}

type ImageInfo struct {
	Publisher string
	Offer     string
	Sku       string
	Version   string
}

type VMInfo struct {
	UserName string
	Password string
	SshPublicKeyPath string
}


func Connect(credentialFilePath string) ConnectionInfo {
	var connInfo ConnectionInfo

	//ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	//defer cancel()
	ctx, _ := context.WithTimeout(context.Background(), 6000*time.Second)
	connInfo.context = ctx

        // get subscritionID from auth file.
        authInfo, authErr := readJSON(credentialFilePath)
        if authErr != nil {
                log.Fatal(authErr)
        }
        connInfo.subscriptionID = (*authInfo)["subscriptionId"].(string)

        // get autorest.Authorizer Object.
        var err error
        connInfo.authorizer, err = auth.NewAuthorizerFromFile(azure.PublicCloud.ResourceManagerEndpoint)
        if err != nil {
                log.Fatal(err)
        }

	return connInfo
}

func readJSON(path string) (*map[string]interface{}, error) {
        data, err := ioutil.ReadFile(path)
        if err != nil {
                log.Fatalf("failed to read file: %v", err)
        }
        contents := make(map[string]interface{})
        json.Unmarshal(data, &contents)
        return &contents, nil
}

