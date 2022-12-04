<template>
	<main :class="$style.main">
		<h1>#Häxe-News</h1>
		<p>
			Kein
			<a href="https://de-de.facebook.com/dorfhaexe/">Facebook</a>? Kein
			<a href="https://instagram.com/dorfhaexe1975?igshid=YmMyMTA2M2Y="
				>Instagram</a
			>? <br />Kein Problem unsere aktuellsten News erhaltet ihr auch hier. 🤓
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
			text-align: left;
			font-size: 2rem;
			line-height: 2.8rem;
			a {
				color: inherit;
			}
		}

		@media (max-width: 600px) {
			h1 {
				font-size: 1.5rem;
			}

			p {
				font-size: 1rem;
				line-height: 1.3rem;
			}
		}
	}
</style>
