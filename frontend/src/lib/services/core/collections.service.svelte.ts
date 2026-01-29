/**
 * Collections service
 */

import { BaseService } from '../shared/service-base';
import { coreApiClient } from '../shared/api-client';
import type { ServiceResult } from '../shared/types';
import { toastStore } from '../shared/toast-store.svelte';
import type { z } from 'zod';
import type { schemas } from '$lib/genapi/core';

type Collection = z.infer<typeof schemas.Collection>;
type NewCollectionRequest = z.infer<typeof schemas.NewCollectionRequest>;
type NewCollectionResponse = z.infer<typeof schemas.NewCollectionResponse>;
type AnalyzeCollectionRequest = z.infer<typeof schemas.AnalyzeCollectionRequest>;
type CollectionAnalysis = z.infer<typeof schemas.CollectionAnalysis>;

class CollectionsService extends BaseService {
	collections = $state<Collection[]>([]);
	currentCollection = $state<Collection | null>(null);
	analyses = $state<CollectionAnalysis[]>([]);
	loading = $state(false);
	analyzing = $state(false);
	error = $state<string | null>(null);

	/**
	 * Create a new collection
	 */
	async create(data: NewCollectionRequest): Promise<ServiceResult<NewCollectionResponse>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const response = await coreApiClient.newCollection(data);
			toastStore.success('Collection created successfully!');

			// Refresh collections for this course
			await this.getByCourse(data.course);

			return response;
		});

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Get collection by ID
	 */
	async getById(id: string): Promise<ServiceResult<Collection>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const collection = await coreApiClient.getCollection({ params: { id } });
			this.currentCollection = collection;
			return collection;
		});

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Get collections by course
	 */
	async getByCourse(courseId: string): Promise<ServiceResult<Collection[]>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const collections = await coreApiClient.getCourseCollections({
				params: { courseID: courseId }
			});
			this.collections = collections;
			return collections;
		}, false); // Don't show toast for read operations

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Filter collections by course and type
	 */
	async filter(courseId: string, type: string): Promise<ServiceResult<Collection[]>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const collections = await coreApiClient.filterCollections({
				params: { courseID: courseId, type }
			});
			this.collections = collections;
			return collections;
		}, false);

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Get analyses for a collection
	 */
	async getAnalyses(collectionId: string): Promise<ServiceResult<CollectionAnalysis[]>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const analyses = await coreApiClient.getCollectionAnalyses({
				params: { id: collectionId }
			});
			this.analyses = analyses;
			return analyses;
		}, false);

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Create analysis for a collection
	 */
	async analyze(
		collectionId: string,
		type: AnalyzeCollectionRequest['type']
	): Promise<ServiceResult<CollectionAnalysis>> {
		this.analyzing = true;
		this.error = null;

		const result = await this.execute(async () => {
			const analysis = await coreApiClient.analyzeCollection(
				{ type },
				{ params: { id: collectionId } }
			);

			toastStore.success('Analysis created successfully!');

			// Refresh analyses
			await this.getAnalyses(collectionId);

			return analysis;
		});

		this.analyzing = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}
}

export const collectionsService = new CollectionsService();
