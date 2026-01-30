<script lang="ts">
	import katex from 'katex';

	let { content }: { content: string } = $props();

	// Regex patterns for explicit LaTeX delimiters
	const DISPLAY_MATH_REGEX = /\$\$([\s\S]*?)\$\$|\\\[([\s\S]*?)\\\]/g;
	const INLINE_MATH_REGEX = /\$((?!\$)[\s\S]*?)\$|\\\(([\s\S]*?)\\\)/g;

	// Map text representations to LaTeX
	const LATEX_REPLACEMENTS: Record<string, string> = {
		'alpha': '\\alpha', 'beta': '\\beta', 'gamma': '\\gamma', 'delta': '\\delta',
		'epsilon': '\\epsilon', 'zeta': '\\zeta', 'eta': '\\eta', 'theta': '\\theta',
		'iota': '\\iota', 'kappa': '\\kappa', 'lambda': '\\lambda', 'mu': '\\mu',
		'nu': '\\nu', 'xi': '\\xi', 'omicron': 'o', 'pi': '\\pi',
		'rho': '\\rho', 'sigma': '\\sigma', 'tau': '\\tau', 'upsilon': '\\upsilon',
		'phi': '\\phi', 'chi': '\\chi', 'psi': '\\psi', 'omega': '\\omega',
		'Alpha': 'A', 'Beta': 'B', 'Gamma': '\\Gamma', 'Delta': '\\Delta',
		'Epsilon': 'E', 'Zeta': 'Z', 'Eta': 'H', 'Theta': '\\Theta',
		'Iota': 'I', 'Kappa': 'K', 'Lambda': '\\Lambda', 'Mu': 'M',
		'Nu': 'N', 'Xi': '\\Xi', 'Omicron': 'O', 'Pi': '\\Pi',
		'Rho': 'P', 'Sigma': '\\Sigma', 'Tau': 'T', 'Upsilon': '\\Upsilon',
		'Phi': '\\Phi', 'Chi': 'X', 'Psi': '\\Psi', 'Omega': '\\Omega',
		'infinity': '\\infty', 'infty': '\\infty',
		'sum': '\\sum', 'prod': '\\prod', 'int': '\\int', 'integral': '\\int',
		'lim': '\\lim',
		'>=': '\\geq', '<=': '\\leq', '!=': '\\neq', '≥': '\\geq', '≤': '\\leq', '≠': '\\neq',
		'±': '\\pm', '×': '\\times', '·': '\\cdot',
	};

	function convertToLatex(expr: string): string {
		let result = expr;

		// Replace Greek letters and symbols
		for (const [text, latex] of Object.entries(LATEX_REPLACEMENTS)) {
			const regex = new RegExp(`\\b${text}\\b`, 'g');
			result = result.replace(regex, latex);
		}

		// Convert sqrt(x) to \sqrt{x}
		result = result.replace(/sqrt\(([^)]+)\)/gi, '\\sqrt{$1}');
		result = result.replace(/√(\w+)/g, '\\sqrt{$1}');

		// Convert fractions a/b to \frac{a}{b} for simple cases
		result = result.replace(/(\d+|\w+)\s*\/\s*(\d+|\w+)/g, '\\frac{$1}{$2}');

		return result;
	}

	function autoDetectMath(text: string): string {
		if (!text) return '';

		// First check if it already has LaTeX delimiters (use non-capturing test)
		const hasDisplayMath = /\$\$([\s\S]*?)\$\$|\\\[([\s\S]*?)\\\]/.test(text);
		const hasInlineMath = /\$((?!\$)[\s\S]*?)\$|\\\(([\s\S]*?)\\\)/.test(text);

		if (hasDisplayMath || hasInlineMath) {
			return text;
		}

		let result = text;

		// Only auto-detect clear mathematical patterns
		// Exponents like x^2, |E|^2, n^{k+1} - any character followed by ^ and number/expression
		result = result.replace(/(\|?\w+\|?)\^(\{[^}]+\}|\d+)/g, (match, base, exp) => {
			const cleanExp = exp.startsWith('{') ? exp : `{${exp}}`;
			return `$${base}^${cleanExp}$`;
		});

		// Subscripts like x_1, x_{ij} (must be letter followed by underscore and number)
		result = result.replace(/\b([a-zA-Z])_(\{[^}]+\}|\d+)/g, (match, base, sub) => {
			const cleanSub = sub.startsWith('{') ? sub : `{${sub}}`;
			return `$${base}_${cleanSub}$`;
		});

		// Square roots - explicit sqrt() or √
		result = result.replace(/sqrt\(([^)]+)\)/gi, (match, inner) => {
			return `$\\sqrt{${inner}}$`;
		});
		result = result.replace(/√(\w+|\d+)/g, (match, inner) => {
			return `$\\sqrt{${inner}}$`;
		});

		// Greek letters (standalone words only)
		const greekPattern = /\b(alpha|beta|gamma|delta|epsilon|zeta|eta|theta|iota|kappa|lambda|mu|nu|xi|pi|rho|sigma|tau|upsilon|phi|chi|psi|omega)\b/gi;
		result = result.replace(greekPattern, (match) => {
			const latex = LATEX_REPLACEMENTS[match.toLowerCase()];
			return latex ? `$${latex}$` : match;
		});

		// Math functions like sin(x), cos(x), log(x) - must have parentheses
		result = result.replace(/\b(sin|cos|tan|cot|sec|csc|arcsin|arccos|arctan|sinh|cosh|tanh|log|ln|exp)\(([^)]+)\)/gi, (match, func, arg) => {
			return `$\\${func.toLowerCase()}(${arg})$`;
		});

		// Fractions - numeric or simple variable fractions like 1/2, n/2, n/m
		result = result.replace(/\b(\d+|[a-zA-Z])\s*\/\s*(\d+|[a-zA-Z])\b/g, (match, num, denom) => {
			return `$\\frac{${num}}{${denom}}$`;
		});

		// Merge adjacent $ delimiters
		result = result.replace(/\$\s*\$/g, ' ');

		return result;
	}

	function renderLatex(text: string): string {
		if (!text) return '';

		// First, auto-detect and wrap mathematical expressions
		let result = autoDetectMath(text);

		// Handle display math ($$...$$ or \[...\])
		result = result.replace(DISPLAY_MATH_REGEX, (match, p1, p2) => {
			const latex = p1 || p2;
			try {
				return katex.renderToString(latex.trim(), {
					displayMode: true,
					throwOnError: false,
					trust: true
				});
			} catch (e) {
				console.warn('KaTeX display error:', e);
				return match;
			}
		});

		// Handle inline math ($...$ or \(...\))
		result = result.replace(INLINE_MATH_REGEX, (match, p1, p2) => {
			const latex = p1 || p2;
			try {
				return katex.renderToString(latex.trim(), {
					displayMode: false,
					throwOnError: false,
					trust: true
				});
			} catch (e) {
				console.warn('KaTeX inline error:', e);
				return match;
			}
		});

		return result;
	}

	let renderedContent = $derived(renderLatex(content));
</script>

<span class="latex-content">{@html renderedContent}</span>

<style>
	.latex-content {
		display: inline;
	}

	.latex-content :global(.katex) {
		font-size: 1.1em;
	}

	.latex-content :global(.katex-display) {
		margin: 1em 0;
		overflow-x: auto;
		overflow-y: hidden;
	}
</style>
