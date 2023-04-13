{{
# Creating docker image for logger service
docker build -f logger-service.dockerfile -t javohir11111/logger-service:1.0.0 .

# For pushing image use this command
docker push javohir11111/logger-service:1.0.0

// We have to create new docker image and push to docker hub like these two comment for every service
}}

# Docker auth in terminal. Docker hub is at this link https://hub.docker.com/
docker login

#
docker swarm init

# Creating network and services on Docker HUB
docker stack deploy -c swarm.yml myapp

docker service ls

docker service scale myapp_listener-service=3

docker service update --image javohir11111/logger-service:1.0.1 myapp_logger-service

# Removing swarm container from docker dashboard
docker service scale myapp_broker-service=0
docker stack rm myapp 

docker swarm leave