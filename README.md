# Kafka Health Check Manager

Status page and alerting for 
[Kafka Health Check](https://github.com/andreas-schroeder/kafka-health-check)

## Use Case
Display a status page for those always asking the status of kafka. 
Send alerts when there is a problem.
Prometheus and Burrow are good for finite alerts. Kafka Health Check is good for a simple overall check and alert for impacted topics.

### Notifiers
* Slack
* SNS
* Web Hook
