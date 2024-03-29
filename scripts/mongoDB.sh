#!/bin/bash
#script to pull and run mongoDB in single node inside a docker container

# Set env var
hostname=localhost
port=27888
username=admin
password=pass

export DB_HOSTNAME=$hostname
export DB_PORT=$port
export DB_USERNAME=$username
export DB_PASSWORD=$password

# If mongoDB container is running, exit.
running_app_container=`docker ps | grep local-mongo | wc -l`
if [ $running_app_container -gt "0" ]
then
	printf "container already running... \n"
	printf "press any key to close the terminal \n"	
	read junk
	exit 1
fi
		
# If mongoDB container exists and is off, run.
existing_app_container=`docker ps -a | grep local-mongo | grep Exit | wc -l`
if [ $existing_app_container -gt "0" ]
then
	printf "starting container... \n"
	docker start local-mongo
	printf "container is now running \n"
	printf "press any key to close the terminal \n"
	read junk
	exit 1
fi
		
# Else, pull image and run.
printf "pulling image for MongoDB... \n\n"
docker pull mongo
printf "running container... \n\n"
docker run -d --name local-mongo -p $port:27017 -e MONGO_INITDB_ROOT_USERNAME=$username -e MONGO_INITDB_ROOT_PASSWORD=$password mongo
