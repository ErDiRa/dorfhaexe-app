<template>
	<main :class="$style.main">
		<h1>#Häxe-News</h1>
		<p>
			Kein Facebook und kein Instagram ?! Kein Problem unsere aktuellsten News
			erhaltet ihr auch hier. 🤓
		</p>
		<news-item></news-item>
	</main>
</template>

<script setup>
	import NewsItem from '../elements/news-item.vue';
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		flex-direction: column;

		h1 {
			margin-bottom: 0;
			text-align: left;
			font-size: 2.5rem;
		}

		p {
			font-size: 2rem;
			line-height: 2.3rem;
		}

		@media (max-width: 600px) {
			h1 {
				text-align: left;
				font-size: 1.5rem;
			}

			p {
				text-align: left;
				font-size: 1rem;
				line-height: 1.3rem;
			}
		}
	}
</style>
