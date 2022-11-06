<template>
	<div :class="$style.footerContainer">
		<hr :class="$style.separator" />
		<div :class="$style.footer">Dorfhäxe Rümmingen 1975 e.V.</div>
	</div>
</template>

<style lang="scss" module>
	.footerContainer {
		.separator {
			border: 0;
			border-top: 2px solid #d9d9d9;
		}

		.footer {
			min-height: 4rem;
			display: flex;
			justify-self: flex-start;
			margin-top: 2rem;
		}
	}
</style>
