# Doitpay Integration Samples

## Getting Started
Welcome to the [Doitpay Integration Samples] repository! This readme provides an overview of the project and how to get started. This is a very simple and straightforward example how to run the [Doitpay Golang SDK](https://github.com/automotechnologies/doitpay-go).


### Prerequsites
The sample system requires Go 1.20 or higher on the local machine in order to run the binary.

#### Go Programming Language
You need to have golang 1.20 installed in your machine.
Follow this step if you don't have golang 1.20 on your machine :
- Download the Go 1.20 binary package from the official Go website (https://golang.org/dl/).
- Install the package by following the instructions provided during the installation process.
- Once the installation is complete, verify that Go has been installed correctly by running the following command in your terminal: "go version".

## How to run locally
### Clone Repository:
Once you have all the prerequisites properly installed, you can start by cloning this repository.
- To clone the repository, run the following command in your terminal:
```
git clone https://github.com/automotechnologies/samples.git
```
- To navigate to the repository directory, run the following command in your terminal:
```
cd samples
```
Note: These instructions assume that you have Git installed on your machine. If you don't have Git installed, you can follow the instructions on the Git website to install it.

### Change Environment Variable
After cloning the repository, you can customize the environment variables in the .env file. Specifically, you should replace the placeholder API_KEY with your own API key. If you don't have an API key, you can create one by visiting the following link:
```
[TODO] Update this to the valid link
sandbox:
https://dev.doit-frontend.pages.dev/settings?type=api
production:
https://dev.doit-frontend.pages.dev/settings?type=api
```
Additionally, you may want to update the CHECKOUT_LINK based on your environment:
```
[TODO] Update this to the valid link
sandbox:
https://dev.doit-frontend.pages.dev/checkout/ 

production:
https://dev.doit-frontend.pages.dev/checkout/
```


### Running Binary:
Once you have cloned the repository and set up the dependencies, you can run the binary using either of the following methods:

Run vendor to download package dependencies

```
go mod vendor
```

Note: If you have change the docker config please change the config in /config/config.yaml before run it

And run using :

```
go run ./main.go
```

### Postman Collection
You can import postman collection in Repo File with Name : 
```
Doit Sample API.postman_collection.json
```

Note: The details mentioned in these steps may vary depending on your configuration.

### Setup Webhook Simulation:
In this guide, we will walk through the steps to set up webhook simulation using ngrok.

- First, download ngrok from its official website: [https://ngrok.com/download](https://ngrok.com/download).
- Register your ngrok account on the official ngrok website to obtain an authentication token. You will need this token to run ngrok.

To set up your authentication token, run the following command:
```
ngrok config add-authtoken <YOUR_AUTH_TOKEN>
```

Once you have added the authentication token, you can start ngrok with the following command:
```
ngrok http 8000
```

After running the command, ngrok will provide you with a URL. To use this as your webhook URL for simulation, append /api/v1/webhook to it."

```
Example : https://4b2e-118-99-106-57.ngrok-free.app/api/v1/webhook
```

This revised version provides clear steps and includes the command to add the authentication token as well as how to use the ngrok-generated URL for your webhook simulation.
