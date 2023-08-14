// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file tendermint/libs/bits/types.proto (package tendermint.libs.bits, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message tendermint.libs.bits.BitArray
 */
export class BitArray extends Message<BitArray> {
  /**
   * @generated from field: int64 bits = 1;
   */
  bits = protoInt64.zero;

  /**
   * @generated from field: repeated uint64 elems = 2;
   */
  elems: bigint[] = [];

  constructor(data?: PartialMessage<BitArray>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tendermint.libs.bits.BitArray";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bits", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "elems", kind: "scalar", T: 4 /* ScalarType.UINT64 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BitArray {
    return new BitArray().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BitArray {
    return new BitArray().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BitArray {
    return new BitArray().fromJsonString(jsonString, options);
  }

  static equals(a: BitArray | PlainMessage<BitArray> | undefined, b: BitArray | PlainMessage<BitArray> | undefined): boolean {
    return proto3.util.equals(BitArray, a, b);
  }
}

