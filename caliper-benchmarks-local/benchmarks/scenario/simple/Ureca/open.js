'use strict';

const { builtinModules } = require("module");

let bc, ctx, clientArgs;
let txnPerBatch, initMoney;

const characters = 'abcdefghijklmnopqrstuvwxyz';

module.exports.info = "Set random value";

module.exports.init = async function(blockchain, context, args) {
    if(!args.hasOwnProperty('money')) {
        return Promise.reject(new Error('simple.open - \'money\' is missed in the arguments'));
    }

    if(!args.hasOwnProperty('txnPerBatch')) {
        args.txnPerBatch = 1;
    }
    bc = blockchain;
    ctx = context;
    clientArgs = args;
    txnPerBatch = args.txnPerBatch;
    initMoney = args.money;

    return Promise.resolve();
}

module.exports.end = function() {
    return Promise.resolve();
}

function generateRandomVar(length) {
    let result = '';
    for ( var i = 0; i < length; i++ ) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    return result;
}

function generateWorkload() {
    let workload = [];
    for (let i = 0; i < txnPerBatch; i++){
        let varName = generateRandomVar(5);
        let value = (Math.floor(Math.random()));
        
        if (bc.getType() === 'fabric') {
            workload.push({
                chaincodeFunction: 'set',
                chaincodeArguments: [varName.toString(), value.toString()],
            });
        } else if (bc.getType() === 'ethereum') {
                workload.push({
                    verb: 'set',
                    args: [varName, value]
                });
        } else {
            workload.push({
                'verb': 'set',
                'account': varName,
                'money': value,
            });
        }
    }
}

module.exports.run = function() {
    let args = generateWorkload();
    return bc.invokeSmartContract(ctx, 'sample', 'v0', args);
}