// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file ethermint/feemarket/v1/query.proto (package ethermint.feemarket.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Params } from "./feemarket_pb.js";

/**
 * QueryParamsRequest defines the request type for querying x/evm parameters.
 *
 * @generated from message ethermint.feemarket.v1.QueryParamsRequest
 */
export class QueryParamsRequest extends Message<QueryParamsRequest> {
  constructor(data?: PartialMessage<QueryParamsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryParamsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryParamsRequest {
    return new QueryParamsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryParamsRequest | PlainMessage<QueryParamsRequest> | undefined, b: QueryParamsRequest | PlainMessage<QueryParamsRequest> | undefined): boolean {
    return proto3.util.equals(QueryParamsRequest, a, b);
  }
}

/**
 * QueryParamsResponse defines the response type for querying x/evm parameters.
 *
 * @generated from message ethermint.feemarket.v1.QueryParamsResponse
 */
export class QueryParamsResponse extends Message<QueryParamsResponse> {
  /**
   * params define the evm module parameters.
   *
   * @generated from field: ethermint.feemarket.v1.Params params = 1;
   */
  params?: Params;

  constructor(data?: PartialMessage<QueryParamsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryParamsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "params", kind: "message", T: Params },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryParamsResponse {
    return new QueryParamsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryParamsResponse | PlainMessage<QueryParamsResponse> | undefined, b: QueryParamsResponse | PlainMessage<QueryParamsResponse> | undefined): boolean {
    return proto3.util.equals(QueryParamsResponse, a, b);
  }
}

/**
 * QueryBaseFeeRequest defines the request type for querying the EIP1559 base
 * fee.
 *
 * @generated from message ethermint.feemarket.v1.QueryBaseFeeRequest
 */
export class QueryBaseFeeRequest extends Message<QueryBaseFeeRequest> {
  constructor(data?: PartialMessage<QueryBaseFeeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryBaseFeeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBaseFeeRequest {
    return new QueryBaseFeeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBaseFeeRequest {
    return new QueryBaseFeeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBaseFeeRequest {
    return new QueryBaseFeeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBaseFeeRequest | PlainMessage<QueryBaseFeeRequest> | undefined, b: QueryBaseFeeRequest | PlainMessage<QueryBaseFeeRequest> | undefined): boolean {
    return proto3.util.equals(QueryBaseFeeRequest, a, b);
  }
}

/**
 * BaseFeeResponse returns the EIP1559 base fee.
 *
 * @generated from message ethermint.feemarket.v1.QueryBaseFeeResponse
 */
export class QueryBaseFeeResponse extends Message<QueryBaseFeeResponse> {
  /**
   * @generated from field: string base_fee = 1;
   */
  baseFee = "";

  constructor(data?: PartialMessage<QueryBaseFeeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryBaseFeeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "base_fee", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBaseFeeResponse {
    return new QueryBaseFeeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBaseFeeResponse {
    return new QueryBaseFeeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBaseFeeResponse {
    return new QueryBaseFeeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBaseFeeResponse | PlainMessage<QueryBaseFeeResponse> | undefined, b: QueryBaseFeeResponse | PlainMessage<QueryBaseFeeResponse> | undefined): boolean {
    return proto3.util.equals(QueryBaseFeeResponse, a, b);
  }
}

/**
 * QueryBlockGasRequest defines the request type for querying the EIP1559 base
 * fee.
 *
 * @generated from message ethermint.feemarket.v1.QueryBlockGasRequest
 */
export class QueryBlockGasRequest extends Message<QueryBlockGasRequest> {
  constructor(data?: PartialMessage<QueryBlockGasRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryBlockGasRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBlockGasRequest {
    return new QueryBlockGasRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBlockGasRequest {
    return new QueryBlockGasRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBlockGasRequest {
    return new QueryBlockGasRequest().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBlockGasRequest | PlainMessage<QueryBlockGasRequest> | undefined, b: QueryBlockGasRequest | PlainMessage<QueryBlockGasRequest> | undefined): boolean {
    return proto3.util.equals(QueryBlockGasRequest, a, b);
  }
}

/**
 * QueryBlockGasResponse returns block gas used for a given height.
 *
 * @generated from message ethermint.feemarket.v1.QueryBlockGasResponse
 */
export class QueryBlockGasResponse extends Message<QueryBlockGasResponse> {
  /**
   * @generated from field: int64 gas = 1;
   */
  gas = protoInt64.zero;

  constructor(data?: PartialMessage<QueryBlockGasResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "ethermint.feemarket.v1.QueryBlockGasResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "gas", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): QueryBlockGasResponse {
    return new QueryBlockGasResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): QueryBlockGasResponse {
    return new QueryBlockGasResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): QueryBlockGasResponse {
    return new QueryBlockGasResponse().fromJsonString(jsonString, options);
  }

  static equals(a: QueryBlockGasResponse | PlainMessage<QueryBlockGasResponse> | undefined, b: QueryBlockGasResponse | PlainMessage<QueryBlockGasResponse> | undefined): boolean {
    return proto3.util.equals(QueryBlockGasResponse, a, b);
  }
}

