/**
 * Copyright (c) 2025 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file gitpod/v1/envvar.proto (package gitpod.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateConfigurationEnvironmentVariableRequest, CreateConfigurationEnvironmentVariableResponse, CreateOrganizationEnvironmentVariableRequest, CreateOrganizationEnvironmentVariableResponse, CreateUserEnvironmentVariableRequest, CreateUserEnvironmentVariableResponse, DeleteConfigurationEnvironmentVariableRequest, DeleteConfigurationEnvironmentVariableResponse, DeleteOrganizationEnvironmentVariableRequest, DeleteOrganizationEnvironmentVariableResponse, DeleteUserEnvironmentVariableRequest, DeleteUserEnvironmentVariableResponse, ListConfigurationEnvironmentVariablesRequest, ListConfigurationEnvironmentVariablesResponse, ListOrganizationEnvironmentVariablesRequest, ListOrganizationEnvironmentVariablesResponse, ListUserEnvironmentVariablesRequest, ListUserEnvironmentVariablesResponse, ResolveWorkspaceEnvironmentVariablesRequest, ResolveWorkspaceEnvironmentVariablesResponse, UpdateConfigurationEnvironmentVariableRequest, UpdateConfigurationEnvironmentVariableResponse, UpdateOrganizationEnvironmentVariableRequest, UpdateOrganizationEnvironmentVariableResponse, UpdateUserEnvironmentVariableRequest, UpdateUserEnvironmentVariableResponse } from "./envvar_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gitpod.v1.EnvironmentVariableService
 */
export const EnvironmentVariableService = {
  typeName: "gitpod.v1.EnvironmentVariableService",
  methods: {
    /**
     * ListUserEnvironmentVariables returns all environment variables for the
     * authenticated user.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.ListUserEnvironmentVariables
     */
    listUserEnvironmentVariables: {
      name: "ListUserEnvironmentVariables",
      I: ListUserEnvironmentVariablesRequest,
      O: ListUserEnvironmentVariablesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateUserEnvironmentVariable updates an environment variable for the
     * authenticated user.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.UpdateUserEnvironmentVariable
     */
    updateUserEnvironmentVariable: {
      name: "UpdateUserEnvironmentVariable",
      I: UpdateUserEnvironmentVariableRequest,
      O: UpdateUserEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateUserEnvironmentVariable creates a new environment variable for the
     * authenticated user.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.CreateUserEnvironmentVariable
     */
    createUserEnvironmentVariable: {
      name: "CreateUserEnvironmentVariable",
      I: CreateUserEnvironmentVariableRequest,
      O: CreateUserEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteUserEnvironmentVariable deletes an environment variable for the
     * authenticated user.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.DeleteUserEnvironmentVariable
     */
    deleteUserEnvironmentVariable: {
      name: "DeleteUserEnvironmentVariable",
      I: DeleteUserEnvironmentVariableRequest,
      O: DeleteUserEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListConfigurationEnvironmentVariables returns all environment variables in
     * a configuration.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.ListConfigurationEnvironmentVariables
     */
    listConfigurationEnvironmentVariables: {
      name: "ListConfigurationEnvironmentVariables",
      I: ListConfigurationEnvironmentVariablesRequest,
      O: ListConfigurationEnvironmentVariablesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateConfigurationEnvironmentVariable updates an environment variable in
     * a configuration.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.UpdateConfigurationEnvironmentVariable
     */
    updateConfigurationEnvironmentVariable: {
      name: "UpdateConfigurationEnvironmentVariable",
      I: UpdateConfigurationEnvironmentVariableRequest,
      O: UpdateConfigurationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateConfigurationEnvironmentVariable creates a new environment variable
     * in a configuration.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.CreateConfigurationEnvironmentVariable
     */
    createConfigurationEnvironmentVariable: {
      name: "CreateConfigurationEnvironmentVariable",
      I: CreateConfigurationEnvironmentVariableRequest,
      O: CreateConfigurationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteConfigurationEnvironmentVariable deletes an environment variable in
     * a configuration.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.DeleteConfigurationEnvironmentVariable
     */
    deleteConfigurationEnvironmentVariable: {
      name: "DeleteConfigurationEnvironmentVariable",
      I: DeleteConfigurationEnvironmentVariableRequest,
      O: DeleteConfigurationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListOrganizationEnvironmentVariables returns all environment variables in
     * an organization.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.ListOrganizationEnvironmentVariables
     */
    listOrganizationEnvironmentVariables: {
      name: "ListOrganizationEnvironmentVariables",
      I: ListOrganizationEnvironmentVariablesRequest,
      O: ListOrganizationEnvironmentVariablesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateOrganizationEnvironmentVariable updates an environment variable in
     * an organization.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.UpdateOrganizationEnvironmentVariable
     */
    updateOrganizationEnvironmentVariable: {
      name: "UpdateOrganizationEnvironmentVariable",
      I: UpdateOrganizationEnvironmentVariableRequest,
      O: UpdateOrganizationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CreateOrganizationEnvironmentVariable creates a new environment variable
     * in an organization.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.CreateOrganizationEnvironmentVariable
     */
    createOrganizationEnvironmentVariable: {
      name: "CreateOrganizationEnvironmentVariable",
      I: CreateOrganizationEnvironmentVariableRequest,
      O: CreateOrganizationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteOrganizationEnvironmentVariable deletes an environment variable in
     * an organization.
     *
     * @generated from rpc gitpod.v1.EnvironmentVariableService.DeleteOrganizationEnvironmentVariable
     */
    deleteOrganizationEnvironmentVariable: {
      name: "DeleteOrganizationEnvironmentVariable",
      I: DeleteOrganizationEnvironmentVariableRequest,
      O: DeleteOrganizationEnvironmentVariableResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc gitpod.v1.EnvironmentVariableService.ResolveWorkspaceEnvironmentVariables
     */
    resolveWorkspaceEnvironmentVariables: {
      name: "ResolveWorkspaceEnvironmentVariables",
      I: ResolveWorkspaceEnvironmentVariablesRequest,
      O: ResolveWorkspaceEnvironmentVariablesResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
