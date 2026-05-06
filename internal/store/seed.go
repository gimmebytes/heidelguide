package store

import "fmt"

// seedLandmark holds the data for a single landmark to be seeded.
type seedLandmark struct {
	Latitude      float64
	Longitude     float64
	ImageFilename string
	YearBuilt     int
	YearDestroyed *int
	Translations  []seedTranslation
}

// seedTranslation holds locale-specific content for a landmark.
type seedTranslation struct {
	Locale      string
	Name        string
	Description string
	History     string
}

// intPtr is a helper to create a pointer to an int.
func intPtr(v int) *int {
	return &v
}

// Seed inserts initial landmark data if the landmarks table is empty.
func (s *Store) Seed() error {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM landmarks").Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check landmarks count: %w", err)
	}
	if count > 0 {
		return nil
	}

	landmarks := seedData()

	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin seed transaction: %w", err)
	}
	defer tx.Rollback()

	landmarkStmt, err := tx.Prepare(`INSERT INTO landmarks (latitude, longitude, image_filename, year_built, year_destroyed) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare landmark insert: %w", err)
	}
	defer landmarkStmt.Close()

	translationStmt, err := tx.Prepare(`INSERT INTO landmark_translations (landmark_id, locale, name, description, history) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed to prepare translation insert: %w", err)
	}
	defer translationStmt.Close()

	for _, lm := range landmarks {
		result, err := landmarkStmt.Exec(lm.Latitude, lm.Longitude, lm.ImageFilename, lm.YearBuilt, lm.YearDestroyed)
		if err != nil {
			return fmt.Errorf("failed to insert landmark: %w", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			return fmt.Errorf("failed to get landmark id: %w", err)
		}
		for _, t := range lm.Translations {
			if _, err := translationStmt.Exec(id, t.Locale, t.Name, t.Description, t.History); err != nil {
				return fmt.Errorf("failed to insert translation for landmark %d: %w", id, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit seed transaction: %w", err)
	}
	return nil
}

// seedData returns the 8 Heidelberg landmarks with German and English translations.
func seedData() []seedLandmark {
	return []seedLandmark{
		{
			Latitude:      49.4105,
			Longitude:     8.7153,
			ImageFilename: "castle.jpg",
			YearBuilt:     1214,
			YearDestroyed: intPtr(1693),
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Heidelberger Schloss",
					Description: "Die berühmte Schlossruine thront über der Altstadt und ist eines der bedeutendsten Renaissancebauwerke nördlich der Alpen.",
					History:     "Erstmals 1214 urkundlich erwähnt, wurde das Schloss im Pfälzischen Erbfolgekrieg 1693 durch französische Truppen schwer beschädigt. Die Ruine wurde im 19. Jahrhundert zum Symbol der Romantik. Heute beherbergt das Schloss das Deutsche Apothekenmuseum und das größte Weinfass der Welt.",
				},
				{
					Locale:      "en",
					Name:        "Heidelberg Castle",
					Description: "The famous castle ruin towers above the old town and is one of the most significant Renaissance structures north of the Alps.",
					History:     "First mentioned in records in 1214, the castle was severely damaged by French troops during the War of the Palatinate Succession in 1693. The ruin became a symbol of Romanticism in the 19th century. Today the castle houses the German Pharmacy Museum and the world's largest wine barrel.",
				},
			},
		},
		{
			Latitude:      49.4133,
			Longitude:     8.7105,
			ImageFilename: "old_bridge.jpg",
			YearBuilt:     1788,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Alte Brücke (Karl-Theodor-Brücke)",
					Description: "Die barocke Steinbrücke über den Neckar gehört zu den ältesten Brücken Deutschlands und ist ein Wahrzeichen der Stadt.",
					History:     "Die heutige Brücke wurde 1788 unter Kurfürst Karl Theodor erbaut und ersetzte zahlreiche hölzerne Vorgängerbauten, die durch Hochwasser und Eisgang zerstört worden waren. Im Zweiten Weltkrieg wurde sie 1945 von der Wehrmacht gesprengt und bis 1947 wieder aufgebaut. Das Brückentor am Südufer stammt aus dem Mittelalter und diente einst als Stadttor.",
				},
				{
					Locale:      "en",
					Name:        "Old Bridge (Karl Theodor Bridge)",
					Description: "The baroque stone bridge over the Neckar is one of the oldest bridges in Germany and an iconic landmark of the city.",
					History:     "The current bridge was built in 1788 under Elector Karl Theodor, replacing numerous wooden predecessors destroyed by floods and ice. It was blown up by the Wehrmacht in 1945 during World War II and rebuilt by 1947. The bridge gate on the south bank dates from the Middle Ages and once served as a city gate.",
				},
			},
		},
		{
			Latitude:      49.4167,
			Longitude:     8.7000,
			ImageFilename: "philosophers_walk.jpg",
			YearBuilt:     1817,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Philosophenweg",
					Description: "Der berühmte Höhenweg am Nordufer des Neckars bietet einen der schönsten Ausblicke auf die Altstadt, das Schloss und das Neckartal.",
					History:     "Der Weg wurde ab 1817 angelegt und erhielt seinen Namen, weil Heidelberger Professoren und Denker hier ihre Spaziergänge unternahmen. Dank des milden Mikroklimas am Südhang des Heiligenbergs gedeihen hier mediterrane Pflanzen wie Zypressen und Zitronenbäume. Der Philosophengarten am oberen Ende wurde 1979 als öffentliche Parkanlage gestaltet.",
				},
				{
					Locale:      "en",
					Name:        "Philosophers' Walk",
					Description: "The famous hillside path on the north bank of the Neckar offers one of the most beautiful views of the old town, the castle, and the Neckar valley.",
					History:     "The path was laid out from 1817 onward and received its name because Heidelberg professors and thinkers took their walks here. Thanks to the mild microclimate on the south-facing slope of the Heiligenberg, Mediterranean plants such as cypresses and lemon trees thrive here. The Philosophers' Garden at the upper end was created as a public park in 1979.",
				},
			},
		},
		{
			Latitude:      49.4118,
			Longitude:     8.7063,
			ImageFilename: "holy_spirit_church.jpg",
			YearBuilt:     1398,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Heiliggeistkirche",
					Description: "Die gotische Hallenkirche am Marktplatz ist die größte und bedeutendste Kirche Heidelbergs und prägt die Silhouette der Altstadt.",
					History:     "Der Bau der Heiliggeistkirche begann 1398 unter König Ruprecht III. und wurde um 1544 vollendet. Sie diente als Grablege der pfälzischen Kurfürsten und beherbergte einst die berühmte Bibliotheca Palatina, die 1623 als Kriegsbeute nach Rom gebracht wurde. Von 1706 bis 1936 war die Kirche durch eine Scheidemauer in einen protestantischen und einen katholischen Teil getrennt.",
				},
				{
					Locale:      "en",
					Name:        "Church of the Holy Spirit",
					Description: "The Gothic hall church on the market square is the largest and most important church in Heidelberg, defining the old town's skyline.",
					History:     "Construction of the Church of the Holy Spirit began in 1398 under King Ruprecht III and was completed around 1544. It served as the burial place of the Palatinate Electors and once housed the famous Bibliotheca Palatina, which was taken to Rome as war spoils in 1623. From 1706 to 1936, the church was divided by a partition wall into Protestant and Catholic sections.",
				},
			},
		},
		{
			Latitude:      49.4107,
			Longitude:     8.7068,
			ImageFilename: "student_prison.jpg",
			YearBuilt:     1778,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Studentenkarzer",
					Description: "Das historische Studentengefängnis der Universität Heidelberg diente von 1778 bis 1914 zur Bestrafung ungehorsamer Studenten.",
					History:     "Der Karzer wurde 1778 eingerichtet, nachdem die Universität das Recht zur eigenen Gerichtsbarkeit über ihre Studenten erhalten hatte. Studenten wurden hier für Vergehen wie Ruhestörung, Trunkenheit oder Duelle für zwei bis vier Wochen eingesperrt. Die Wände sind mit farbenfrohen Graffiti, Schattenrissen und Inschriften der Insassen bedeckt, die heute als einzigartiges Kulturdenkmal gelten.",
				},
				{
					Locale:      "en",
					Name:        "Student Prison",
					Description: "The historic student prison of Heidelberg University was used from 1778 to 1914 to punish disobedient students.",
					History:     "The prison was established in 1778 after the university received the right to its own jurisdiction over students. Students were locked up here for two to four weeks for offenses such as disturbing the peace, drunkenness, or dueling. The walls are covered with colorful graffiti, silhouettes, and inscriptions by former inmates, now considered a unique cultural monument.",
				},
			},
		},
		{
			Latitude:      49.4098,
			Longitude:     8.7060,
			ImageFilename: "university_library.jpg",
			YearBuilt:     1905,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Universitätsbibliothek Heidelberg",
					Description: "Die Universitätsbibliothek ist die älteste Universitätsbibliothek Deutschlands und beherbergt wertvolle mittelalterliche Handschriften.",
					History:     "Das heutige Hauptgebäude im Jugendstil wurde 1905 nach Plänen von Josef Durm errichtet. Die Bibliothek wurde 1386 zusammen mit der Universität gegründet und ist damit die älteste in Deutschland. Zu ihren bedeutendsten Schätzen zählt der Codex Manesse, die umfangreichste Sammlung mittelhochdeutscher Liederdichtung aus dem 14. Jahrhundert.",
				},
				{
					Locale:      "en",
					Name:        "Heidelberg University Library",
					Description: "The University Library is the oldest university library in Germany and houses valuable medieval manuscripts.",
					History:     "The current Art Nouveau main building was constructed in 1905 based on designs by Josef Durm. The library was founded in 1386 together with the university, making it the oldest in Germany. Among its most significant treasures is the Codex Manesse, the most comprehensive collection of Middle High German lyric poetry from the 14th century.",
				},
			},
		},
		{
			Latitude:      49.3990,
			Longitude:     8.7280,
			ImageFilename: "koenigstuhl.jpg",
			YearBuilt:     0,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Königstuhl",
					Description: "Der 568 Meter hohe Königstuhl ist der Hausberg Heidelbergs und seit 1890 mit der historischen Bergbahn erreichbar.",
					History:     "Der Name Königstuhl geht vermutlich auf die mittelalterliche Königswahl zurück, bei der die Kurfürsten auf dem Berg zusammenkamen. Die Heidelberger Bergbahn, eine der ältesten elektrischen Standseilbahnen Deutschlands, wurde 1890 eröffnet und verbindet die Altstadt über die Molkenkur mit dem Gipfel. Auf dem Königstuhl befinden sich heute eine Sternwarte und ein Märchenparadies für Kinder.",
				},
				{
					Locale:      "en",
					Name:        "Königstuhl",
					Description: "The 568-meter Königstuhl is Heidelberg's local mountain and has been accessible by the historic funicular railway since 1890.",
					History:     "The name Königstuhl (King's Throne) likely dates back to medieval royal elections when the Electors gathered on the mountain. The Heidelberg funicular, one of the oldest electric funicular railways in Germany, opened in 1890 and connects the old town via the Molkenkur station to the summit. Today the Königstuhl is home to an observatory and a fairy-tale theme park for children.",
				},
			},
		},
		{
			Latitude:      49.4150,
			Longitude:     8.6950,
			ImageFilename: "neckar_meadow.jpg",
			YearBuilt:     0,
			YearDestroyed: nil,
			Translations: []seedTranslation{
				{
					Locale:      "de",
					Name:        "Neckarwiese",
					Description: "Die weitläufige Grünfläche am Nordufer des Neckars ist der beliebteste Treffpunkt der Heidelberger für Freizeit und Erholung.",
					History:     "Die Neckarwiese entstand in ihrer heutigen Form durch die Neckarregulierung im 19. Jahrhundert, als die Ufer befestigt und die Überschwemmungsgebiete in Parkflächen umgewandelt wurden. Sie erstreckt sich vom Bismarckplatz bis zur Theodor-Heuss-Brücke und bietet einen unverstellten Blick auf Schloss und Altstadt. Im Sommer ist sie Schauplatz von Grillabenden, Sportveranstaltungen und dem traditionellen Heidelberger Herbst.",
				},
				{
					Locale:      "en",
					Name:        "Neckar Meadow",
					Description: "The expansive green space on the north bank of the Neckar is Heidelberg's most popular gathering place for leisure and recreation.",
					History:     "The Neckar Meadow took its current form through the Neckar river regulation in the 19th century, when the banks were reinforced and flood plains were converted into parkland. It stretches from Bismarckplatz to the Theodor Heuss Bridge and offers an unobstructed view of the castle and old town. In summer it hosts barbecue evenings, sports events, and the traditional Heidelberg Autumn festival.",
				},
			},
		},
	}
}
