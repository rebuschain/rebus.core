<!--
order: 7
-->

# Params

Claim module provides below params

```protobuf
// Params defines the claim module's parameters.
message Params {
  google.protobuf.Timestamp airdrop_start_time = 1 [
    (gogoproto.stdtime) = true,
    (gogoproto.nullable) = false,
    (gogoproto.moretags) = "yaml:\"airdrop_start_time\""
  ];
  google.protobuf.Duration airdrop_duration = 2 [
    (gogoproto.nullable) = false,
    (gogoproto.stdduration) = true,
    (gogoproto.jsontag) = "airdrop_duration,omitempty",
    (gogoproto.moretags) = "yaml:\"airdrop_duration\""
  ];
  // denom of claimable asset
  string claim_denom = 3;
}
```

1. `airdrop_start_time` refers to the time when user can start to claim airdrop.
2. `airdrop_duration` refers to the airdrop duration.
3. `claim_denom` refers to the denomination of claiming tokens. As a default, it's `arebus`.