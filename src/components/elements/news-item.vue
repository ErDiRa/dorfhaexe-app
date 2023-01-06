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
					:srcset="`${NewsNovemberTiny} 780w,
									${NewsNovemberSmall} 1024w,
									${NewsNovemberMedium} 1440w,
									${NewsNovemberLarge} 2400w`"
					:class="$style.image"
				/>
				<a :href="pdfNov" target="_blank">
					<div :class="$style.iconContainer">
						<maximize :class="$style.maximize"></maximize>
					</div>
				</a>
			</div>
		</div>
	</div>
	<div :class="$style.container">
		<div :class="$style.content">
			<div :class="$style.text">
				<h1 :class="$style.title">Ausgabe #2 - Januar 2023</h1>
				<h1>Neues Jahr bedeutet Vollgas! 🏎️💨</h1>
				<p>Mit dem neuen Jahr starten die Fasnachts-Events</p>
			</div>
			<div :class="$style.thumbnail">
				<img
					:src="NewsNovemberTiny"
					:srcset="`${NewsJanuaryTiny} 780w,
									${NewsJanuarySmall} 1024w,
									${NewsJanuaryMedium} 1440w,
									${NewsJanuaryLarge} 2400w`"
					:class="$style.image"
				/>
				<a :href="pdfJan" target="_blank">
					<div :class="$style.iconContainer">
						<maximize :class="$style.maximize"></maximize>
					</div>
				</a>
			</div>
		</div>
	</div>
</template>

<script setup>
	import pdfNov from '../../assets/ElfterimElften.pdf';
	import pdfJan from '../../assets/january.pdf';
	import Maximize from '../../assets/maximize.svg';
	import NewsNovemberSmall from '../../assets/new-november-1024w.png';
	import NewsNovemberMedium from '../../assets/new-november-1440w.png';
	import NewsNovemberLarge from '../../assets/new-november-2400w.png';
	import NewsNovemberTiny from '../../assets/new-november-780w.png';
	import NewsJanuarySmall from '../../assets/news-january-1024w.png';
	import NewsJanuaryMedium from '../../assets/news-january-1440w.png';
	import NewsJanuaryLarge from '../../assets/news-january-2400w.png';
	import NewsJanuaryTiny from '../../assets/news-january-780w.png';
</script>

<style lang="scss" module>
	.container {
		border-bottom: 1px solid #bdbdbd;
		border-top: 1px solid #bdbdbd;
		padding: 1rem 0;

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
					line-height: 2rem;
					text-align: left;
				}

				@media (max-width: 600px) {
					h1 {
						font-size: 1rem;
						line-height: 1.3rem;
					}

					p {
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

	.container + .container {
		border-top: none;
	}

	@media (max-width: 600px) {
		.container {
			padding: 0.75rem 0.5rem;
		}
	}
</style>
