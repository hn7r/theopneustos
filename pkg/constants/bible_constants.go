package constants

const (
	GenesisID             = 1
	ExodusID              = 2
	LeviticusID           = 3
	NumbersID             = 4
	DeuteronomyID         = 5
	JoshuaID              = 6
	JudgesID              = 7
	RuthID                = 8
	FirstSamuelID         = 9
	SecondSamuelID        = 10
	FirstKingsID          = 11
	SecondKingsID         = 12
	FirstChroniclesID     = 13
	SecondChroniclesID    = 14
	EzraID                = 15
	NehemiahID            = 16
	EstherID              = 17
	JobID                 = 18
	PsalmsID              = 19
	ProverbsID            = 20
	EcclesiastesID        = 21
	SongOfSolomonID       = 22
	IsaiahID              = 23
	JeremiahID            = 24
	LamentationsID        = 25
	EzekielID             = 26
	DanielID              = 27
	HoseaID               = 28
	JoelID                = 29
	AmosID                = 30
	ObadiahID             = 31
	JonahID               = 32
	MicahID               = 33
	NahumID               = 34
	HabakkukID            = 35
	ZephaniahID           = 36
	HaggaiID              = 37
	ZechariahID           = 38
	MalachiID             = 39
	MatthewID             = 40
	MarkID                = 41
	LukeID                = 42
	JohnID                = 43
	ActsID                = 44
	RomansID              = 45
	FirstCorinthiansID    = 46
	SecondCorinthiansID   = 47
	GalatiansID           = 48
	EphesiansID           = 49
	PhilippiansID         = 50
	ColossiansID          = 51
	FirstThessaloniansID  = 52
	SecondThessaloniansID = 53
	FirstTimothyID        = 54
	SecondTimothyID       = 55
	TitusID               = 56
	PhilemonID            = 57
	HebrewsID             = 58
	JamesID               = 59
	FirstPeterID          = 60
	SecondPeterID         = 61
	FirstJohnID           = 62
	SecondJohnID          = 63
	ThirdJohnID           = 64
	JudeID                = 65
	RevelationID          = 66
)

var BookNumberToName = map[int]string{
	GenesisID:             "GENESIS",
	ExodusID:              "EXODUS",
	LeviticusID:           "LEVITICUS",
	NumbersID:             "NUMBERS",
	DeuteronomyID:         "DEUTERONOMY",
	JoshuaID:              "JOSHUA",
	JudgesID:              "JUDGES",
	RuthID:                "RUTH",
	FirstSamuelID:         "1 SAMUEL",
	SecondSamuelID:        "2 SAMUEL",
	FirstKingsID:          "1 KINGS",
	SecondKingsID:         "2 KINGS",
	FirstChroniclesID:     "1 CHRONICLES",
	SecondChroniclesID:    "2 CHRONICLES",
	EzraID:                "EZRA",
	NehemiahID:            "NEHEMIAH",
	EstherID:              "ESTHER",
	JobID:                 "JOB",
	PsalmsID:              "PSALMS",
	ProverbsID:            "PROVERBS",
	EcclesiastesID:        "ECCLESIASTES",
	SongOfSolomonID:       "SONG OF SOLOMON",
	IsaiahID:              "ISAIAH",
	JeremiahID:            "JEREMIAH",
	LamentationsID:        "LAMENTATIONS",
	EzekielID:             "EZEKIEL",
	DanielID:              "DANIEL",
	HoseaID:               "HOSEA",
	JoelID:                "JOEL",
	AmosID:                "AMOS",
	ObadiahID:             "OBADIAH",
	JonahID:               "JONAH",
	MicahID:               "MICAH",
	NahumID:               "NAHUM",
	HabakkukID:            "HABAKKUK",
	ZephaniahID:           "ZEPHANIAH",
	HaggaiID:              "HAGGAI",
	ZechariahID:           "ZECHARIAH",
	MalachiID:             "MALACHI",
	MatthewID:             "MATTHEW",
	MarkID:                "MARK",
	LukeID:                "LUKE",
	JohnID:                "JOHN",
	ActsID:                "ACTS",
	RomansID:              "ROMANS",
	FirstCorinthiansID:    "1 CORINTHIANS",
	SecondCorinthiansID:   "2 CORINTHIANS",
	GalatiansID:           "GALATIANS",
	EphesiansID:           "EPHESIANS",
	PhilippiansID:         "PHILIPPIANS",
	ColossiansID:          "COLOSSIANS",
	FirstThessaloniansID:  "1 THESSALONIANS",
	SecondThessaloniansID: "2 THESSALONIANS",
	FirstTimothyID:        "1 TIMOTHY",
	SecondTimothyID:       "2 TIMOTHY",
	TitusID:               "TITUS",
	PhilemonID:            "PHILEMON",
	HebrewsID:             "HEBREWS",
	JamesID:               "JAMES",
	FirstPeterID:          "1 PETER",
	SecondPeterID:         "2 PETER",
	FirstJohnID:           "1 JOHN",
	SecondJohnID:          "2 JOHN",
	ThirdJohnID:           "3 JOHN",
	JudeID:                "JUDE",
	RevelationID:          "REVELATION",
}

var ChapterCount = map[int]int{
	GenesisID:             50,
	ExodusID:              40,
	LeviticusID:           27,
	NumbersID:             36,
	DeuteronomyID:         34,
	JoshuaID:              24,
	JudgesID:              21,
	RuthID:                4,
	FirstSamuelID:         31,
	SecondSamuelID:        24,
	FirstKingsID:          22,
	SecondKingsID:         25,
	FirstChroniclesID:     29,
	SecondChroniclesID:    36,
	EzraID:                10,
	NehemiahID:            13,
	EstherID:              10,
	JobID:                 42,
	PsalmsID:              150,
	ProverbsID:            31,
	EcclesiastesID:        12,
	SongOfSolomonID:       8,
	IsaiahID:              66,
	JeremiahID:            52,
	LamentationsID:        5,
	EzekielID:             48,
	DanielID:              12,
	HoseaID:               14,
	JoelID:                3,
	AmosID:                9,
	ObadiahID:             1,
	JonahID:               4,
	MicahID:               7,
	NahumID:               3,
	HabakkukID:            3,
	ZephaniahID:           3,
	HaggaiID:              2,
	ZechariahID:           14,
	MalachiID:             4,
	MatthewID:             28,
	MarkID:                16,
	LukeID:                24,
	JohnID:                21,
	ActsID:                28,
	RomansID:              16,
	FirstCorinthiansID:    16,
	SecondCorinthiansID:   13,
	GalatiansID:           6,
	EphesiansID:           6,
	PhilippiansID:         4,
	ColossiansID:          4,
	FirstThessaloniansID:  5,
	SecondThessaloniansID: 3,
	FirstTimothyID:        6,
	SecondTimothyID:       4,
	TitusID:               3,
	PhilemonID:            1,
	HebrewsID:             13,
	JamesID:               5,
	FirstPeterID:          5,
	SecondPeterID:         3,
	FirstJohnID:           5,
	SecondJohnID:          1,
	ThirdJohnID:           1,
	JudeID:                1,
	RevelationID:          22,
}

var VerseCounts = map[int][]int{
	GenesisID: {
		31, 25, 24, 26, 32, 22, 24, 22, 29, 32,
		32, 20, 18, 24, 21, 16, 27, 33, 38, 18,
		34, 24, 20, 67, 34, 35, 46, 22, 35, 43,
		55, 32, 20, 31, 29, 43, 36, 30, 23, 23,
		57, 38, 34, 34, 28, 34, 31, 22, 33, 26,
	},
	ExodusID: {
		22, 25, 22, 31, 23, 30, 25, 32, 35, 29,
		10, 51, 22, 31, 27, 36, 16, 27, 25, 26,
		36, 31, 33, 18, 40, 37, 21, 43, 46, 38,
		18, 35, 23, 35, 35, 38, 29, 31, 43, 38,
	},
	LeviticusID: {
		17, 16, 17, 35, 19, 30, 38, 36, 24, 20,
		47, 8, 59, 57, 33, 34, 16, 30, 37, 27,
		24, 33, 44, 23, 55, 46, 34,
	},
	NumbersID: {
		54, 34, 51, 49, 31, 27, 89, 26, 23, 36,
		35, 16, 33, 45, 41, 50, 13, 32, 22, 29,
		35, 41, 30, 25, 18, 65, 23, 31, 39, 17,
		54, 42, 56, 29, 34, 13,
	},
	DeuteronomyID: {
		46, 37, 29, 49, 33, 25, 26, 20, 29, 22,
		32, 32, 18, 29, 23, 22, 20, 22, 21, 20,
		25, 22, 29, 22, 19, 19, 26, 68, 29, 20,
		30, 52, 29, 12,
	},
	JoshuaID: {
		18, 24, 17, 24, 15, 27, 26, 35, 27, 43,
		23, 24, 33, 15, 63, 10, 18, 28, 51, 9,
		45, 34, 16, 33,
	},
	JudgesID: {
		36, 23, 31, 24, 31, 40, 25, 35, 57, 18,
		40, 15, 25, 20, 20, 31, 13, 31, 30, 48,
		25,
	},
	RuthID: {
		22, 23, 18, 22,
	},
	FirstSamuelID: {
		28, 36, 21, 22, 12, 21, 17, 22, 27, 27,
		15, 25, 23, 52, 35, 23, 58, 30, 24, 42,
		15, 23, 29, 22, 44, 25, 12, 25, 11, 31,
		13,
	},
	SecondSamuelID: {
		27, 32, 39, 12, 25, 23, 29, 18, 13, 19,
		27, 31, 39, 33, 37, 23, 29, 33, 43, 26,
		22, 51, 39, 25,
	},
	FirstKingsID: {
		53, 46, 28, 34, 18, 38, 51, 66, 28, 29,
		43, 33, 34, 31, 34, 34, 24, 46, 21, 43,
		29, 53,
	},
	SecondKingsID: {
		18, 25, 27, 44, 27, 33, 20, 29, 37, 36,
		20, 21, 25, 29, 38, 20, 41, 37, 37, 21,
		26, 20, 37, 20, 30,
	},
	FirstChroniclesID: {
		54, 55, 24, 43, 26, 81, 40, 40, 44, 14,
		47, 40, 14, 17, 29, 43, 27, 17, 19, 8,
		30, 19, 32, 31, 31, 32, 34, 21, 30,
	},
	SecondChroniclesID: {
		17, 18, 17, 22, 14, 42, 22, 18, 31, 19,
		23, 16, 23, 14, 19, 14, 19, 34, 11, 37,
		20, 12, 21, 27, 28, 23, 9, 27, 36, 27,
		21, 33, 25, 33, 27, 23,
	},
	EzraID: {
		11, 70, 13, 24, 17, 22, 28, 36, 15, 44,
	},
	NehemiahID: {
		11, 20, 32, 23, 19, 19, 73, 18, 38, 39,
		36, 47, 31,
	},
	EstherID: {
		22, 23, 15, 17, 14, 14, 10, 17, 32, 3,
	},
	JobID: {
		22, 13, 26, 21, 27, 30, 21, 22, 35, 22,
		20, 25, 28, 22, 35, 22, 16, 21, 29, 29,
		34, 30, 17, 25, 6, 14, 23, 28, 25, 31,
		40, 22, 33, 37, 16, 33, 24, 41, 30, 24,
		34, 17,
	},
	PsalmsID: {
		6, 12, 8, 8, 12, 10, 17, 9, 20, 18,
		7, 8, 6, 7, 5, 11, 15, 50, 14, 9,
		13, 31, 6, 10, 22, 12, 14, 9, 11, 12,
		24, 11, 22, 22, 28, 12, 40, 22, 13, 17,
		13, 11, 5, 26, 17, 11, 9, 14, 20, 23,
		19, 9, 6, 7, 23, 13, 11, 11, 17, 12,
		8, 12, 11, 10, 13, 20, 7, 35, 36, 5,
		24, 20, 28, 23, 10, 12, 20, 72, 13, 19,
		16, 8, 18, 12, 13, 17, 7, 18, 52, 17,
		16, 15, 5, 23, 11, 13, 12, 9, 9, 5,
		8, 28, 22, 35, 45, 48, 43, 13, 31, 7,
		10, 10, 9, 8, 18, 19, 2, 29, 176, 7,
		8, 9, 4, 8, 5, 6, 5, 6, 8, 8,
		3, 18, 3, 3, 21, 26, 9, 8, 24, 14,
		10, 8, 12, 15, 21, 10, 20, 14, 9, 6,
	},
	ProverbsID: {
		33, 22, 35, 27, 23, 35, 27, 36, 18, 32,
		31, 28, 25, 35, 33, 33, 28, 24, 29, 30,
		31, 29, 35, 34, 28, 28, 27, 28, 27, 33,
		31,
	},
	EcclesiastesID: {
		18, 26, 22, 16, 20, 12, 29, 17, 18, 20,
		10, 14,
	},
	SongOfSolomonID: {
		17, 17, 11, 16, 16, 13, 13, 14,
	},
	IsaiahID: {
		31, 22, 26, 6, 30, 13, 25, 22, 21, 34,
		16, 6, 22, 32, 9, 14, 14, 7, 25, 6,
		17, 25, 18, 23, 12, 21, 13, 29, 24, 33,
		9, 20, 24, 17, 10, 22, 38, 22, 8, 31,
		29, 25, 28, 28, 25, 13, 15, 22, 26, 11,
		23, 15, 12, 17, 13, 12, 21, 14, 21, 22,
		11, 12, 19, 12, 25, 24,
	},
	JeremiahID: {
		19, 37, 25, 31, 31, 30, 34, 22, 26, 25,
		23, 17, 27, 22, 21, 21, 27, 23, 15, 18,
		14, 30, 40, 10, 38, 24, 22, 17, 32, 24,
		40, 44, 26, 22, 19, 32, 21, 28, 18, 16,
		18, 22, 13, 30, 5, 28, 7, 47, 39, 46,
		64, 34,
	},
	LamentationsID: {
		22, 22, 66, 22, 22,
	},
	EzekielID: {
		28, 10, 27, 17, 17, 14, 27, 18, 11, 22,
		25, 28, 23, 23, 8, 63, 24, 32, 14, 49,
		32, 31, 49, 27, 17, 21, 36, 26, 21, 26,
		18, 32, 33, 31, 15, 38, 28, 23, 29, 49,
		26, 20, 27, 31, 25, 24, 23, 35,
	},
	DanielID: {
		21, 49, 30, 37, 31, 28, 28, 27, 27, 21,
		45, 13,
	},
	HoseaID: {
		11, 23, 5, 19, 15, 11, 16, 14, 17, 15,
		12, 14, 16, 9,
	},
	JoelID: {
		20, 32, 21,
	},
	AmosID: {
		15, 16, 15, 13, 27, 14, 17, 14, 15,
	},
	ObadiahID: {
		21,
	},
	JonahID: {
		17, 10, 10, 11,
	},
	MicahID: {
		16, 13, 12, 13, 15, 16, 20,
	},
	NahumID: {
		15, 13, 19,
	},
	HabakkukID: {
		17, 20, 19,
	},
	ZephaniahID: {
		18, 15, 20,
	},
	HaggaiID: {
		15, 23,
	},
	ZechariahID: {
		21, 13, 10, 14, 11, 15, 14, 23, 17, 12,
		17, 14, 9, 21,
	},
	MalachiID: {
		14, 17, 18, 6,
	},
	MatthewID: {
		25, 23, 17, 25, 48, 34, 29, 34, 38, 42,
		30, 50, 58, 36, 39, 28, 27, 35, 30, 34,
		46, 46, 39, 51, 46, 75, 66, 20,
	},
	MarkID: {
		45, 28, 35, 41, 43, 56, 37, 38, 50, 52,
		33, 44, 37, 72, 47, 20,
	},
	LukeID: {
		80, 52, 38, 44, 39, 49, 50, 56, 62, 42,
		54, 59, 35, 35, 32, 31, 37, 43, 48, 47,
		38, 71, 56, 53,
	},
	JohnID: {
		51, 25, 36, 54, 47, 71, 53, 59, 41, 42,
		57, 50, 38, 31, 27, 33, 26, 40, 42, 31,
		25,
	},
	ActsID: {
		26, 47, 26, 37, 42, 15, 60, 40, 43, 48,
		30, 25, 52, 28, 41, 40, 34, 28, 41, 38,
		40, 30, 35, 27, 27, 32, 44, 31,
	},
	RomansID: {
		32, 29, 31, 25, 21, 23, 25, 39, 33, 21,
		36, 21, 14, 23, 33, 27,
	},
	FirstCorinthiansID: {
		31, 16, 23, 21, 13, 20, 40, 13, 27, 33,
		34, 31, 13, 40, 58, 24,
	},
	SecondCorinthiansID: {
		24, 17, 18, 18, 21, 18, 16, 24, 15, 18,
		33, 21, 14,
	},
	GalatiansID: {
		24, 21, 29, 31, 26, 18,
	},
	EphesiansID: {
		23, 22, 21, 32, 33, 24,
	},
	PhilippiansID: {
		30, 30, 21, 23,
	},
	ColossiansID: {
		29, 23, 25, 18,
	},
	FirstThessaloniansID: {
		10, 20, 13, 18, 28,
	},
	SecondThessaloniansID: {
		12, 17, 18,
	},
	FirstTimothyID: {
		20, 15, 16, 16, 25, 21,
	},
	SecondTimothyID: {
		18, 26, 17, 22,
	},
	TitusID: {
		16, 15, 15,
	},
	PhilemonID: {
		25,
	},
	HebrewsID: {
		14, 18, 19, 16, 14, 20, 28, 13, 28, 39,
		40, 29, 25,
	},
	JamesID: {
		27, 26, 18, 17, 20,
	},
	FirstPeterID: {
		25, 25, 22, 19, 14,
	},
	SecondPeterID: {
		21, 22, 18,
	},
	FirstJohnID: {
		10, 29, 24, 21, 21,
	},
	SecondJohnID: {
		13,
	},
	ThirdJohnID: {
		14,
	},
	JudeID: {
		25,
	},
	RevelationID: {
		20, 29, 22, 11, 14, 17, 17, 13, 21, 11,
		19, 17, 18, 20, 8, 21, 18, 24, 21, 15,
		27, 21,
	},
}
