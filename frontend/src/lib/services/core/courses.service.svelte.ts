/**
 * Courses service
 */

import { BaseService } from '../shared/service-base';
import { coreApiClient } from '../shared/api-client';
import type { ServiceResult } from '../shared/types';

class CoursesService extends BaseService {
	courses = $state<string[]>([]);
	loading = $state(false);
	error = $state<string | null>(null);

	/**
	 * Get all courses
	 */
	async getCourses(): Promise<ServiceResult<string[]>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const courses = await coreApiClient.getCourses();
			// Handle null response from backend by defaulting to empty array
			const courseList = courses ?? [];
			this.courses = courseList;
			return courseList;
		});

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}

	/**
	 * Create a new course
	 */
	async createCourse(name: string): Promise<ServiceResult<string>> {
		this.loading = true;
		this.error = null;

		const result = await this.execute(async () => {
			const response = await coreApiClient.newCourse({ name });
			// Refresh the courses list
			await this.getCourses();
			return response.courseName;
		});

		this.loading = false;
		if (result.error) {
			this.error = result.error.message;
		}

		return result;
	}
}

export const coursesService = new CoursesService();
