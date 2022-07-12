docker rm -f $(docker ps -aq)

# remove existed network
docker network rm  bitnetwork-br7
# create new network
docker network create --subnet=172.172.0.0/24 bitnetwork-br7

# remove pre data
sudo rm -rf ~/bitnetwork_dev
sudo cp -r ../bitnetwork_dev ~/bitnetwork_dev

# start validators


docker logs -f peer0 --tail=500
