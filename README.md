# go-echo-eureka-zuul-oauth2
Sample code for integrating labstack echo microservice with netflix eureka and zuul with spring boot oauth2 as the security handler

To add new service to Eureka use the following format

http://host:port/eureka/apps/{{service-name}}
{
    "instance": {
        "hostName": "{{hostname/ip}}", 
        "app": "{{service-name}}",
        "vipAddress": "{{service-name}}",
        "secureVipAddress": "{{service-name}}",
        "ipAddr": "{{ip-address}}",
        "status": "STARTING",
        "port": {"$": "{{http-port}}", "@enabled": "true"},
        "securePort": {"$": "{{https-port}}", "@enabled": "false"},
        "healthCheckUrl": "http://{{hostname/ip}}:{{http-port}}/healthcheck",
        "statusPageUrl": "http://{{hostname/ip}}:{{http-port}}/healthcheck",
        "homePageUrl": "http://{{hostname/ip}}:{{http-port}}",
        "dataCenterInfo": {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", 
            "name": "MyOwn"
        }
    }
}

I used following configuration

{
    "instance": {
        "hostName": "192.168.0.104", 
        "app": "goUserService",
        "vipAddress": "goUserService",
        "secureVipAddress": "goUserService",
        "ipAddr": "192.168.0.104",
        "status": "STARTING",
        "port": {"$": "1323", "@enabled": "true"},
        "securePort": {"$": "8443", "@enabled": "false"},
        "healthCheckUrl": "http://192.168.0.104:1323/healthcheck",
        "statusPageUrl": "http://192.168.0.104:1323/healthcheck",
        "homePageUrl": "http://192.168.0.104:1323",
        "dataCenterInfo": {
            "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo", 
            "name": "MyOwn"
        }
    }
}

