<template>
	<main :class="$style.main">
		<h1>Geschichtsstunde</h1>
		<div>
			<p>
				Man munkelt, dass es in jedem Dorf Hexen gibt. Auch in Rümmingen soll
				dies der Fall sein. In der Fasnachtszeit tauchen sie immer wieder auf
				und verschwinden dann wieder im Dunkeln der Nacht.
			</p>
			<p>
				Dementsprechend schlossen sich im Jahr 1975 ein paar Jugendliche
				zusammen und gründeten eine Hexenclique. Es war kein leichter Start für
				die jungen Leute, denn es fehlte das erforderliche Geld. Motivation war
				genügend vorhanden, aus diesem Grund wurden Utensilien wie Ringelhemd
				und Strümpfe für das Häs im Ort, im „Lade vo de Maria“ besorgt und der
				Stoff für Rock und Schurz in der KBC geholt und alles selbst finanziert.
				Hilfe beim Nähen bekam man dann günstig im Ort von Frau Kapp. Die ersten
				Masken besorgten Helga und Manfred Gempp von einem Larvenmacher aus
				Basel. Sie wurden aus Pappmasche gefertigt. Zähne und Haare stammten
				allerdings von Schweinen und Rindern. Strohschuhe waren zu diesem
				Zeitpunkt schlicht zu teuer, weshalb man alles Verfügbare an Schuhwerk
				auftrug.
			</p>
			<p>
				Der Wagenbau war eines der schönsten Dinge dieser Zeit. Mit einfachen
				Mitteln galt es einen Umzugswagen zu bauen, welcher Jahr für Jahr das
				Motto der Kampagnen ausspielte und mit den „Großen“ konkurrieren sollte.
				Durch immer größer werdende Auflagen und strengere Vorschriften wurde
				diese Tradition leider nicht mehr fortgeführt. Auch die fehlenden
				Möglichkeiten den Wagen unter zustellen machten einen Strich durch die
				Rechnung.
			</p>
			<p>
				Im Jahr 1985 folgte dann der Eintrag in das Vereinsregister des
				Amtsgerichts. Aus den Dorfhexe Rümmingen wurden die heutigen Dorfhäxe
				Rümmingen 1975 e.V. Auch wurde man in Sachen Häs etwas professioneller.
				Hochwertige Holzmasken folgten den Pappmasche-Masken und auch der Rest
				des Häs wurde modernisiert. Heute besteht es aus Strohschuhen, roten
				Socken, einer langen weißen Unterhose mit Rüschensaum und Glöckchen,
				einem grünen Rock und darüber einer roten Schürze. Auch das Oberteil aus
				einem Ringelhemd musste weichen. Dies wurde erst durch eine Rot-Weiß
				karierte und später dann durch die aktuelle Gelbe Bluse ersetzt, auf der
				auf der Brust das Emblem der Dorfhäxe ragt. Die Schultern werden von
				einem roten Dreiecktuch bedeckt und rote Handschuhe bedecken auch den
				Rest der freien Haut. In diesem Dress gehen die Dorfhäxe Rümmingen an
				Fasnacht auf die Straßen und machen diese unsicher!
			</p>
			<p>
				Mit einem bunten Mix aus „alten Hasen“ und „jungen Füchsen“ bleiben die
				Dorfhäxen Rümmingen 1975 e.V. auch weiterhin mit sehr viel Spaß, Freude
				und Motivation bei ihren Aktivitäten am Start.
			</p>
		</div>
		<h1>Vereinsaktivitäten</h1>
		<div>
			<p>
				Nicht nur an Fasnacht sind die Dorfhäxen unterwegs. Mit der
				Generalversammlung im Oktober startet jedes Jahr eine neue Kampagne.
				Damit beginnen auch Planung und Organisation für die Fasnacht.
			</p>
			<p>
				Am 11.11. starten die Häxen mit Cliquen aus der Nachbarschaft in die
				neue Fasnachstssaison. Außerdem wird das Fasnachtsmotto für die
				jeweilige Kampagne an diesem Tag bekannt gegeben. Nach den „Heiligen
				Drei Königen“ geht es dann so richtig los und die Zeit der Narrentreffen
				beginnt. In der näheren Umgebung, aber auch darüber hinaus, nehmen die
				Häxen an verschiedensten Narrentreffen, Umzügen und
				Hallenveranstaltungen teil.
			</p>

			<p>
				Mit dem „Schmutzige Dunnschdig" beginnt die Hauptfasnacht. Hier wird der
				Bürgermeister entmachtet und die Häxen regieren im Dorf.
			</p>

			<p>
				Vier Tage später am Rosenmontag, finden zahlreiche Cliquen,
				Guggenmusiken und sonstige Vereine und Leute den Weg nach Rümmingen,
				denn hier findet der traditionelle Rosenmontagsumzug statt. Auch gibt es
				im kompletten Dorf verteilt diverse Bwirtungsstände verschiedener
				Vereine und Organisationen, sowie ein großes Partyzelt auf dem Schulhof
				in dem bis in die Nacht zusammen gefeiert wird.
			</p>

			<p>
				Am Buurefasnachtssonntag heißt es dann schon wieder Abschied nehmen von
				der Fasnacht. Symbolisch wird hier die „Häx“ verbannt und somit das Ende
				der Fasnacht dargestellt.
			</p>

			<p>
				Das Ende der Fasnacht ist jedoch nicht das Ende der Aktivitäten der
				Dorfhäxe. Teilnahmen an Grümpeltunier, Maiwanderung, Dorfhock,
				Dreikampfturnier Weihnachtsbaumsammlung und vielem mehr gestalten das
				Jahr.
			</p>
			<p>
				So eine Kampagne kann ganz schön anstrengend sein, dennoch haben die
				Häxen immer sehr viel Spaß und Freude dabei! Wenn Ihr an einer
				Mitgliedschaft der Dorfhäxe Rümmingen 1975 e.V. interessiert seid,
				kontaktiert uns doch einfach über unsere
				<router-link to="/kontakt">Kontakt Möglichkeiten.</router-link>

			  <div :class="$style.bye">
					Eure Dorfhäxe Rümmingen 1975 e.V.
					<haexe-logo></haexe-logo>
				</div>
			</p>
		</div>
	</main>
</template>

<script setup>
	import HaexeLogo from '../../assets/dorfhaexe-emblem.svg';
</script>

<style lang="scss" module>
	.main {
		min-height: 100vh;
		display: flex;
		flex-direction: column;

		.contact {
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

		@media (max-width: 600px) {
			h1 {
				font-size: 1.5rem;
			}

			p {
				font-size: 1rem;
				line-height: 1.4rem;
			}
		}

		.bye{
			margin-top: 16px;
			display: flex;
			flex-direction: column;
			align-items: flex-start;
			font-weight: 600;

			svg {
				margin-top: 8px;
			}
		}
	}
</style>
