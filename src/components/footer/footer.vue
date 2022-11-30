<template>
	<div :class="$style.footerContainer">
		<hr :class="$style.separator" />
		<div :class="$style.footer">Dorfhäxe Rümmingen 1975 e.V.</div>
		<div :class="$style.links">
			<!--TODO: add views and links via router-->
			<a @click="showImpressum">Impressum</a>
			<a @click="showPrivacy">Datenschutz</a>
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

		<modal
			:is-open="privacyIsOpen"
			@close="privacyIsOpen = false"
			:class="$style.modal"
		>
			<template #title>Datenschutzerklärung</template>
			<template #content>
				<div :class="$style.content">
					<h2 id="m14">Einleitung</h2>
					<p>
						Mit der folgenden Datenschutzerklärung möchten wir Sie darüber
						aufklären, welche Arten Ihrer personenbezogenen Daten (nachfolgend
						auch kurz als "Daten“ bezeichnet) wir zu welchen Zwecken und in
						welchem Umfang verarbeiten. Die Datenschutzerklärung gilt für alle
						von uns durchgeführten Verarbeitungen personenbezogener Daten,
						sowohl im Rahmen der Erbringung unserer Leistungen als auch
						insbesondere auf unseren Webseiten, in mobilen Applikationen sowie
						innerhalb externer Onlinepräsenzen, wie z.B. unserer
						Social-Media-Profile (nachfolgend zusammenfassend bezeichnet als
						"Onlineangebot“).
					</p>
					<p>Die verwendeten Begriffe sind nicht geschlechtsspezifisch.</p>
					<p>Stand: 30. November 2022</p>
					<h2>Inhaltsübersicht</h2>
					<ul class="index">
						<li><a class="index-link" href="#m14">Einleitung</a></li>
						<li><a class="index-link" href="#m3">Verantwortlicher</a></li>
						<li>
							<a class="index-link" href="#mOverview"
								>Übersicht der Verarbeitungen</a
							>
						</li>
						<li>
							<a class="index-link" href="#m13">Maßgebliche Rechtsgrundlagen</a>
						</li>
						<li><a class="index-link" href="#m27">Sicherheitsmaßnahmen</a></li>
						<li>
							<a class="index-link" href="#m354"
								>Wahrnehmung von Aufgaben nach Satzung oder Geschäftsordnung</a
							>
						</li>
						<li>
							<a class="index-link" href="#m136"
								>Präsenzen in sozialen Netzwerken (Social Media)</a
							>
						</li>
						<li>
							<a class="index-link" href="#m15"
								>Änderung und Aktualisierung der Datenschutzerklärung</a
							>
						</li>
						<li><a class="index-link" href="#m42">Begriffsdefinitionen</a></li>
					</ul>
					<h2 id="m3">Verantwortlicher</h2>
					<p>
						Dorfhäxe Rümmingen 1975 e.V.<br />Marco Pollini<br />Lörracher Str.
						1<br />79595 Rümmingen<br /><a href="mailto:kontakt@dorfhaexe.de"
							>kontakt@dorfhaexe.de</a
						><br />
					</p>
					E-Mail-Adresse:
					<p><a href="mailto:kontakt@dorfhaexe.de">kontakt@dorfhaexe.de</a></p>
					<h2 id="mOverview">Übersicht der Verarbeitungen</h2>
					<p>
						Die nachfolgende Übersicht fasst die Arten der verarbeiteten Daten
						und die Zwecke ihrer Verarbeitung zusammen und verweist auf die
						betroffenen Personen.
					</p>
					<h3>Arten der verarbeiteten Daten</h3>
					<ul>
						<li>Bestandsdaten.</li>
						<li>Zahlungsdaten.</li>
						<li>Kontaktdaten.</li>
						<li>Inhaltsdaten.</li>
						<li>Vertragsdaten.</li>
						<li>Nutzungsdaten.</li>
						<li>Meta-/Kommunikationsdaten.</li>
					</ul>
					<h3>Kategorien betroffener Personen</h3>
					<ul>
						<li>Nutzer.</li>
						<li>Mitglieder.</li>
						<li>Geschäfts- und Vertragspartner.</li>
					</ul>
					<h3>Zwecke der Verarbeitung</h3>
					<ul>
						<li>Erbringung vertraglicher Leistungen und Kundenservice.</li>
						<li>Kontaktanfragen und Kommunikation.</li>
						<li>Verwaltung und Beantwortung von Anfragen.</li>
						<li>Feedback.</li>
						<li>Marketing.</li>
					</ul>
					<h3 id="m13">Maßgebliche Rechtsgrundlagen</h3>
					<p>
						Im Folgenden erhalten Sie eine Übersicht der Rechtsgrundlagen der
						DSGVO, auf deren Basis wir personenbezogene Daten verarbeiten. Bitte
						nehmen Sie zur Kenntnis, dass neben den Regelungen der DSGVO
						nationale Datenschutzvorgaben in Ihrem bzw. unserem Wohn- oder
						Sitzland gelten können. Sollten ferner im Einzelfall speziellere
						Rechtsgrundlagen maßgeblich sein, teilen wir Ihnen diese in der
						Datenschutzerklärung mit.
					</p>
					<ul>
						<li>
							<strong
								>Vertragserfüllung und vorvertragliche Anfragen (Art. 6 Abs. 1
								S. 1 lit. b) DSGVO)</strong
							>
							- Die Verarbeitung ist für die Erfüllung eines Vertrags, dessen
							Vertragspartei die betroffene Person ist, oder zur Durchführung
							vorvertraglicher Maßnahmen erforderlich, die auf Anfrage der
							betroffenen Person erfolgen.
						</li>
						<li>
							<strong
								>Berechtigte Interessen (Art. 6 Abs. 1 S. 1 lit. f)
								DSGVO)</strong
							>
							- Die Verarbeitung ist zur Wahrung der berechtigten Interessen des
							Verantwortlichen oder eines Dritten erforderlich, sofern nicht die
							Interessen oder Grundrechte und Grundfreiheiten der betroffenen
							Person, die den Schutz personenbezogener Daten erfordern,
							überwiegen.
						</li>
					</ul>
					<p>
						Zusätzlich zu den Datenschutzregelungen der
						Datenschutz-Grundverordnung gelten nationale Regelungen zum
						Datenschutz in Deutschland. Hierzu gehört insbesondere das Gesetz
						zum Schutz vor Missbrauch personenbezogener Daten bei der
						Datenverarbeitung (Bundesdatenschutzgesetz – BDSG). Das BDSG enthält
						insbesondere Spezialregelungen zum Recht auf Auskunft, zum Recht auf
						Löschung, zum Widerspruchsrecht, zur Verarbeitung besonderer
						Kategorien personenbezogener Daten, zur Verarbeitung für andere
						Zwecke und zur Übermittlung sowie automatisierten
						Entscheidungsfindung im Einzelfall einschließlich Profiling. Des
						Weiteren regelt es die Datenverarbeitung für Zwecke des
						Beschäftigungsverhältnisses (§ 26 BDSG), insbesondere im Hinblick
						auf die Begründung, Durchführung oder Beendigung von
						Beschäftigungsverhältnissen sowie die Einwilligung von
						Beschäftigten. Ferner können Landesdatenschutzgesetze der einzelnen
						Bundesländer zur Anwendung gelangen.
					</p>
					<h2 id="m27">Sicherheitsmaßnahmen</h2>
					<p>
						Wir treffen nach Maßgabe der gesetzlichen Vorgaben unter
						Berücksichtigung des Stands der Technik, der Implementierungskosten
						und der Art, des Umfangs, der Umstände und der Zwecke der
						Verarbeitung sowie der unterschiedlichen
						Eintrittswahrscheinlichkeiten und des Ausmaßes der Bedrohung der
						Rechte und Freiheiten natürlicher Personen geeignete technische und
						organisatorische Maßnahmen, um ein dem Risiko angemessenes
						Schutzniveau zu gewährleisten.
					</p>
					<p>
						Zu den Maßnahmen gehören insbesondere die Sicherung der
						Vertraulichkeit, Integrität und Verfügbarkeit von Daten durch
						Kontrolle des physischen und elektronischen Zugangs zu den Daten als
						auch des sie betreffenden Zugriffs, der Eingabe, der Weitergabe, der
						Sicherung der Verfügbarkeit und ihrer Trennung. Des Weiteren haben
						wir Verfahren eingerichtet, die eine Wahrnehmung von
						Betroffenenrechten, die Löschung von Daten und Reaktionen auf die
						Gefährdung der Daten gewährleisten. Ferner berücksichtigen wir den
						Schutz personenbezogener Daten bereits bei der Entwicklung bzw.
						Auswahl von Hardware, Software sowie Verfahren entsprechend dem
						Prinzip des Datenschutzes, durch Technikgestaltung und durch
						datenschutzfreundliche Voreinstellungen.
					</p>
					<p>
						TLS-Verschlüsselung (https): Um Ihre via unserem Online-Angebot
						übermittelten Daten zu schützen, nutzen wir eine
						TLS-Verschlüsselung. Sie erkennen derart verschlüsselte Verbindungen
						an dem Präfix https:// in der Adresszeile Ihres Browsers.
					</p>
					<h2 id="m354">
						Wahrnehmung von Aufgaben nach Satzung oder Geschäftsordnung
					</h2>
					<p>
						Wir verarbeiten die Daten unserer Mitglieder, Unterstützer,
						Interessenten, Geschäftspartner oder sonstiger Personen
						(Zusammenfassend "Betroffene"), wenn wir mit ihnen in einem
						Mitgliedschafts- oder sonstigem geschäftlichen Verhältnis stehen und
						unsere Aufgaben wahrnehmen sowie Empfänger von Leistungen und
						Zuwendungen sind. Im Übrigen verarbeiten wir die Daten Betroffener
						auf Grundlage unserer berechtigten Interessen, z.B. wenn es sich um
						administrative Aufgaben oder Öffentlichkeitsarbeit handelt.
					</p>
					<p>
						Die hierbei verarbeiteten Daten, die Art, der Umfang und der Zweck
						und die Erforderlichkeit ihrer Verarbeitung, bestimmen sich nach dem
						zugrundeliegenden Mitgliedschafts- oder Vertragsverhältnis, aus dem
						sich auch die Erforderlichkeit etwaiger Datenangaben ergeben (im
						Übrigen weisen wir auf erforderliche Daten hin).
					</p>
					<p>
						Wir löschen Daten, die zur Erbringung unserer satzungs- und
						geschäftsmäßigen Zwecke nicht mehr erforderlich sind. Dies bestimmt
						sich entsprechend der jeweiligen Aufgaben und vertraglichen
						Beziehungen. Wir bewahren die Daten so lange auf, wie sie zur
						Geschäftsabwicklung, als auch im Hinblick auf etwaige
						Gewährleistungs- oder Haftungspflichten auf Grundlage unserer
						berechtigten Interesse an deren Regelung relevant sein können. Die
						Erforderlichkeit der Aufbewahrung der Daten wird regelmäßig
						überprüft; im Übrigen gelten die gesetzlichen
						Aufbewahrungspflichten.
					</p>
					<ul class="m-elements">
						<li>
							<strong>Verarbeitete Datenarten:</strong> Bestandsdaten (z.B.
							Namen, Adressen); Zahlungsdaten (z.B. Bankverbindungen,
							Rechnungen, Zahlungshistorie); Kontaktdaten (z.B. E-Mail,
							Telefonnummern); Vertragsdaten (z.B. Vertragsgegenstand, Laufzeit,
							Kundenkategorie).
						</li>
						<li>
							<strong>Betroffene Personen:</strong> Nutzer (z.B.
							Webseitenbesucher, Nutzer von Onlinediensten); Mitglieder;
							Geschäfts- und Vertragspartner.
						</li>
						<li>
							<strong>Zwecke der Verarbeitung:</strong> Erbringung vertraglicher
							Leistungen und Kundenservice; Kontaktanfragen und Kommunikation;
							Verwaltung und Beantwortung von Anfragen.
						</li>
						<li>
							<strong>Rechtsgrundlagen:</strong> Vertragserfüllung und
							vorvertragliche Anfragen (Art. 6 Abs. 1 S. 1 lit. b) DSGVO);
							Berechtigte Interessen (Art. 6 Abs. 1 S. 1 lit. f) DSGVO).
						</li>
					</ul>
					<h2 id="m136">Präsenzen in sozialen Netzwerken (Social Media)</h2>
					<p>
						Wir unterhalten Onlinepräsenzen innerhalb sozialer Netzwerke und
						verarbeiten in diesem Rahmen Daten der Nutzer, um mit den dort
						aktiven Nutzern zu kommunizieren oder um Informationen über uns
						anzubieten.
					</p>
					<p>
						Wir weisen darauf hin, dass dabei Daten der Nutzer außerhalb des
						Raumes der Europäischen Union verarbeitet werden können. Hierdurch
						können sich für die Nutzer Risiken ergeben, weil so z.B. die
						Durchsetzung der Rechte der Nutzer erschwert werden könnte.
					</p>
					<p>
						Ferner werden die Daten der Nutzer innerhalb sozialer Netzwerke im
						Regelfall für Marktforschungs- und Werbezwecke verarbeitet. So
						können z.B. anhand des Nutzungsverhaltens und sich daraus ergebender
						Interessen der Nutzer Nutzungsprofile erstellt werden. Die
						Nutzungsprofile können wiederum verwendet werden, um z.B.
						Werbeanzeigen innerhalb und außerhalb der Netzwerke zu schalten, die
						mutmaßlich den Interessen der Nutzer entsprechen. Zu diesen Zwecken
						werden im Regelfall Cookies auf den Rechnern der Nutzer gespeichert,
						in denen das Nutzungsverhalten und die Interessen der Nutzer
						gespeichert werden. Ferner können in den Nutzungsprofilen auch Daten
						unabhängig der von den Nutzern verwendeten Geräte gespeichert werden
						(insbesondere, wenn die Nutzer Mitglieder der jeweiligen Plattformen
						sind und bei diesen eingeloggt sind).
					</p>
					<p>
						Für eine detaillierte Darstellung der jeweiligen Verarbeitungsformen
						und der Widerspruchsmöglichkeiten (Opt-Out) verweisen wir auf die
						Datenschutzerklärungen und Angaben der Betreiber der jeweiligen
						Netzwerke.
					</p>
					<p>
						Auch im Fall von Auskunftsanfragen und der Geltendmachung von
						Betroffenenrechten weisen wir darauf hin, dass diese am effektivsten
						bei den Anbietern geltend gemacht werden können. Nur die Anbieter
						haben jeweils Zugriff auf die Daten der Nutzer und können direkt
						entsprechende Maßnahmen ergreifen und Auskünfte geben. Sollten Sie
						dennoch Hilfe benötigen, dann können Sie sich an uns wenden.
					</p>
					<ul class="m-elements">
						<li>
							<strong>Verarbeitete Datenarten:</strong> Kontaktdaten (z.B.
							E-Mail, Telefonnummern); Inhaltsdaten (z.B. Eingaben in
							Onlineformularen); Nutzungsdaten (z.B. besuchte Webseiten,
							Interesse an Inhalten, Zugriffszeiten); Meta-/Kommunikationsdaten
							(z.B. Geräte-Informationen, IP-Adressen).
						</li>
						<li>
							<strong>Betroffene Personen:</strong> Nutzer (z.B.
							Webseitenbesucher, Nutzer von Onlinediensten).
						</li>
						<li>
							<strong>Zwecke der Verarbeitung:</strong> Kontaktanfragen und
							Kommunikation; Feedback (z.B. Sammeln von Feedback via
							Online-Formular); Marketing.
						</li>
						<li>
							<strong>Rechtsgrundlagen:</strong> Berechtigte Interessen (Art. 6
							Abs. 1 S. 1 lit. f) DSGVO).
						</li>
					</ul>
					<p>
						<strong
							>Weitere Hinweise zu Verarbeitungsprozessen, Verfahren und
							Diensten:</strong
						>
					</p>
					<ul class="m-elements">
						<li>
							<strong>Instagram: </strong>Soziales Netzwerk;
							<strong>Dienstanbieter:</strong> Meta Platforms Irland Limited, 4
							Grand Canal Square, Grand Canal Harbour, Dublin 2, Irland;
							<strong>Rechtsgrundlagen:</strong> Berechtigte Interessen (Art. 6
							Abs. 1 S. 1 lit. f) DSGVO); <strong>Website:</strong>
							<a href="https://www.instagram.com" target="_blank"
								>https://www.instagram.com</a
							>; <strong>Datenschutzerklärung:</strong>
							<a
								href="https://instagram.com/about/legal/privacy"
								target="_blank"
								>https://instagram.com/about/legal/privacy</a
							>.
						</li>
						<li>
							<strong>Facebook-Seiten: </strong>Profile innerhalb des sozialen
							Netzwerks Facebook - Wir sind gemeinsam mit Meta Platforms Ireland
							Limited für die Erhebung (jedoch nicht die weitere Verarbeitung)
							von Daten der Besucher unserer Facebook-Seite (sog. "Fanpage")
							verantwortlich. Zu diesen Daten gehören Informationen zu den Arten
							von Inhalten, die Nutzer sich ansehen oder mit denen sie
							interagieren, oder die von ihnen vorgenommenen Handlungen (siehe
							unter „Von dir und anderen getätigte und bereitgestellte Dinge“ in
							der Facebook-Datenrichtlinie:
							<a href="https://www.facebook.com/policy" target="_blank"
								>https://www.facebook.com/policy</a
							>), sowie Informationen über die von den Nutzern genutzten Geräte
							(z. B. IP-Adressen, Betriebssystem, Browsertyp,
							Spracheinstellungen, Cookie-Daten; siehe unter
							„Geräteinformationen“ in der Facebook-Datenrichtlinie:
							<a href="https://www.facebook.com/policy" target="_blank"
								>https://www.facebook.com/policy</a
							>). Wie in der Facebook-Datenrichtlinie unter „Wie verwenden wir
							diese Informationen?“ erläutert, erhebt und verwendet Facebook
							Informationen auch, um Analysedienste, so genannte
							"Seiten-Insights", für Seitenbetreiber bereitzustellen, damit
							diese Erkenntnisse darüber erhalten, wie Personen mit ihren Seiten
							und mit den mit ihnen verbundenen Inhalten interagieren. Wir haben
							mit Facebook eine spezielle Vereinbarung abgeschlossen
							("Informationen zu Seiten-Insights",
							<a
								href="https://www.facebook.com/legal/terms/page_controller_addendum"
								target="_blank"
								>https://www.facebook.com/legal/terms/page_controller_addendum</a
							>), in der insbesondere geregelt wird, welche Sicherheitsmaßnahmen
							Facebook beachten muss und in der Facebook sich bereit erklärt hat
							die Betroffenenrechte zu erfüllen (d. h. Nutzer können z. B.
							Auskünfte oder Löschungsanfragen direkt an Facebook richten). Die
							Rechte der Nutzer (insbesondere auf Auskunft, Löschung,
							Widerspruch und Beschwerde bei zuständiger Aufsichtsbehörde),
							werden durch die Vereinbarungen mit Facebook nicht eingeschränkt.
							Weitere Hinweise finden sich in den "Informationen zu
							Seiten-Insights" (<a
								href="https://www.facebook.com/legal/terms/information_about_page_insights_data"
								target="_blank"
								>https://www.facebook.com/legal/terms/information_about_page_insights_data</a
							>); <strong>Dienstanbieter:</strong> Meta Platforms Ireland
							Limited, 4 Grand Canal Square, Grand Canal Harbour, Dublin 2,
							Irland; <strong>Rechtsgrundlagen:</strong> Berechtigte Interessen
							(Art. 6 Abs. 1 S. 1 lit. f) DSGVO); <strong>Website:</strong>
							<a href="https://www.facebook.com" target="_blank"
								>https://www.facebook.com</a
							>; <strong>Datenschutzerklärung:</strong>
							<a href="https://www.facebook.com/about/privacy" target="_blank"
								>https://www.facebook.com/about/privacy</a
							>;
							<strong
								>Standardvertragsklauseln (Gewährleistung Datenschutzniveau bei
								Verarbeitung in Drittländern):</strong
							>
							<a
								href="https://www.facebook.com/legal/EU_data_transfer_addendum"
								target="_blank"
								>https://www.facebook.com/legal/EU_data_transfer_addendum</a
							>; <strong>Weitere Informationen:</strong> Vereinbarung
							gemeinsamer Verantwortlichkeit:
							<a
								href="https://www.facebook.com/legal/terms/information_about_page_insights_data"
								target="_blank"
								>https://www.facebook.com/legal/terms/information_about_page_insights_data</a
							>. Die gemeinsame Verantwortlichkeit beschränkt sich auf die
							Erhebung durch und Übermittlung von Daten an Meta Platforms
							Ireland Limited, ein Unternehmen mit Sitz in der EU. Die weitere
							Verarbeitung der Daten liegt in der alleinigen Verantwortung von
							Meta Platforms Ireland Limited, was insbesondere die Übermittlung
							der Daten an die Muttergesellschaft Meta Platforms, Inc. in den
							USA betrifft (auf der Grundlage der zwischen Meta Platforms
							Ireland Limited und Meta Platforms, Inc. geschlossenen
							Standardvertragsklauseln).
						</li>
					</ul>
					<h2 id="m15">Änderung und Aktualisierung der Datenschutzerklärung</h2>
					<p>
						Wir bitten Sie, sich regelmäßig über den Inhalt unserer
						Datenschutzerklärung zu informieren. Wir passen die
						Datenschutzerklärung an, sobald die Änderungen der von uns
						durchgeführten Datenverarbeitungen dies erforderlich machen. Wir
						informieren Sie, sobald durch die Änderungen eine
						Mitwirkungshandlung Ihrerseits (z.B. Einwilligung) oder eine
						sonstige individuelle Benachrichtigung erforderlich wird.
					</p>
					<p>
						Sofern wir in dieser Datenschutzerklärung Adressen und
						Kontaktinformationen von Unternehmen und Organisationen angeben,
						bitten wir zu beachten, dass die Adressen sich über die Zeit ändern
						können und bitten die Angaben vor Kontaktaufnahme zu prüfen.
					</p>
					<h2 id="m42">Begriffsdefinitionen</h2>
					<p>
						In diesem Abschnitt erhalten Sie eine Übersicht über die in dieser
						Datenschutzerklärung verwendeten Begrifflichkeiten. Viele der
						Begriffe sind dem Gesetz entnommen und vor allem im Art. 4 DSGVO
						definiert. Die gesetzlichen Definitionen sind verbindlich. Die
						nachfolgenden Erläuterungen sollen dagegen vor allem dem Verständnis
						dienen. Die Begriffe sind alphabetisch sortiert.
					</p>
					<ul class="glossary">
						<li>
							<strong>Personenbezogene Daten:</strong> "Personenbezogene Daten“
							sind alle Informationen, die sich auf eine identifizierte oder
							identifizierbare natürliche Person (im Folgenden "betroffene
							Person“) beziehen; als identifizierbar wird eine natürliche Person
							angesehen, die direkt oder indirekt, insbesondere mittels
							Zuordnung zu einer Kennung wie einem Namen, zu einer Kennnummer,
							zu Standortdaten, zu einer Online-Kennung (z.B. Cookie) oder zu
							einem oder mehreren besonderen Merkmalen identifiziert werden
							kann, die Ausdruck der physischen, physiologischen, genetischen,
							psychischen, wirtschaftlichen, kulturellen oder sozialen Identität
							dieser natürlichen Person sind.
						</li>
						<li>
							<strong>Verantwortlicher:</strong> Als "Verantwortlicher“ wird die
							natürliche oder juristische Person, Behörde, Einrichtung oder
							andere Stelle, die allein oder gemeinsam mit anderen über die
							Zwecke und Mittel der Verarbeitung von personenbezogenen Daten
							entscheidet, bezeichnet.
						</li>
						<li>
							<strong>Verarbeitung:</strong> "Verarbeitung" ist jeder mit oder
							ohne Hilfe automatisierter Verfahren ausgeführte Vorgang oder jede
							solche Vorgangsreihe im Zusammenhang mit personenbezogenen Daten.
							Der Begriff reicht weit und umfasst praktisch jeden Umgang mit
							Daten, sei es das Erheben, das Auswerten, das Speichern, das
							Übermitteln oder das Löschen.
						</li>
					</ul>
					<p class="seal">
						<a
							href="https://datenschutz-generator.de/"
							title="Rechtstext von Dr. Schwenke - für weitere Informationen bitte anklicken."
							target="_blank"
							rel="noopener noreferrer nofollow"
							>Erstellt mit kostenlosem Datenschutz-Generator.de von Dr. Thomas
							Schwenke</a
						>
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
		impressumIsOpen.value = true;
	};

	let privacyIsOpen = ref(false);
	const showPrivacy = () => {
		privacyIsOpen.value = true;
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
