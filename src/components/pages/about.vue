<template>
	<main :class="$style.main">
		<person-card
			:class="$style.card"
			v-for="person in mainClub"
			:person="person"
		>
		</person-card>
	</main>
</template>

<script setup>
	import vorstandschaftJSON from '../../data/vorstandschaft.json';
	import PersonCard from '../elements/person-card.vue';

	let mainClub = JSON.parse(JSON.stringify(vorstandschaftJSON));
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		justify-content: space-between;
		flex-wrap: wrap;
		flex-direction: row;
		padding: 0 10rem;

		.card {
			width: 260px;
		}
	}

	@media (max-width: 650px) {
		.main {
			display: block;
			padding: 0;
			.card {
				width: 260px;
				margin-bottom: 0.5rem;
				margin-left: auto;
				margin-right: auto;
			}
		}
	}
</style>
