import type { AnyZodiosRequestOptions, ZodiosEndpointDefinitions, ZodiosPlugin } from "@zodios/core";
import type { DeepReadonlyObject } from "@zodios/core/lib/utils.types";

const authPlugin: ZodiosPlugin = {
  request: async (api: ZodiosEndpointDefinitions, config: DeepReadonlyObject<AnyZodiosRequestOptions>) => {
    const cookie = document.cookie
      .split("; ")
      .find(row => row.startsWith("auth="))
      ?.split("=")[1];

    if (!cookie) return config;

    const mutableConfig = { ...(config as AnyZodiosRequestOptions) };
    mutableConfig.headers = {
      ...mutableConfig.headers,
      Authorization: `Bearer ${cookie}`,
    };

    return mutableConfig;
  },
};

export default authPlugin;