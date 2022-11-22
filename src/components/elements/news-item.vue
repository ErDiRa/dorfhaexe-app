<template>
	<div :class="$style.container">
		<div :class="$style.content">
			<div :class="$style.text">
				<h1 :class="$style.title">Ausgabe #1 - November 2022</h1>
				<h1>Endlich widda Fasnacht <br />Narri Narro 🎉🎉</h1>
				<p>Nach zwei Jahren ohne geht’s heute wieder los</p>
			</div>
			<div :class="$style.thumbnail">
				<img
					:src="NewsNovemberTiny"
					:srcset="`${NewsNovemberSmall} 780w,
									${NewsNovemberSmall} 1024w,
									${NewsNovemberMedium} 1440w,
									${NewsNovemberLarge} 2400w`"
					:class="$style.image"
				/>
				<a :href="src" target="_blank">
					<div :class="$style.iconContainer">
						<maximize :class="$style.maximize"></maximize>
					</div>
				</a>
			</div>
		</div>
	</div>
</template>

<script setup>
	import { ref } from 'vue';
	import Maximize from '../../assets/maximize.svg';
	import NewsNovemberSmall from '../../assets/new-november-1024w.png';
	import NewsNovemberMedium from '../../assets/new-november-1440w.png';
	import NewsNovemberLarge from '../../assets/new-november-2400w.png';
	import NewsNovemberTiny from '../../assets/new-november-780w.png';

	// TODO: use a blob to lazy load pdf to prevent loading per default from assets
	// TODO: checkout how hosting on ionos work to provide resources
	// TODO: add base url in href
	const src = ref(
		new URL(
			'../../dorfhaexe-app/assets/' + 'ElfterimElften.pdf',
			import.meta.url
		)
	);
</script>

<style lang="scss" module>
	.container {
		border-bottom: 1px solid #bdbdbd;
		border-top: 1px solid #bdbdbd;
		padding: 0.5rem;

		.content {
			display: flex;
			flex-direction: row;
			.title {
				text-align: left;
				margin-bottom: 1rem;
			}

			.text {
				display: flex;
				margin-right: auto;
				flex-direction: column;
				.title {
					text-align: left;
					margin-bottom: 1rem;
				}

				h1 {
					margin: 0;
					text-align: left;
					font-size: 2rem;
					line-height: 2.3rem;
				}

				p {
					margin: 0;
					margin-top: 1rem;
					line-height: 1.3rem;
					text-align: left;
				}

				@media (max-width: 600px) {
					h1 {
						font-size: 1rem;
						line-height: 1.3rem;
					}
				}
			}

			.thumbnail {
				position: relative;

				a {
					color: inherit;
				}
				.image {
					border: 1px solid #ddd; /* Gray border */
					border-radius: 4px; /* Rounded border */
					padding: 5px; /* Some padding */
					width: 150px; /* Set a small width */
				}

				.iconContainer {
					position: absolute;
					width: 32px;
					height: 32px;
					bottom: 8px;
					right: 6px;
					background-color: #443c3c83;
					border-radius: 0;
					display: flex;
					border-radius: 5px;
					cursor: pointer;

					.maximize {
						margin: auto;
						stroke: white;
					}
				}
			}
			@media (max-width: 600px) {
				.thumbnail {
					margin-top: 1rem;
					margin-right: auto;
				}
			}
		}

		@media (max-width: 600px) {
			.content {
				flex-direction: column;
			}
		}
	}
</style>
