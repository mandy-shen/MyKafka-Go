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
         * if blank, default is 'add-row neu-student, My Name, ID0000001'.
         * it could be changed when typing other inputs starting from add-row
4. command examples
   * docker run mandyshen/mykafka
   * docker run mandyshen/mykafka 'add-row neu-student, Julia, ID0000123'
5. see Capture.png
   