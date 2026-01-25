import { makeApi, Zodios, type ZodiosOptions } from "@zodios/core";
import { z } from "zod";

const NewCollectionRequest = z
  .object({ title: z.string(), course: z.string(), type: z.string() })
  .passthrough();
const NewCollectionResponse = z
  .object({ collectionID: z.string().uuid() })
  .passthrough();
const Collection = z
  .object({
    ID: z.string().uuid(),
    title: z.string(),
    course: z.string(),
    type: z.string(),
  })
  .passthrough();
const Collections = z.array(Collection);
const UploadFileRequest = z
  .object({ collectionID: z.string().uuid(), mimeType: z.string() })
  .passthrough();
const UploadFileResponse = z.object({ uploadURL: z.string() }).passthrough();
const Document = z
  .object({
    ID: z.string().uuid(),
    collectionID: z.string().uuid(),
    mimeType: z.string(),
    downloadURL: z.string().url(),
  })
  .passthrough();

export const schemas = {
  NewCollectionRequest,
  NewCollectionResponse,
  Collection,
  Collections,
  UploadFileRequest,
  UploadFileResponse,
  Document,
};

const endpoints = makeApi([
  {
    method: "post",
    path: "/core/collection",
    alias: "newCollection",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z
          .object({ title: z.string(), course: z.string(), type: z.string() })
          .passthrough(),
      },
    ],
    response: z.object({ collectionID: z.string().uuid() }).passthrough(),
  },
  {
    method: "get",
    path: "/core/collection/:id",
    alias: "getCollection",
    requestFormat: "json",
    parameters: [
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z
      .object({
        ID: z.string().uuid(),
        title: z.string(),
        course: z.string(),
        type: z.string(),
      })
      .passthrough(),
  },
  {
    method: "get",
    path: "/core/collections/:courseID/:type",
    alias: "filterCollections",
    requestFormat: "json",
    parameters: [
      {
        name: "courseID",
        type: "Path",
        schema: z.string(),
      },
      {
        name: "type",
        type: "Path",
        schema: z.string(),
      },
    ],
    response: z.array(Collection),
  },
  {
    method: "post",
    path: "/core/document",
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
  {
    method: "get",
    path: "/core/document/:id",
    alias: "getDocument",
    requestFormat: "json",
    parameters: [
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z
      .object({
        ID: z.string().uuid(),
        collectionID: z.string().uuid(),
        mimeType: z.string(),
        downloadURL: z.string().url(),
      })
      .passthrough(),
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
