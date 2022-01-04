# kafka_producer_consumer_app
This application is for Publishing and Consuming live messages using event streaming platfrom kafka and golang

## How to run 
  - First run docker file using command 
    ```
    $ docker-compose up
    ```
      > - *do not kill the terminal*
      > - *this docker file will form a cluster for your application so that you consumer and producer can communicate.*
      > - *once the cluster is created you can run your consumer.go and producer.go files* 
      
      
  - Open new CLI then run consumer.go file using command
    ```
    $ go run consumer.go
    ```
      - type the topic name on which you have to get message
      ```
      Enter Topic Name have to get message
      $ example_topic
      ```
  - Open new CLI then run producer.go file using command
    ```
    $ go run producer.go
    ```
      - type the topic name on which you have to publish message in that format
        ```
        To Enter topic name type ===> /topic <topic_name>
        $ /topic example_topic
        ```
      - type msg which you want to publish
        ```
         To message in example_topic type ===> /msg <your message>
         $ /msg Hii everyone this is example message !
        ```
     
  > Now you can see the message on consumer cli
## Note
  > - you can create multiple consumer who can subscribe diffrent topics 
  > - and also can create multiple producer who can write msg on a different topics
  > - for creating new consumer or producer just open new terminal and run consumer.go or producer.go scripts. 
  

    
