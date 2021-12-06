# Integrate Back-End Interview Exercise
The center of Integrate's product is the __Demand Acceleration Platform (DAP)__. DAP allows users to interact with marketing lead data. A __marketing lead__ is a person who shows interest in a company's products or services. Our customers need marketing leads that are correct and likely to generate business. 

We are striving to build a world-class application with microservices that expose their functionality through RESTful APIs. For this exercise, you will be asked to create a simple RESTful API. This repo provides a working solution for you to use as a base.

1. The RESTful API should include a Lead data model with the following properties:
    * First Name (mandatory)
    * Last Name (mandatory)
    * Email (mandatory)
    * Company (optional)
    * ZIP Code (optional)
    * Phone Number (mandatory)
    * Date Created

2. The API should allow the user to add a new lead to the dataset.
    * A persistent datastore isn't required. The datastore can instead be memory-based. But, the code should be designed in such a way that a persistent datastore could be substituted later.
3. The API should allow the user to retrieve all people in the dataset.
    
If you're using Golang, you can use the base solution in this repo to get started.

## Dependencies
You must have the following dependencies installed:
```
go mod tidy
```

## Pulling down the code
Use Git to clone this repo. You can run the following from a command line:
`git clone https://github.com/IntegrateDev/Integrate.Interview.Backend.Golang.git`


## Running from the command line
In a command line, run these commands from the root of the repo:
```
cd Integrate.Interview.Backend.Golang
go run main.go
```

## Writing code
This solution has a file called `server/controllers/leadsController.go` that will help you get started. A single GET endpoint has been provided. 

## Testing your code
With your solution running, navigate to http://localhost:1323/swagger/index.html to view the Swagger documentation.

Each endpoint has a "Try It Out" section that allows you to test.


### Swagger Doc
Regenerate Swagger Doc

```
swag init -g "server/routes.go"
```
