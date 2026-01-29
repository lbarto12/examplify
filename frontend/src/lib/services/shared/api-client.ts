/**
 * Singleton API client setup for service layer
 */

import { createApiClient as createCoreApi } from '$lib/genapi/core';
import { createApiClient as createSessionsApi } from '$lib/genapi/sessions';
import { PUBLIC_API_BASE_PRIVATE, PUBLIC_API_BASE_PUBLIC } from '$env/static/public';
import authPlugin from '$lib/apis/utils/authPlugin';

/**
 * Core API client (authenticated) - for collections, documents, analyses
 */
export const coreApiClient = createCoreApi(PUBLIC_API_BASE_PRIVATE, {});
coreApiClient.use(authPlugin);

/**
 * Sessions API client (public) - for authentication
 */
export const sessionsApiClient = createSessionsApi(PUBLIC_API_BASE_PUBLIC, {});
