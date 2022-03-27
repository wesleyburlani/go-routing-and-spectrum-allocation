# Routing and Spectrum Allocation in Elastic Optical Networks Simulator

This project implements some algorithms to simulate Routing and Spectrum Allocation in Elastic Optical Networks using GoLang

As a brief explanation, given a Graph <em>g</em>, with nodes(<em>n</em>) that represents places and edges(<em>e</em>) that represents network optical links between these places and a demand <em>d</em>, to simulation will try to allocate this demand <em>d</em> into the Network using a Rouing and Spectrum Allocation Algorithm <em>a</em>, filling the routing table using a Table Fill algorithm <em>t</em>.

## Requirements
* [GoLang 1.7](https://go.dev/dl/)

OR 

* [Docker](https://www.docker.com/products/docker-desktop/) 
* [Docker Compose](https://docs.docker.com/compose/install/)


## Configuring the project

This project has the intention of being totally configurable in terms of which algorithm will be used on each of its steps. Because of that, there are some variables that must be set in order to configure your simulation. 

All configurations are centralized on a `.env` file, so, the first step is to copy the `.env.example` file and paste is into a `.env.` file on the root of the project. You can also run:

```
cp .env.example .env
```

Now, using the new created `.env` file you should configure the following variables:

**NODES_FILE_PATH**: the path of the file containing the nodes of the network (required)

**EDGES_FILE_PATH**: the path of the file containing the links of the network (required)

**LOG_TYPE**: configures where the logs must be written (required). Possible values: stdout, file
  
**LOG_FILE_PATH**: If log type is `file` configures the file path to log

**DEMANDS_SOURCE**: defines how the demands will be provided (required). Possible values: generate, file

**DEMANDS_FILE_PATH**: If demand source is `file` configures the file path to provide the demands and if demand source is `generate` saves the generated demands into this file.

**PATH_SEARCH_ALGORITHM**: defines which path search algorithm to use (required). Possible values: dijkstra.

**DISJOINTED_PATH_PAIR_SEARCH_ALGORITHM**: defines which disjointed path pair search algorithm to use (required). Possible values: suurballe

**TABLE_FILL_ALGORITHM**: defines which algorithm to use to fill the routing table (required). Possible values: first_fit_rsa, first_fit_rmlsa

**RSA_TYPE**: defines the type of the RSA algorithm that should be used. Possible values: single, dedicated_protection, shared_protection

The `.env.example` file is already configured with a valid configuration


## Running Locally

If you have installed Go Lang 1.7 locally, you can run the following command:

```
./run.sh
```

If you are using docker:

```
docker build . -t rsa-simulator 
docker run rsa-simulator 
```

If you are using docker-compose:

```
docker-compose build
docker-compose run rsa-simulator
```
