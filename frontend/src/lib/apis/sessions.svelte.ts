import { createApiClient } from "../genapi/sessions";
import { PUBLIC_API_BASE_PUBLIC } from '$env/static/public';
import authPlugin from "./utils/authPlugin";

let client = createApiClient(PUBLIC_API_BASE_PUBLIC, {});

export function getClient() {
  client.use(authPlugin);
  return client;
}
