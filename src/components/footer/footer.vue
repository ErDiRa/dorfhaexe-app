<template>
	<div :class="$style.footerContainer">
		<hr :class="$style.separator" />
		<div :class="$style.footer">Dorfhäxe Rümmingen 1975 e.V.</div>
		<div :class="$style.links">
			<!--TODO: add views and links via router-->
			<a>Impressum</a>
			<a>Datenschutz</a>
		</div>
	</div>
</template>

<style lang="scss" module>
	.footerContainer {
		.separator {
			border: 0;
			border-top: 2px solid #d9d9d9;
		}

		.footer {
			min-height: 3rem;
			display: flex;
			justify-self: flex-start;
			margin-top: 2rem;
		}

		.links {
			display: flex;
			flex-direction: row;

			a {
				border-bottom: 1px solid black;
				padding-bottom: 0.1rem;
			}

			a + a {
				margin-left: 1rem;
			}
		}
	}
</style>
