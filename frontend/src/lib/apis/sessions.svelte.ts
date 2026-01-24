import { createApiClient } from "../genapi/sessions";
import { PUBLIC_API_BASE_PUBLIC } from '$env/static/public';
import type { AnyZodiosRequestOptions, ZodiosEndpointDefinitions, ZodiosPlugin } from "@zodios/core";
import type { DeepReadonlyObject } from "@zodios/core/lib/utils.types";

let client = createApiClient(PUBLIC_API_BASE_PUBLIC, {});

// Auth plugin using cookieStore (browser only)
const authPlugin: ZodiosPlugin = {
  request: async (
    api: ZodiosEndpointDefinitions,
    config: DeepReadonlyObject<AnyZodiosRequestOptions>
  ) => {
    const auth = await cookieStore.get("auth");
    if (!auth) return config;

    const mutableConfig = { ...(config as AnyZodiosRequestOptions) };
    mutableConfig.headers = {
      ...mutableConfig.headers,
      Authorization: `Bearer ${auth.value}`,
    };

    return mutableConfig;
  }
};

export function getClient() {
  client.use(authPlugin);
  return client;
}
