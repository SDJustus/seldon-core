/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: v2_dataplane.proto

package io.seldon.mlops.inference.v2;

@kotlin.jvm.JvmSynthetic
public inline fun serverReadyRequest(block: io.seldon.mlops.inference.v2.ServerReadyRequestKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest =
  io.seldon.mlops.inference.v2.ServerReadyRequestKt.Dsl._create(io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest.newBuilder()).apply { block() }._build()
public object ServerReadyRequestKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest = _builder.build()
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest.copy(block: io.seldon.mlops.inference.v2.ServerReadyRequestKt.Dsl.() -> kotlin.Unit): io.seldon.mlops.inference.v2.V2Dataplane.ServerReadyRequest =
  io.seldon.mlops.inference.v2.ServerReadyRequestKt.Dsl._create(this.toBuilder()).apply { block() }._build()