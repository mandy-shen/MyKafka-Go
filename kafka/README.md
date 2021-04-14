# edu.neu.csye6225.assign4

## Required
* terminal
* docker

## Prepare environment
* the image is prepared publicly in Docker Hub, named "mandyshen/mykafka"

## How to docker run my image
1. open terminal
2. download the image
   * docker pull mandyshen/mykafka
3. run image's formula
   * docker run mandyshen/mykafka {param}
      * {param} :
         * if blank, default is 'add-row neu-student, My Name, ID0000001'
         * it could be changed when typing other inputs starting from "add-row"
         * ps. MacOS is 'add-row neu-student, My Name, ID0000001', but Windows is "add-row neu-student, My Name, ID0000001", different quotation marks
4. command examples
   * docker run mandyshen/mykafka
   * docker run mandyshen/mykafka 'add-row neu-student, Julia, ID0000123'
5. see Capture.png

## How to approach this assignment
### step of containerizing the software
* use example kafka as software
    * download kafka - https://kafka.apache.org/downloads.html
    * choose: Scala 2.12  - kafka_2.12-2.7.0.tgz
    * unzip kafka_2.12-2.7.0.tgz
* customize kafka's config - config/server.properties
    * unmarked advertised.listeners
        * advertised.listeners=PLAINTEXT://localhost:9092
* write executable files - kfk_start.sh & main.go
    * kfk_start.sh
        * make sure that zookeeper and kafka are starting in order
    * main.go
        * because the input is " docker run my-container-image 'add-row neu-student, My Name, ID0000001' "
        * use main.go get parameter, and then customized output through kafka
        * use main.go as container's entrypoint
* write Dockerfile
    * start from scratch
    * use multi-stage to create smaller binary
    * use mini os - alpine
    * leverage CMD as default parameter
    * jre is necessary - kafka base on jvm
* push to Docker Hub
    * push code to Github
    * use Github Actions to push to Docker Hub
        * create secrets
            * DOCKER_HUB_USERNAME >> Docker Hub ID
            * DOCKER_HUB_ACCESS_TOKEN >> Docker Hub's Access Token
                * account setting >> security >> New Access Token
        * create CI workflow by Github Actions UI
            * see https://docs.docker.com/ci-cd/github-actions/
            * example: pushDockerHub.yml
    
### another better way
