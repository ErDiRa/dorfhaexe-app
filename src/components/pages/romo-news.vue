<template>
	<main :class="$style.main">
		<h1>Rosenmontagsumzug - 12.02.2024</h1>
		<p>Hier findet ihr alle aktuellen Info zu unserem Rosenmontagsumzug</p>
		<div :class="$style.container">
			<div :class="$style.content">
				<div :class="$style.text">
					<h1 :class="$style.title">Umzugsstrecke und Aufstellungsplan 🤓</h1>
					<a :href="RomoInfo" target="_blank">
						Hier findet ihr die Umzugsstrecke sowie den aktuellen
						Aufstellungsplan zum <span>herunterladen</span>.
					</a>
				</div>
				<div :class="$style.thumbnail">
					<img :src="RomoThumb" :class="$style.image" />
					<a :href="RomoInfo" target="_blank">
						<div :class="$style.iconContainer">
							<maximize :class="$style.maximize"></maximize>
						</div>
					</a>
				</div>
			</div>
		</div>
	</main>
</template>

<script setup>
	import Maximize from '../../assets/maximize.svg';
	import RomoInfo from '../../assets/romo-2024-1.pdf';
	import RomoThumb from '../../assets/romo-thumbnail-2024-1.png';
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
		}

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
					font-size: 1.25rem;
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

					a {
						margin: 0;
						color: inherit;
						margin-top: 1rem;
						font-size: 1.5rem;
						line-height: 2rem;
						text-align: left;
						text-decoration: none;
						span {
							text-decoration: underline;
							text-underline-offset: 2px;
						}
					}

					@media (max-width: 600px) {
						h1 {
							font-size: 1rem;
							line-height: 1.3rem;
						}

						a {
							font-size: 1rem;
							line-height: 1.3rem;
							margin-top: 0;
						}
					}
				}

				.thumbnail {
					position: relative;
					margin-left: 1rem;
					cursor: pointer;

					a {
						color: inherit;
					}
					.image {
						border: 1px solid #ddd; /* Gray border */
						border-radius: 4px; /* Rounded border */
						padding: 5px; /* Some padding */
						width: 175px; /* Set a small width */
						height: auto;
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
				@media (max-width: 920px) {
					.thumbnail {
						margin-top: 1rem;
						margin-right: auto;
						margin-left: 0;
					}
				}
			}

			@media (max-width: 920px) {
				.content {
					flex-direction: column;
				}
			}
		}

		@media (max-width: 600px) {
			.container {
				padding: 1rem 0;
			}
		}

		@media (max-width: 600px) {
			h1 {
				font-size: 1.25rem;
			}

			p {
				font-size: 1rem;
				line-height: 1.3rem;
			}
		}
	}
</style>
