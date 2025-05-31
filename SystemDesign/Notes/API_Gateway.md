
Q: Is API Gateway Single point of Failure (SPF) ?
Ans: No, AWS API gateways  Regional service specific which itself has multiple AZ's , so a API G/W routes load b/w diff AZ's and each region has it's own API Gateway and these G/Ws are behind DNS service which routes b/w diff regions


Regional Endpoint: When you create an API Gateway in AWS, it's a regional service. This means that the API Gateway and its underlying infrastructure are deployed within the specific AWS region you choose (e.g., us-east-1, eu-west-2, ap-south-1 - which includes Bengaluru). You get a regional endpoint for your API, something like https://your-api-id.execute-api.your-region.amazonaws.com.

DNS Resolution: When a client (anywhere in the world) makes a request to your API's regional endpoint, the DNS system resolves this regional endpoint to one of the IP addresses associated with the API Gateway in that specific region.

AWS Network Routing: Once the DNS resolves to an IP address within the chosen region, the request enters the AWS network. AWS's internal network infrastructure is designed to efficiently route traffic within a region.

Internal Load Balancing (Managed by AWS): Within the AWS API Gateway service in that region, there are internal load balancing mechanisms managed by AWS. These mechanisms are responsible for distributing the incoming traffic across the multiple API Gateway instances that are running in different Availability Zones within that region.

Traffic Distribution Across AZs: The internal load balancing within API Gateway ensures that requests are spread across the healthy API Gateway instances in all the Availability Zones where it's deployed in that region. This is transparent to the client. The client sends a request to the regional endpoint, and AWS takes care of routing it to an available API Gateway instance in any of the healthy AZs within that region.


# API Gateway
+-----------------+       HTTPS/HTTP       +-----------------+
|     Client      |--------------------->  |   API Gateway   |
|(e.g., Browser,  |                        | - WAF           |
| Mobile App, etc.)|                       | - Routing       |
+-----------------+                        | - Auth/Authz    |
                                           | - Rate Limiting |
                                           | - Transformation|
                                           | - Composition   |
                                           | - Monitoring    |
                                           +--------+--------+
                                                |        | (Internal Network)
                                       HTTP/HTTPS|        | HTTP/HTTPS
                                                v        v
                                       +-----------------+   +-----------------+
                                       | Microservice A  |   | Microservice B  |
                                       |(e.g., User     |    |(e.g., Orders   |
                                       |  Service)       |   |  Service)       |
                                       +-----------------+   +-----------------+
                                                ^
                                                | HTTP/HTTPS
                                       +-----------------+
                                       | Microservice C  |
                                       |(e.g., Inventory|
                                       |  Service)       |
                                       +-----------------+


# Load Balancer
+-----------------+       HTTP/HTTPS     +-----------------+
|     Client      |--------------------->|  Load Balancer  |
+-----------------+                      +--------+--------+
                                                |         |-------------
                                       HTTP/HTTP|                      |  HTTP/HTTPS
                                                v                      v 
                                       +-----------------+   +-----------------+
                                       | Microservice A  |   | Microservice A  |
                                       |(Instance 1)     |   |(Instance 2)    |
                                       +-----------------+   +-----------------+

# TLS(Transport layer security)
-------------------------------------------------------------------------------------------
| Feature           | TLS (One-way Authentication) | mTLS (Two-way/Mutual Authentication) |
| :---------------- | :--------------------------- | :----------------------------------- |
| Authentication    | Server authenticates to client | Client and server authenticate each other |
| Security          | Good for server verification   | Enhanced security due to mutual verification |
| Complexity        | Simpler to implement         | More complex due to client certificate management |
| Use Cases         | Public websites (HTTPS)      | API security, microservices, IoT, B2B communication, Zero Trust environments |
