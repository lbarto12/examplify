<script lang="ts">
	interface Props {
		variant?: 'text' | 'circle' | 'rectangle';
		width?: string;
		height?: string;
		class?: string;
	}

	let {
		variant = 'rectangle',
		width,
		height,
		class: className = ''
	}: Props = $props();

	const variantClasses = {
		text: 'h-4 rounded',
		circle: 'rounded-full aspect-square',
		rectangle: 'rounded-xl'
	};
</script>

<div
	class={`shimmer bg-base-300 ${variantClasses[variant]} ${className}`}
	style={`${width ? `width: ${width};` : ''} ${height ? `height: ${height};` : ''}`}
></div>
