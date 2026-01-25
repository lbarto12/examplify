import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const UploadFileRequest = z
  .object({ collectionID: z.string().uuid(), mimeType: z.string() })
  .passthrough();
const UploadFileResponse = z.object({ uploadURL: z.string() }).passthrough();

export const schemas = {
  UploadFileRequest,
  UploadFileResponse,
};

const endpoints = makeApi([
  {
    method: "post",
    path: "/core/upload-file",
    alias: "uploadFile",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ collectionID: z.string().uuid(), mimeType: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ uploadURL: z.string() }).passthrough(),
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
