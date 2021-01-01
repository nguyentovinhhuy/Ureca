set -ev

export ORDERER_CA="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/users/Admin@example.com/msp/cacerts/ca.example.com-cert.pem"

export CORE_PEER_ADDRESS="peer0.org1.example.com:7051"
peer chaincode install -n sample -p github.com/sample_chaincode/ -v 1.0

export CORE_PEER_ADDRESS="peer1.org1.example.com:7051"
peer chaincode install -n sample -p github.com/sample_chaincode/ -v 1.0

export CORE_PEER_ADDRESS="peer2.org1.example.com:7051"
peer chaincode install -n sample -p github.com/sample_chaincode/ -v 1.0

#instantiate chaicode
peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n sample -v 1.0 -l golang -c '{"Args":["a", "100"]}'