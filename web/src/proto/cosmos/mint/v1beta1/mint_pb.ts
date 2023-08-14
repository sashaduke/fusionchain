// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file cosmos/mint/v1beta1/mint.proto (package cosmos.mint.v1beta1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * Minter represents the minting state.
 *
 * @generated from message cosmos.mint.v1beta1.Minter
 */
export class Minter extends Message<Minter> {
  /**
   * current annual inflation rate
   *
   * @generated from field: string inflation = 1;
   */
  inflation = "";

  /**
   * current annual expected provisions
   *
   * @generated from field: string annual_provisions = 2;
   */
  annualProvisions = "";

  constructor(data?: PartialMessage<Minter>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.mint.v1beta1.Minter";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "inflation", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "annual_provisions", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Minter {
    return new Minter().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Minter {
    return new Minter().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Minter {
    return new Minter().fromJsonString(jsonString, options);
  }

  static equals(a: Minter | PlainMessage<Minter> | undefined, b: Minter | PlainMessage<Minter> | undefined): boolean {
    return proto3.util.equals(Minter, a, b);
  }
}

/**
 * Params holds parameters for the mint module.
 *
 * @generated from message cosmos.mint.v1beta1.Params
 */
export class Params extends Message<Params> {
  /**
   * type of coin to mint
   *
   * @generated from field: string mint_denom = 1;
   */
  mintDenom = "";

  /**
   * maximum annual change in inflation rate
   *
   * @generated from field: string inflation_rate_change = 2;
   */
  inflationRateChange = "";

  /**
   * maximum inflation rate
   *
   * @generated from field: string inflation_max = 3;
   */
  inflationMax = "";

  /**
   * minimum inflation rate
   *
   * @generated from field: string inflation_min = 4;
   */
  inflationMin = "";

  /**
   * goal of percent bonded atoms
   *
   * @generated from field: string goal_bonded = 5;
   */
  goalBonded = "";

  /**
   * expected blocks per year
   *
   * @generated from field: uint64 blocks_per_year = 6;
   */
  blocksPerYear = protoInt64.zero;

  constructor(data?: PartialMessage<Params>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cosmos.mint.v1beta1.Params";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "mint_denom", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "inflation_rate_change", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "inflation_max", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "inflation_min", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "goal_bonded", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "blocks_per_year", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Params {
    return new Params().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Params {
    return new Params().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Params {
    return new Params().fromJsonString(jsonString, options);
  }

  static equals(a: Params | PlainMessage<Params> | undefined, b: Params | PlainMessage<Params> | undefined): boolean {
    return proto3.util.equals(Params, a, b);
  }
}

