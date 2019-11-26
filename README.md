# # Logistics in Blockchain


[![Node.js](https://img.shields.io/badge/Node.js-10.15.3-brightgreen)](https://nodejs.org/)
[![npm](https://img.shields.io/badge/npm-5.6.0-brightgreen)](https://www.npmjs.com/)
![Platforms](https://img.shields.io/badge/platform-linux%20%7C%20osx-brightgreen)


## Prerequisites For Hyperledger Fabric-1.4
- Hyperledger Fabric 1.4 Prerequisites https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html

## Dependencies of  Hyperledger Fabric-1.4
- Pulling Docker images from Dockethub/hyperledgerfabric https://hyperledger-fabric.readthedocs.io/en/release-1.4/install.html

## Description/Summary

A blockchain based solution which records the temperature of a shipment and gives the status of the
shipment to Seller, Buyer. <br />

A blockchain based system which records the temperature,location and other paramaters of a shipment during transport. Depending upon our requirements of tracing of the product , we can keep them into blockchain for tracing the location,status, time,temperature and others.


## Architecture diagram :
![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/Architeure%20diagram.png)

## Network Diagram-Blockchain Network setup details:
![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/NetworkDetails.png)

## Usage
 Pull the docker images 1.4 and place bin directory into network folder. <br />
 ```
1)Run ./remove.sh to remove certificate. Removes all the certificates generated ---> Inside network folder 
2)Run ./stopdocker.sh For removing all the running container and dev images related to chaincode --> inside network folder 
3)Run ./generate.sh Generate new certificats for all the orgization --> Inside network folder
3)Come back to root folder where your ./start.sh is present  
5) Run ./start.sh script to bring your network and server will run on 8080 port.

```


## Postman Results:
 ### RequestLogistic:
 Generally Logistic transfer flow starts from the buyer who need to get some product from seller, will request for the products with id,type,Buyerid,BuyerLocation,Sellerid,SellerLocation,status will. The seller has to accept the request and logistics will come into picture who have to transport the goods from  seller to buyer.
 #### Postman for generating transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/RequestLogistics%20Postman.png)
 
 #### console with transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/RequestLogistic%20consolewith%20txid.png)
 #### Postman for  transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/RequestLogistics_postman_txid.png)
 
  #### Couchb with transaction id
![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/RequestLogistic%20coouchdb%20txid.png)

 # TransitLogistics:
 
 In this phase, logistics will be transported from seller to buyer and the same time blockchain will read the location,status, time,temperature and othersof the products for a fixed time period either from GPS/IOT and updates the details into blockchain 
  #### Postman for  TransitLogistics id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/TransitLogist_postman.png)
 
 #### console with TransitLogistics transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/TransitLogistci_console_txid.png)
 #### Postman for TransitLogistics  transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/TransitLogistic_Postman_txid.png)
 
  #### Couchb with  TransitLogistics transaction id
![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/TransitLogistic_couchdb_Tx.png)
 
 
 
 
 # DeliveryLogistics:
 In this phase,logistics will delivery the product to the buyer.
  #### Postman for  DeliveryLogistics transction  id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/DeliveryLogistic_postman.png)
 
 #### console with DeliveryLogistics transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/DeliveryLogistic_txid_console.png)
 #### Postman for DeliveryLogistics  transaction id
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/DeliveryLogistic_Consolse_txid_postman.png)
 
  #### Couchb with  DeliveryLogistics transaction id
![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/DeliveryLogistic_REJECTEDSTATUS_couchdb.png)
 
 
 # GetAllProducts:
 Get all the Product details within the blockchain
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/GetAllProducts.png)
 
 

 # QueryProduct:
 Get the product details of a particular product.
 ## Console Results
 ![alt text]https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/QueryNAme_console.png)
 ## Postman Results
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/QuerywithName_postman.png)
 
 
 # GetTransactionHistoryForKey:
 Get all the transactions for a particular logistics product id like request logistic,transitlogistic and deliveryof logistics of the product.
 
  ## Console Results
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/GetTxHistory_Console.png)

 ## Postman Results
 ![alt text](https://github.com/adineshreddy1/LogisticsIntoBlockchain/blob/master/screenshots/GetTxHistory_Postman.png)




# Written by
Dinesh Reddy A <br />
https://www.linkedin.com/in/adineshreddy1/ <br />
https://stackoverflow.com/users/9599959/adineshreddy1  <br />
