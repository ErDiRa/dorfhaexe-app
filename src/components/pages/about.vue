<template>
	<main :class="$style.main">
		<h1>Vorstandschaft</h1>
		<p>Hier erhaltet ihr einen Überblick über die aktuelle Vorstandschaft</p>
		<header-collapsable
			:title="'Hauptverein'"
			:init-is-open="true"
			:class="$style.collapsable"
		>
			<div :class="$style.content">
				<person-card
					:class="$style.card"
					v-for="person in board.main"
					:person="person"
				>
				</person-card>
			</div>
		</header-collapsable>
		<header-collapsable
			:title="'Förderverein'"
			:init-is-open="true"
			:class="[$style.collapsable, $style.footerGap]"
		>
			<div :class="$style.content">
				<person-card
					:class="$style.card"
					v-for="person in board.support"
					:person="person"
				>
				</person-card>
			</div>
		</header-collapsable>
	</main>
</template>

<script setup>
	import boardJSON from '../../data/vorstandschaft.json';
	import HeaderCollapsable from '../elements/header-collapsable.vue';
	import PersonCard from '../elements/person-card.vue';

	let board = JSON.parse(JSON.stringify(boardJSON));
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		h1 {
			text-align: left;
		}

		p {
			text-align: left;
			font-size: 1.5rem;
			margin-top: 0;
		}

		@media (max-width: 600px) {
			h1 {
				font-size: 1rem;
			}

			p {
				font-size: 1rem;
			}
		}

		.collapsable {
			.content {
				display: flex;
				justify-content: space-between;
				flex-wrap: wrap;
				flex-direction: row;
				padding: 0 10rem;
				margin-top: 1rem;

				.card {
					width: 260px;
				}
			}

			@media (max-width: 950px) {
				.content {
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
		}

		.collapsable + .collapsable {
			margin-top: 1rem;
		}
	}

	.footerGap {
		margin-bottom: 4rem;
	}
</style>
