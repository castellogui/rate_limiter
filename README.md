# Rate Limiter

### Introduction
The service that restricts the number of requests per user/IP within a given period of time. To make it distributed, it has been implemented as a microservice that different applications call to check whether a request can be accepted or not.
