<template>
	<main>
		<h1 :class="$style.header">Dorfhäxe Rümmingen 1975 e.V.</h1>
		<!--TODO: add emblem-->
		<p :class="$style.welcome">
			Herzlich Willkommen und viel Spass auf unserer Homepage
		</p>
		<gallery></gallery>
		<div :class="$style.newsContainer">
			<hr :class="$style.separator" />
			<div>
				<h2 :class="$style.subHeader">
					Der Rosenmontag steht vor der Tür ☝️☝️
				</h2>
				<p @click="navigateTo(navigation.romo.ROUTE)">
					<a>Hier</a> findet ihr aktuelle Infos für unsere Rosenmontagsumzug am
					<b>12.02.2024</b>
				</p>
			</div>
			<hr :class="$style.separator" />
		</div>

		<div :class="$style.socialMedia">
			<h2 :class="$style.subHeader">
				Ihr wollt nichts verpassen und up to date bleiben?
			</h2>
			<p>Dann folgt uns einfach auf Facebook oder Instagram</p>
			<div :class="$style.icons">
				<Facebook
					:class="$style.facebook"
					@click="openLink(`https://de-de.facebook.com/dorfhaexe/`)"
				/>

				<div :class="$style.stamp">
					<div :class="$style.card">#geilschdicliquederwelt</div>

					<div
						:class="[$style.card, $style.rubber]"
						:style="`background-image: url(${rubberImg})`"
						@click="openLink(`https://instagram.com/dorfhaexe1975`)"
					></div>
				</div>
			</div>
		</div>
	</main>
</template>

<script setup>
	import { useRouter } from 'vue-router';
	import Facebook from '../../assets/facebook.svg';
	import rubberImg from '../../assets/rubber.png';
	import { navigation } from '../../const/strings';
	import Gallery from '../elements/gallery.vue';

	const router = useRouter();

	const openLink = (link) => {
		window.open(link, '_blank');
	};

	const navigateTo = (path) => {
		router.push(path);
	};
</script>

<style lang="scss" module>
	.header {
		font-size: 3rem;
		display: flex;
		justify-content: left;
		margin-top: 1.5rem;
		margin-bottom: 0;
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

	.separator {
		border: 0;
		border-top: 2px solid #d9d9d9;
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

	.newsContainer {
		margin-top: 0.5rem;
		margin-bottom: 0.5rem;

		div {
			padding: 0.5rem 0;
		}

		p {
			margin-top: 0;
			text-align: left;
			font-size: 1.5rem;

			a {
				cursor: pointer;
				text-decoration: underline;
				text-underline-offset: 2px;
			}
		}
		@media (max-width: 960px) {
			p {
				font-size: 1rem;
			}
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

			.stamp {
				position: relative;
				.card {
					display: flex;
					margin-left: 2.5rem;
					padding: 0.5rem;
					border-radius: 3px;
					transform: rotate(-8deg);
					color: #f44336;
					border: 4px solid #f44336;
					cursor: pointer;
					position: absolute;
					top: 0;
					left: 0;
					font-weight: 600;
				}
				.rubber {
					border: none;
					width: 205px;
					height: 42px;
					z-index: 7;
					opacity: 0.5;
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
