/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file gitpod/experimental/v1/scm.proto (package gitpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetSuggestedRepoURLsRequest, GetSuggestedRepoURLsResponse } from "./scm_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gitpod.experimental.v1.SCMService
 */
export const SCMService = {
  typeName: "gitpod.experimental.v1.SCMService",
  methods: {
    /**
     * GetSuggestedRepoURLs returns a list of suggested repositories to open for
     * the user.
     *
     * @generated from rpc gitpod.experimental.v1.SCMService.GetSuggestedRepoURLs
     */
    getSuggestedRepoURLs: {
      name: "GetSuggestedRepoURLs",
      I: GetSuggestedRepoURLsRequest,
      O: GetSuggestedRepoURLsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
