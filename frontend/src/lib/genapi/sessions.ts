import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const SessionItem = z
  .object({ email: z.string().email(), password: z.string() })
  .passthrough();
const SessionResponse = z.object({ token: z.string() }).passthrough();

export const schemas = {
  SessionItem,
  SessionResponse,
};

const endpoints = makeApi([
  {
    method: "post",
    path: "/signin",
    alias: "signIn",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ email: z.string().email(), password: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ token: z.string() }).passthrough(),
  },
  {
    method: "post",
    path: "/signup",
    alias: "signUp",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ email: z.string().email(), password: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ token: z.string() }).passthrough(),
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
