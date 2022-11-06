<template>
	<h1 :class="$style.header">Dorfhäxe Rümmingen 1975 e.V.</h1>
	<!--TODO: add emblem-->
	<p :class="$style.welcome">
		Herzlich Willkommen und viel Spass auf unserer Homepage
	</p>
	<gallery></gallery>
	<div :class="$style.socialMedia">
		<h2 :class="$style.subHeader">
			Ihr wollt nichts Verpassen und up to date bleiben?
		</h2>
		<p>Dann Folgt uns einfach auf Facebook oder Instagram</p>
		<div :class="$style.icons">
			<Facebook
				:class="$style.facebook"
				@click="openLink(`https://de-de.facebook.com/dorfhaexe/`)"
			/>
			<!-- <Instagram
				:class="$style.instagram"
				@click="
					openLink(
						`https://www.instagram.com/explore/tags/geilschdicliquederwelt/`
					)
				"
			/> -->
			<div
				:class="$style.card"
				@click="
					openLink(
						`https://www.instagram.com/explore/tags/geilschdicliquederwelt/`
					)
				"
			>
				<div :class="$style.hashTag">#geilschdicliquederwelt</div>
			</div>
		</div>
	</div>
</template>

<script setup>
	import Facebook from '../../assets/facebook.svg';
	// import Instagram from '../../assets/instagram.svg';
	import Gallery from '../elements/gallery.vue';

	const openLink = (link) => {
		window.open(link, '_blank');
	};
</script>

<style lang="scss" module>
	.header {
		font-size: 3rem;
		display: flex;
		justify-content: left;
		margin-top: 1.5rem;
	}

	@media (max-width: 960px) {
		.header {
			font-size: 1.25rem;
		}
	}

	.welcome {
		text-align: left;
		font-size: 2rem;
	}

	@media (max-width: 960px) {
		.welcome {
			font-size: 1rem;
		}
	}

	.subHeader {
		font-size: 2rem;
		text-align: left;
		margin-bottom: 1rem;
	}

	@media (max-width: 960px) {
		.subHeader {
			font-size: 1rem;
		}
	}

	.socialMedia {
		margin-bottom: 2rem;
		p {
			margin-top: 0;
			text-align: left;
			font-size: 1.5rem;
		}

		@media (max-width: 960px) {
			p {
				font-size: 1rem;
			}
		}

		.icons {
			display: flex;
			margin-top: 1.5rem;

			.facebook {
				height: 3rem;
				width: 3rem;
				cursor: pointer;
			}

			.instagram {
				height: 3rem;
				width: 3rem;
				cursor: pointer;
			}

			.card {
				display: flex;
				margin-left: 2.5rem;
				padding: 0.5rem;
				border-radius: 3px;
				margin-top: auto;
				margin-bottom: auto;
				transform: rotate(-8deg);
				color: #f44336;
				border: 4px solid #f44336;
				cursor: pointer;

				.hashTag {
					font-weight: 600;
				}
			}

			svg + svg {
				margin-left: 2rem;
			}

			@media (max-width: 960px) {
				svg + svg {
					margin-left: 1.5rem;
				}
			}
		}
	}
</style>
