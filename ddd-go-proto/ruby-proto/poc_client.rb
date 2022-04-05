require 'faraday'
require 'faraday/net_http'
require 'json'
require '../protos/CreateAuthenticationClient_pb'
require 'pry'

Faraday.default_adapter = :net_http

class PocClient
    def initialize
        @conn_json = Faraday.new(
            url: 'http://localhost:8080',
            headers: {'Content-Type' => 'application/json'}
          )
        @conn_proto = Faraday.new(
            url: 'http://localhost:8888',
            headers: {'Content-Type' => 'application/protobuf', 'Accept' => 'application/protobuf' }
          )
    end

    def create_authentication_code_json
        response = @conn_json.post('http://localhost:8080/authentication_clients') do |req|
            req.body = {
                "authentication_client_type":"JWT",
                "client_id":"111111",
                "platform_account_id":2222222,
                "client_secret":"YYYYYYYYYYYYYY",
                "application_uuid":"XXXXXXXXXXXX",
                "enabled":true,
                "user_id":11111112222,
                "api_key":"xxxxx_yyyyyyy_xxxxxx",
                "api_key_name":"my first api key"
             }.to_json
          end

        puts response.status
        puts response.body
    end

    def create_authentication_code_proto
        request = Poc::Model::CreateAuthenticationClient.new(
#             client_id: '123456789012390123456xxxxxx222',
            application_uuid: '123456789012345678901234567890123456',
            platform_account_id: 11222,
           authentication_client_type: 'PUBLIC_TOKEN',
           enabled: true,
           user_id: 112233,
           client_secret: "xxxxx_yyyyyyy_xxxxxx",
#            api_key_name: 'my-api',
#            "api_key": '123456789012345678901234567890123456'

        )
#         puts request
        create_auth_client_request = Poc::Model::CreateAuthenticationClient.encode(request)
#         puts create_auth_client_request

        response = @conn_proto.post('http://localhost:8888/authentication_clients', create_auth_client_request)

#         puts response
#         puts response.body

        begin
            create_auth_client_response = Poc::Model::CreateAuthenticationClientResponse.decode(response.body)
        rescue => e
            puts "RESCUE"
            puts e
        end

        puts create_auth_client_response
    end

    def encode
        begin
            response_type.encode(service_response.body)
        rescue Google::Protobuf::ParseError => error
            raise ServiceError, "raw_error: #{error}, body: #{service_response.inspect}"
        end
    end

    def decode
        request_type.encode(service_request)
    end


end


# PocClient.new().create_authentication_code_json
PocClient.new().create_authentication_code_proto
