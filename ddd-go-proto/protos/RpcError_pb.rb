# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: RpcError.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "poc.model.RpcError" do
    optional :error_code, :string, 1
    optional :message, :string, 2
    repeated :validation_errors, :message, 3, "poc.model.ValidationError"
  end
  add_message "poc.model.ValidationError" do
    optional :field, :string, 1
    optional :restriction, :string, 2
    optional :message, :string, 3
  end
end

module Poc
  module Model
    RpcError = Google::Protobuf::DescriptorPool.generated_pool.lookup("poc.model.RpcError").msgclass
    ValidationError = Google::Protobuf::DescriptorPool.generated_pool.lookup("poc.model.ValidationError").msgclass
  end
end