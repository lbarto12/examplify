/**
 * Documents service
 */

import { BaseService } from '../shared/service-base';
import { coreApiClient } from '../shared/api-client';
import type { ServiceResult } from '../shared/types';
import { toastStore } from '../shared/toast-store.svelte';
import type { z } from 'zod';
import type { schemas } from '$lib/genapi/core';

type Document = z.infer<typeof schemas.Document>;
type UploadFileResponse = z.infer<typeof schemas.UploadFileResponse>;

export interface FileUploadProgress {
	file: File;
	progress: number;
	status: 'pending' | 'uploading' | 'completed' | 'error';
	error?: string;
}

class DocumentsService extends BaseService {
	documents = $state<Document[]>([]);
	uploadQueue = $state<FileUploadProgress[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);

	/**
	 * Get document by ID
	 */
	async getById(id: string): Promise<ServiceResult<Document>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const document = await coreApiClient.getDocument({ params: { id } });
			return document;
		}, false);

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Get all documents in a collection
	 */
	async getByCollection(collectionId: string): Promise<ServiceResult<Document[]>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const documents = await coreApiClient.getCollectionDocuments({
				params: { id: collectionId }
			});
			this.documents = documents;
			return documents;
		}, false);

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Upload a single file to a collection
	 */
	async uploadFile(collectionId: string, file: File): Promise<ServiceResult<void>> {
		// Add to upload queue
		const queueItem: FileUploadProgress = {
			file,
			progress: 0,
			status: 'pending'
		};
		this.uploadQueue = [...this.uploadQueue, queueItem];

		const result = await this.execute(async () => {
			// Update status to uploading
			queueItem.status = 'uploading';
			this.uploadQueue = [...this.uploadQueue];

			// Get presigned upload URL
			const { uploadURL } = await coreApiClient.uploadFile({
				collectionID: collectionId,
				mimeType: file.type || 'application/octet-stream'
			});

			// Upload file to presigned URL
			const uploadResponse = await fetch(uploadURL, {
				method: 'PUT',
				headers: {
					'Content-Type': file.type || 'application/octet-stream'
				},
				body: file
			});

			if (!uploadResponse.ok) {
				throw new Error('Failed to upload file');
			}

			// Update progress to 100%
			queueItem.progress = 100;
			queueItem.status = 'completed';
			this.uploadQueue = [...this.uploadQueue];

			// Refresh documents list
			await this.getByCollection(collectionId);
		});

		if (result.error) {
			queueItem.status = 'error';
			queueItem.error = result.error.message;
			this.uploadQueue = [...this.uploadQueue];
		}

		return result;
	}

	/**
	 * Upload multiple files to a collection
	 */
	async uploadFiles(collectionId: string, files: File[]): Promise<ServiceResult<void>> {
		const result = await this.execute(async () => {
			// Upload files sequentially to maintain order
			for (const file of files) {
				const uploadResult = await this.uploadFile(collectionId, file);
				if (uploadResult.error) {
					throw new Error(`Failed to upload ${file.name}: ${uploadResult.error.message}`);
				}
			}

			toastStore.success(`Successfully uploaded ${files.length} file(s)`);
		});

		return result;
	}

	/**
	 * Clear upload queue
	 */
	clearUploadQueue() {
		this.uploadQueue = [];
	}

	/**
	 * Remove a specific file from upload queue
	 */
	removeFromQueue(file: File) {
		this.uploadQueue = this.uploadQueue.filter((item) => item.file !== file);
	}
}

export const documentsService = new DocumentsService();
