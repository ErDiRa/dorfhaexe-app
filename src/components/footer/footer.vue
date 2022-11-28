<template>
	<div :class="$style.footerContainer">
		<hr :class="$style.separator" />
		<div :class="$style.footer">Dorfhäxe Rümmingen 1975 e.V.</div>
		<div :class="$style.links">
			<!--TODO: add views and links via router-->
			<a @click="showImpressum">Impressum</a>
			<a>Datenschutz</a>
		</div>

		<modal
			:is-open="impressumIsOpen"
			@close="impressumIsOpen = false"
			:class="$style.modal"
		>
			<template #title>Impressum</template>
			<template #content>
				<div :class="$style.content">
					<p>
						<strong>Verantwortlich: </strong><br />
						Dorfhäxe Rümmingen 1975 e.V.<br />
						Marco Pollini <br />
						Lörracher Str. 1 79595 Rümmingen <br />
						Kontakt: E-Mail: kontakt@dorfhaexe.de
					</p>

					<p>
						<strong>Rechtliche Angaben</strong><br />
						Vertretungsberechtigter: Marco Pollini<br />
						Registergericht: Amtsgericht Lörrach Registernummer: VR 793
					</p>
					<p>
						<strong>Weitere Infos <br /></strong>
						Trotz sorgfältiger inhaltlicher Kontrolle übernehmen wir keine
						Haftung für die Inhalte externer Links. Für den Inhalt der
						verlinkten Seiten sind ausschließlich deren Betreiber
						verantwortlich.
					</p>
					<p>
						Alle hier verwendeten Namen, Begriffe, Zeichen und Grafiken können
						Marken- oder Warenzeichen im Besitze ihrer rechtlichen Eigentümer
						sein. Die Rechte aller erwähnten und benutzten Marken- und
						Warenzeichen liegen ausschließlich bei deren Besitzern.
					</p>

					<p>
						<strong
							>Allgemeine Nutzungsbedingungen für Dorfhäxe Rümmingen 1975
							e.V.</strong
						>
					</p>

					<p>
						<strong>1. Informationen zum Urheberrecht</strong><br />
						Alle Informationen dieser Web-Seite werden wie angegeben ohne
						Anspruch auf Richtigkeit, Vollständigkeit oder Aktualität zur
						Verfügung gestellt. Wenn nicht ausdrücklich anderweitig in dieser
						Publikation zu verstehen gegeben, und zwar in Zusammenhang mit einem
						bestimmten Ausschnitt, einer Datei, oder einem Dokument, ist
						jedermann dazu berechtigt, dieses Dokument anzusehen, zu kopieren,
						zu drucken und zu verteilen, unter den folgenden Bedingungen: Das
						Dokument darf nur für nichtkommerzielle Informationszwecke genutzt
						werden. Jede Kopie dieses Dokuments oder eines Teils davon muss
						diese urheberrechtliche Erklärung und das urheberrechtliche
						Schutzzeichen des Betreibers enthalten. Das Dokument, jede Kopie des
						Dokuments oder eines Teils davon dürfen nicht ohne schriftliche
						Zustimmung des Betreibers verändert werden. Der Betreiber behält
						sich das Recht vor, diese Genehmigung jederzeit zu widerrufen, und
						jede Nutzung muss sofort eingestellt werden, sobald eine
						schriftliche Bekanntmachung seitens des Betreibers veröffentlicht
						wird.
					</p>

					<p>
						<strong
							>2. Vertragliche Zusicherungen und Verzichterklärungen</strong
						><br />
						Die Website Dorfhäxe Rümmingen 1975 e.V. steht Ihnen - soweit nicht
						anders vereinbart - kostenlos zur Verfügung. Die Betreiber
						übernehmen keinerlei Gewähr für Richtigkeit der enthaltenen
						Informationen, Verfügbarkeit der Dienste, Verlust von auf Dorfhäxe
						Rümmingen 1975 e.V. abgespeicherten Daten oder Nutzbarkeit für
						irgendeinen bestimmten Zweck. Die Betreiber haften auch nicht für
						Folgeschäden, die auf einer Nutzung des Angebotes beruhen. Soweit
						ein Haftungsausschluss nicht in Betracht kommt, haften die Betreiber
						lediglich für grobe Fahrlässigkeit und Vorsatz. Produkt- und
						Firmennamen sind Marken der jeweiligen Eigentümer und werden auf
						diesen Seiten ausschließlich zu Informationszwecken eingesetzt.
						Diese Publikation könnte technische oder andere Ungenauigkeiten
						enthalten oder Schreib-/Tippfehler. Von Zeit zu Zeit werden der
						vorliegenden Information Änderungen hinzugefügt; diese Änderungen
						werden in neuen Ausgaben der Publikation eingefügt. Der Betreiber
						kann jederzeit Verbesserungen und/oder Veränderungen an den
						Angeboten vornehmen, die in dieser Publikation beschrieben werden.
					</p>

					<p>
						<strong>3. Hinweis gemäß Teledienstgesetz</strong> <br />
						Für Internetseiten Dritter, auf die die dieses Angebot durch sog.
						Links verweist, tragen die jeweiligen Anbieter die Verantwortung.
						Der Betreiber ist für den Inhalt solcher Seiten Dritter nicht
						verantwortlich. Des Weiteren kann die Web-Seite ohne unser Wissen
						von anderen Seiten mittels sog. Links angelinkt werden. Der
						Betreiber übernimmt keine Verantwortung für Darstellungen, Inhalt
						oder irgendeine Verbindung zu dieser Web-Seite in Web-Seiten
						Dritter. Für fremde Inhalte ist der Betreiber nur dann
						verantwortlich, wenn von ihnen (d.h. auch von einem rechtswidrigen
						oder strafbaren Inhalt) positive Kenntnis vorliegt und es technisch
						möglich und zumutbar ist, deren Nutzung zu verhindern. Der Betreiber
						ist nach dem Teledienstgesetz jedoch nicht verpflichtet, die fremden
						Inhalte ständig zu überprüfen.
					</p>

					<p>
						<strong>Kontakt</strong> <br />
						Fragen rund um Dorfhäxe Rümmingen 1975 e.V. bitte an den Webmaster
						richten.
					</p>

					<p>
						<strong>Rechtswirksamkeit</strong> <br />
						Diese Allgemeinen Nutzungsbedingungen beziehen sich auf Dorfhäxe
						Rümmingen 1975 e.V.. Sofern Teile oder einzelne Formulierungen
						dieses Textes der geltenden Rechtslage nicht, nicht mehr oder nicht
						vollständig entsprechen sollten, bleiben die übrigen Teile des
						Dokumentes in ihrem Inhalt und ihrer Gültigkeit davon unberührt.﻿
					</p>
				</div>
			</template>
		</modal>
	</div>
</template>

<script setup>
	import { ref } from 'vue';
	import Modal from '../elements/modal.vue';

	let impressumIsOpen = ref(false);
	const showImpressum = () => {
		console.log('show impressum');
		impressumIsOpen.value = true;
	};
</script>

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
