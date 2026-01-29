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
    thumbnailURL: z.string().url().optional(),
  })
  .passthrough();
const Documents = z.array(Document);
const AnalyzeCollectionRequest = z
  .object({ type: z.enum(["summary", "flashcards", "quiz", "deep_summary"]) })
  .passthrough();
const CollectionAnalysis = z
  .object({
    id: z.string().uuid(),
    type: z.enum(["summary", "flashcards", "quiz", "deep_summary"]),
    result: z.string(),
  })
  .passthrough();
const CollectionAnalyses = z.array(CollectionAnalysis);
const NewCourseRequest = z.object({ name: z.string() }).passthrough();
const NewCourseResponse = z.object({ courseName: z.string() }).passthrough();
const CourseNames = z.array(z.string());
const CollectionNames = z.array(Collection);

export const schemas = {
  NewCollectionRequest,
  NewCollectionResponse,
  Collection,
  Collections,
  UploadFileRequest,
  UploadFileResponse,
  Document,
  Documents,
  AnalyzeCollectionRequest,
  CollectionAnalysis,
  CollectionAnalyses,
  NewCourseRequest,
  NewCourseResponse,
  CourseNames,
  CollectionNames,
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
    path: "/core/collection/:id/analyses",
    alias: "getCollectionAnalyses",
    description: `Returns all AI analyses that have been generated
for the given collection.
`,
    requestFormat: "json",
    parameters: [
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z.array(CollectionAnalysis),
    errors: [
      {
        status: 404,
        description: `Collection not found`,
        schema: z.void(),
      },
      {
        status: 500,
        description: `Failed to retrieve analyses`,
        schema: z.void(),
      },
    ],
  },
  {
    method: "get",
    path: "/core/collection/:id/analysis/:analysisID",
    alias: "getAnalysis",
    description: `Returns all AI analyses that have been generated
for the given collection.
`,
    requestFormat: "json",
    parameters: [
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
      {
        name: "analysisID",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z
      .object({
        id: z.string().uuid(),
        type: z.enum(["summary", "flashcards", "quiz", "deep_summary"]),
        result: z.string(),
      })
      .passthrough(),
    errors: [
      {
        status: 404,
        description: `Collection not found`,
        schema: z.void(),
      },
      {
        status: 500,
        description: `Failed to retrieve analyses`,
        schema: z.void(),
      },
    ],
  },
  {
    method: "post",
    path: "/core/collection/:id/analyze",
    alias: "analyzeCollection",
    description: `Performs an AI analysis on all documents in a collection.
The analysis is run on a snapshot of the collection content.
`,
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: AnalyzeCollectionRequest,
      },
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z
      .object({
        id: z.string().uuid(),
        type: z.enum(["summary", "flashcards", "quiz", "deep_summary"]),
        result: z.string(),
      })
      .passthrough(),
    errors: [
      {
        status: 400,
        description: `Invalid request`,
        schema: z.void(),
      },
      {
        status: 404,
        description: `Collection not found`,
        schema: z.void(),
      },
      {
        status: 500,
        description: `Analysis failed`,
        schema: z.void(),
      },
    ],
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
    method: "get",
    path: "/core/collections/:id/documents",
    alias: "getCollectionDocuments",
    requestFormat: "json",
    parameters: [
      {
        name: "id",
        type: "Path",
        schema: z.string().uuid(),
      },
    ],
    response: z.array(Document),
  },
  {
    method: "post",
    path: "/core/course",
    alias: "newCourse",
    requestFormat: "json",
    parameters: [
      {
        name: "body",
        type: "Body",
        schema: z.object({ name: z.string() }).passthrough(),
      },
    ],
    response: z.object({ courseName: z.string() }).passthrough(),
  },
  {
    method: "get",
    path: "/core/course/:courseID/collections",
    alias: "getCourseCollections",
    requestFormat: "json",
    parameters: [
      {
        name: "courseID",
        type: "Path",
        schema: z.string(),
      },
    ],
    response: z.array(Collection),
  },
  {
    method: "get",
    path: "/core/courses",
    alias: "getCourses",
    requestFormat: "json",
    response: z.array(z.string()),
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
        thumbnailURL: z.string().url().optional(),
      })
      .passthrough(),
  },
]);

export const api = new Zodios(endpoints);

export function createApiClient(baseUrl: string, options?: ZodiosOptions) {
  return new Zodios(baseUrl, endpoints, options);
}
