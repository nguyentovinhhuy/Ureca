set -ev

echo "Start adding peer1 and peer2"

docker-compose -f docker-compose-peer1.yaml up -d
docker-compose -f docker-compose-peer2.yaml up -d

echo "Added peer1 and peer2 container"
docker ps 

echo "Timeout for 5 secs"
export TEMP_TIMEOUT=5
sleep ${TEMP_TIMEOUT}

docker cp peer_join.sh cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/peer_join.sh 

docker exec cli ./peer_join.sh