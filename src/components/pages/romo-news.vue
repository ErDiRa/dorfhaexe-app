<template>
	<main :class="$style.main">
		<h1>50 Johr Dorfhäxe Rümminge - 18.01.2025</h1>
		<div :class="$style.container">
			<div :class="$style.content">
				<div :class="$style.text">
					<h1 :class="$style.title">50 Johr 🥳🥳, verruckt!</h1>
					<div>
						<div :class="$style.infos">
							Gerne würden wir das mit euch feiern!!
							<div>
								<ul>
									<li>
										<div :class="$style.row">
											<div :class="$style.key"><b>Wann?&nbsp;</b></div>
											<div :class="$style.value">Am 18.01.2025 ab 19.30Uhr</div>
										</div>
									</li>
									<li>
										<div :class="$style.row">
											<div :class="$style.key"><b>Wo?&nbsp;</b></div>
											<div :class="$style.value">Gemeindehalle Rümmingen</div>
										</div>
									</li>
									<li>
										<div :class="$style.row">
											<div :class="$style.key"><b>Eintritt?&nbsp;</b></div>
											<div :class="$style.value">10€</div>
										</div>
									</li>
									<li>
										<div :class="$style.row">
											<div :class="$style.key"><b>Anmeldung?&nbsp;</b></div>
											<div :class="$style.value">
												<a
													:class="$style.link"
													href="mailto:jubidorfhaexe@gmx.de"
													>jubidorfhaexe@gmx.de</a
												>
											</div>
										</div>
									</li>
								</ul>
							</div>
						</div>
						<a :href="RomoInfo" target="_blank" :class="$style.link"
							>Noch nicht überzeugt? Noch mehr Infos findet ihr in unserem Flyer
							👇👇</a
						>
					</div>
				</div>
				<div :class="$style.thumbnail">
					<a :href="RomoInfo" target="_blank">
						<img :src="RomoThumb" :class="$style.image" />
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
	import RomoInfo from '../../assets/jubi.pdf';
	import RomoThumb from '../../assets/jubi.png';
	import Maximize from '../../assets/maximize.svg';
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
				flex-direction: column;
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

					div {
						margin: 0;
						color: inherit;
						font-size: 1.5rem;
						line-height: 2rem;
						text-align: left;
						a {
							margin: 0;
							color: inherit;
							margin-top: 1rem;
							font-size: 1.5rem;
							line-height: 2rem;
							text-align: left;
							text-decoration: none;
							position: relative;
						}

						.infos {
							display: flex;
							flex-direction: column;
						}

						.footer {
							display: flex;
							flex-direction: row;
						}
						span {
							text-decoration: underline;
							text-underline-offset: 2px;
						}

						ul {
							list-style-type: none;
							padding: 0;

							li {
								.row {
									display: flex;
									flex-direction: column;
									.value {
										a {
											text-decoration: underline;
											text-underline-offset: 2px;
										}
									}
								}
							}

							li + li {
								margin-top: 1rem;
							}
						}
					}

					@media (max-width: 600px) {
						h1 {
							font-size: 1rem;
							line-height: 1.3rem;
						}

						div {
							font-size: 1rem;
							line-height: 1.3rem;
							margin-top: 0;
						}

						.link {
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
						width: 80vw; /* Set a small width */
						max-width: 512px;
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
