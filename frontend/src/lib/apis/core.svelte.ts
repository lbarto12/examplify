import { createApiClient } from "../genapi/core";
import { PUBLIC_API_BASE_PRIVATE } from '$env/static/public';
import authPlugin from "./utils/authPlugin";

let client = createApiClient(PUBLIC_API_BASE_PRIVATE, {});

export function getClient() {
  client.use(authPlugin);
  return client;
}
