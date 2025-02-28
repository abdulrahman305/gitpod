// Copyright (c) 2025 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by connect-kotlin. DO NOT EDIT.
//
// Source: gitpod/v1/token.proto
//
package io.gitpod.publicapi.v1

import com.connectrpc.Headers
import com.connectrpc.MethodSpec
import com.connectrpc.ProtocolClientInterface
import com.connectrpc.ResponseMessage
import com.connectrpc.StreamType

public class TokenServiceClient(
  private val client: ProtocolClientInterface,
) : TokenServiceClientInterface {
  /**
   *  CreateUserToken creates a new temporary access token for the specified user.
   *  +admin – only to be used by installation admins
   */
  override suspend fun createTemporaryAccessToken(request: Token.CreateTemporaryAccessTokenRequest,
      headers: Headers): ResponseMessage<Token.CreateTemporaryAccessTokenResponse> = client.unary(
    request,
    headers,
    MethodSpec(
    "gitpod.v1.TokenService/CreateTemporaryAccessToken",
      io.gitpod.publicapi.v1.Token.CreateTemporaryAccessTokenRequest::class,
      io.gitpod.publicapi.v1.Token.CreateTemporaryAccessTokenResponse::class,
      StreamType.UNARY,
    ),
  )

}
