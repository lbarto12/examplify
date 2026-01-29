/**
 * Analyses service
 */

import { BaseService } from '../shared/service-base';
import { coreApiClient } from '../shared/api-client';
import type { ServiceResult } from '../shared/types';
import type { z } from 'zod';
import type { schemas } from '$lib/genapi/core';

type CollectionAnalysis = z.infer<typeof schemas.CollectionAnalysis>;

class AnalysesService extends BaseService {
	currentAnalysis = $state<CollectionAnalysis | null>(null);
	loading = $state(false);
	error = $state<string | null>(null);

	/**
	 * Get analysis by ID
	 */
	async getById(collectionId: string, analysisId: string): Promise<ServiceResult<CollectionAnalysis>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const analysis = await coreApiClient.getAnalysis({
				params: {
					id: collectionId,
					analysisID: analysisId
				}
			});

			// Parse result if it's a JSON string
			if (typeof analysis.result === 'string') {
				analysis.result = JSON.parse(analysis.result);
			}

			this.currentAnalysis = analysis;
			return analysis;
		}, false);

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Clear current analysis
	 */
	clear() {
		this.currentAnalysis = null;
		this.error = null;
	}
}

export const analysesService = new AnalysesService();
