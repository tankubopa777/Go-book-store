::: mermaid
graph TD
    client[Client] -->|http| ingress[Ingress/Service]

    ingress -->|grpc| middleware[Middleware]
    
    subgraph k8s[Kubernetes Cluster]
    
        middleware -->|grpc| user[User Service]
        middleware -->|grpc| auth[Auth Service]
        middleware -->|grpc| book[Book Service]
        middleware -->|pub| payment[Payment Service]

        
        user -->|sub| kafka[Kafka]
        auth -->|sub| kafka
        book -->|sub| kafka
        payment -->|sub| kafka

        kafka -->|pub| userbooks[Userbooks Service]

        user --> userdb[(User DB)]
        auth --> authdb[(Auth DB)]
        book --> bookdb[(Book DB)]
        userbooks --> userbooksdb[(Userbooks DB)]
        payment --> paymentdb[(Payment DB)]
    end


:::