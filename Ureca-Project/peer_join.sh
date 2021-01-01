set -ev

peer channel fetch newest mychannel.block -c mychannel --orderer orderer.example.com:7050

export CHANNEL_NAME=mychannel

export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer1.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=peer1.org1.example.com:7051

echo "Joining peer1 to mychannel"
peer channel join -b mychannel.block

export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer2.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=peer2.org1.example.com:7051

echo "Joining peer2 to mychannel"
peer channel join -b mychannel.block

exit