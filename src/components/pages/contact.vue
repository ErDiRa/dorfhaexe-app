<template>
	<main :class="$style.main">
		<h1>Kontakt</h1>
		<p>
			Ihr habt Wünsche und Anregungen oder wollt euch über etwas bestimmtes
			informieren?
		</p>
		<p>
			<p>
				Schreibt uns doch einfach eine Email:
			</p>
			✉️
			<a href="mailto:&#107;&#111;&#110;&#116;&#097;&#107;&#116;&#064;&#100;&#111;&#114;&#102;&#104;&#097;&#101;&#120;&#101;&#046;&#100;&#101;">
				 &#107;&#111;&#110;&#116;&#097;&#107;&#116;&#064;&#100;&#111;&#114;&#102;&#104;&#097;&#101;&#120;&#101;&#046;&#100;&#101;
			</a>
		</p>
		<p>
			<p>Oder per Post an:</p>
			<div :class="$style.contact">
				Dorfhäxe Rümmingen 1975 e.V.<br />
				Lörracher Str. 1 <br />
				79595 Rümmingen
			</div>
		</p>
	</main>
</template>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		flex-direction: column;

		.contact{
			font-weight: 600;
		}

		h1 {
			text-align: left;
		}

		p + p {
			margin-top: 16px;
		}

		p {
			margin: 0;
			font-size: 1.5rem;
			text-align: left;
			line-height: 2rem;
			> p {
				margin-bottom: 8px;
			}
			a {
				font-weight: 600;
				color: inherit;
			}
		}

		@media (max-width: 600px){
			h1{
				font-size: 1.5rem;
			}

			p{
				font-size: 1rem;
				line-height: 1.4rem;
			}
		}
	}
</style>
