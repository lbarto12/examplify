import { createApiClient } from "../genapi/sessions";
import { PUBLIC_API_BASE_PUBLIC } from '$env/static/public';

let client = createApiClient(PUBLIC_API_BASE_PUBLIC, {});

export function getClient() {
  return client;
}
