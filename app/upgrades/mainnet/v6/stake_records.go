package v6

// jq '.delegation_responses | map({address:.delegation.delegator_address,amount:((.balance.amount | tonumber)*0.05*((0.42/365)*13+1) | floor) | tostring})' DAN.JSON > to_mint.json

// Slash was 5%
// Lost APR is 42% for 13 days

var recordsJSONString = `[
 {
    "address": "comdex1tq2gr3vv6x30g607r5pf6jjen50fkft6u56zw8",
    "amount": "50000000"
  },
  {
    "address": "comdex1u7qz9yruwpqnq8xkaahazq4z54vf5w58zm4l83",
    "amount": "5000000"
  },
  {
    "address": "comdex1lmjtlmmqvxgnve23nw4xvce0d2xy0c24qk8ss9",
    "amount": "10000000"
  }
]`
